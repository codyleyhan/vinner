package vinner

const (
	Sedan             VehicleType = "sedan"
	Wagon             VehicleType = "wagon"
	Motorcycle        VehicleType = "motorcycle"
	Truck             VehicleType = "truck"
	UnkownVehicleType VehicleType = "unknown"
)

type VehicleType string

type Vehicle struct {
	VIN       string      `json:"vin"`
	Year      int         `json:"year"`
	Make      string      `json:"make"`
	Model     string      `json:"model"`
	Trim      string      `json:"trim"`
	Doors     int         `json:"doors"`
	BodyClass VehicleType `json:"body_class"`
}

type getVehicleResult struct {
	Value    *string
	Variable apiField
}

type getVehicleResponse struct {
	SearchCriteria string
	Results        []getVehicleResult
}

type getMakesResult struct {
	Make string `json:"Make_Name"`
}

type getMakesResponse struct {
	Results []getMakesResult
}

type GetModelsRequest struct {
	Make string
	Year int
}

type getModelsResult struct {
	ModelName string `json:"Model_Name"`
}

type getModelsMakeResponse struct {
	Results []getModelsResult
}

type apiField string

const (
	// GetVin Fields
	// vehicle Fields
	yearField      apiField = "Model Year"
	makeField      apiField = "Make"
	modelField     apiField = "Model"
	trimField      apiField = "Trim"
	doorsField     apiField = "Doors"
	bodyClassField apiField = "Body Class"

	// api response fields
	errorCodeField apiField = "Error Code"
	errorTextField apiField = "Additional Error Text"
)
