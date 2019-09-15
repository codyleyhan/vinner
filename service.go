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
var InvalidModelsRequest = fmt.Errorf("vehicle make is required")

type serviceOptionFunc func(*serviceOptions)

type serviceOptions struct {
	timeout   time.Duration
	transport http.RoundTripper
}

type Service struct {
	client *http.Client
}

func WithTimeout(timeout time.Duration) serviceOptionFunc {
	return func(opts *serviceOptions) {
		opts.timeout = timeout
	}
}

func WithTransport(transport http.RoundTripper) serviceOptionFunc {
	return func(opts *serviceOptions) {
		opts.transport = transport
	}
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
	if len(vin) != 17 {
		return nil, InvalidVIN
	}

	requestURL := fmt.Sprintf("%s/decodevin/%s?format=json", vehicleAPI, vin)

	var response getVehicleResponse

	if err := s.makeRequest(ctx, requestURL, &response); err != nil {
		return nil, fmt.Errorf("failed to get vehicle from VIN: %w", err)
	}

	return vehicleFromResponse(&response)
}

func (s *Service) GetMakes(ctx context.Context) ([]string, error) {
	requestURL := fmt.Sprintf("%s/getallmakes?format=json", vehicleAPI)

	var response getMakesResponse

	if err := s.makeRequest(ctx, requestURL, &response); err != nil {
		return nil, fmt.Errorf("failed to get vehicle makes: %w", err)
	}

	var makes []string

	for _, result := range response.Results {
		makes = append(makes, result.Make)
	}

	return makes, nil
}

func (s *Service) GetModels(ctx context.Context, req GetModelsRequest) ([]string, error) {
	if req.Make == "" {
		return nil, InvalidModelsRequest
	}

	var requestURL string

	if req.Year == 0 {
		requestURL = fmt.Sprintf("%s/getmodelsformake/%s?format=json", vehicleAPI, req.Make)
	} else {
		requestURL = fmt.Sprintf("%s/getmodelsformakeyear/make/%s/modelyear/%d?format=json", vehicleAPI, req.Make, req.Year)
	}

	var response getModelsMakeResponse

	if err := s.makeRequest(ctx, requestURL, &response); err != nil {
		return nil, fmt.Errorf("failed to get vehicle models: %w", err)
	}

	var models []string

	for _, result := range response.Results {
		models = append(models, result.ModelName)
	}

	return models, nil
}

func (s *Service) makeRequest(ctx context.Context, url string, dest interface{}) error {
	if s.client == nil {
		s.client = &http.Client{}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}

	err = withHTTPClose(res, func() error {
		if res.StatusCode != 200 {
			return fmt.Errorf("received bad response: received HTTP response %d", res.StatusCode)
		}

		if err := json.NewDecoder(res.Body).Decode(dest); err != nil {
			return fmt.Errorf("could not decode response: %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
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
		vehicle.BodyClass = apiTypeToVehicleType(*values[bodyClassField])
	}

	return &vehicle, nil
}
