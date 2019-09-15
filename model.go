package vinner

const (
	Convertible                    VehicleType = "Convertible"
	Minivan                        VehicleType = "Minivan"
	Coupe                          VehicleType = "Coupe"
	LowSpeedVehicle                VehicleType = "Low Speed Vehicle (LSV)"
	Hatchback                      VehicleType = "Hatchback"
	StandardMotorcycle             VehicleType = "Standard Motorcycle"
	SportUtilityVehicle            VehicleType = "Sport Utility Vehicle (SUV)"
	CrossoverUtilityVehicle        VehicleType = "Crossover Utility Vehicle (CUV)"
	Van                            VehicleType = "Van"
	Roadster                       VehicleType = "Roadster"
	Truck                          VehicleType = "Truck"
	Scooter                        VehicleType = "Scooter"
	Sedan                          VehicleType = "Sedan"
	Wagon                          VehicleType = "Wagon"
	Bus                            VehicleType = "Bus"
	Pickup                         VehicleType = "Pickup"
	Trailer                        VehicleType = "Trailer"
	TractorTruck                   VehicleType = "Tractor Truck"
	Streetcar                      VehicleType = "Streetcar"
	AllTerrainVehicle              VehicleType = "All Terrain Vehicle (ATV)"
	SchoolBus                      VehicleType = "School Bus"
	RacingCar                      VehicleType = "Racing Car"
	SportMotorcycle                VehicleType = "Sport Motorcycle"
	TouringMotorcycle              VehicleType = "Touring Motorcycle"
	CruiserMotorcycle              VehicleType = "Cruiser Motorcycle"
	TrikeMotorcycle                VehicleType = "Trike Motorcycle"
	DirtBike                       VehicleType = "Dirt Bike"
	DualSportMotorcycle            VehicleType = "Dual Sport Motorcycle"
	EnduroVehicle                  VehicleType = "Enduro Vehicle"
	MiniBikeMotorcycle             VehicleType = "Minibike Motorcycle"
	GoKart                         VehicleType = "Go Kart"
	SideCarMotorcycle              VehicleType = "Side Car Motorcycle"
	CustomMotorcycle               VehicleType = "Custom Motorcycle"
	CargoVan                       VehicleType = "Cargo Van"
	Snowmobile                     VehicleType = "Snowmobile"
	StreetMotorcycle               VehicleType = "Street Motorcycle"
	EnclosedThreeWheelMotorcycle   VehicleType = "Enclosed Three Wheeled Motorcycle"
	UnenclosedThreeWheelMotorcycle VehicleType = "Unenclosed Three Wheeled Motorcycle"
	Moped                          VehicleType = "Moped"
	RecreationalOffRoadVehicle     VehicleType = "Recreational Off-Road Vehicle (ROV)"
	Motorhome                      VehicleType = "Motorhome"
	CrossCountryMotorcycle         VehicleType = "Cross Country Motorcycle"
	UnderboneMotorcycle            VehicleType = "Underbone Motorcycle"
	StepVan                        VehicleType = "Step Van"
	MotocrossVehicle               VehicleType = "MotocrossVehicle"
	CompetitionMotorcycle          VehicleType = "Competition Motorcycle"
	Limousine                      VehicleType = "Limousine"
	SportUtilityTruck              VehicleType = "Sport Utility Truck (SUT)"
	GolfCart                       VehicleType = "Golf Cart"
	UnkownVehicleType              VehicleType = "Unknown"
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
