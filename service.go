package vinner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const vehicleAPI = "https://vpic.nhtsa.dot.gov/api/vehicles"

var InvalidVIN = fmt.Errorf("invalid VIN")

type serviceOptionFunc func(*serviceOptions)

type serviceOptions struct {
	timeout   time.Duration
	transport http.RoundTripper
}

type Service struct {
	client *http.Client
}

func NewService(opts ...serviceOptionFunc) *Service {
	options := serviceOptions{}

	for _, opt := range opts {
		opt(&options)
	}

	return &Service{
		client: &http.Client{
			Transport: options.transport,
			Timeout:   options.timeout,
		},
	}
}

func (s *Service) GetVehicle(ctx context.Context, vin string) (*Vehicle, error) {
	if s.client == nil {
		s.client = &http.Client{}
	}
	if len(vin) != 17 {
		return nil, InvalidVIN
	}

	requestURL := fmt.Sprintf("%s/decodevin/%s?format=json", vehicleAPI, vin)

	var response getVehicleResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = withHTTPClose(res, func() error {
		if res.StatusCode != 200 {
			return fmt.Errorf("could not get vehicle: received HTTP response %d", res.StatusCode)
		}

		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return fmt.Errorf("could not decode get vehicle response: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return vehicleFromResponse(&response)
}

func (s *Service) GetMakes(ctx context.Context) ([]string, error) {
	if s.client == nil {
		s.client = &http.Client{}
	}

	requestURL := fmt.Sprintf("%s/getallmakes?format=json", vehicleAPI)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var response getMakesResponse

	err = withHTTPClose(res, func() error {
		if res.StatusCode != 200 {
			return fmt.Errorf("could not get makes: received HTTP response %d", res.StatusCode)
		}

		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return fmt.Errorf("could not decode get makes response: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var makes []string

	for _, result := range response.Results {
		makes = append(makes, result.Make)
	}

	return makes, nil
}

func vehicleFromResponse(res *getVehicleResponse) (*Vehicle, error) {
	fieldSet := map[apiField]bool{
		yearField:      true,
		makeField:      true,
		modelField:     true,
		trimField:      true,
		doorsField:     true,
		bodyClassField: true,
		errorCodeField: true,
		errorTextField: true,
	}

	values := make(map[apiField]*string)

	for _, res := range res.Results {
		if fieldSet[res.Variable] && res.Value != nil {
			values[res.Variable] = res.Value
		}
	}

	if values[errorCodeField] != nil && *values[errorCodeField] != "0" {
		return nil, InvalidVIN
	}

	vehicle := Vehicle{
		VIN: strings.TrimPrefix(res.SearchCriteria, "VIN:"),
	}

	if values[yearField] != nil {
		year, err := strconv.Atoi(*values[yearField])
		if err != nil {
			return nil, fmt.Errorf("unable to parse the year %q", *values[yearField])
		}
		vehicle.Year = year
	}

	if values[makeField] != nil {
		vehicle.Make = *values[makeField]
	}

	if values[modelField] != nil {
		vehicle.Model = *values[modelField]
	}

	if values[trimField] != nil {
		vehicle.Trim = *values[trimField]
	}

	if values[doorsField] != nil {
		doors, err := strconv.Atoi(*values[doorsField])
		if err != nil {
			return nil, fmt.Errorf("unable to parse the number of doors %q", *values[doorsField])
		}
		vehicle.Doors = doors
	}

	if values[bodyClassField] != nil {
		switch *values[bodyClassField] {
		case "Sedan/Saloon":
			vehicle.BodyClass = Sedan
		case "Truck":
			vehicle.BodyClass = Truck
		case "Motorcycle":
			vehicle.BodyClass = Motorcycle
		case "Wagon":
			vehicle.BodyClass = Wagon
		default:
			vehicle.BodyClass = UnkownVehicleType
		}
	}

	return &vehicle, nil
}

func withHTTPClose(res *http.Response, fn func() error) error {
	err := fn()
	err2 := res.Body.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}

	return nil
}
