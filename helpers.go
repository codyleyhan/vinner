package vinner

import "net/http"

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

func apiTypeToVehicleType(t string) VehicleType {
	switch t {
	case "Cabriolet/Convertible":
		return Convertible
	case "Minivan":
		return Minivan
	case "Coupe":
		return Coupe
	case "Low Speed Vehicle (LSV) / Neighborhood Electric Vehicle (NEV)":
		return LowSpeedVehicle
	case "Hatchback/Liftback/Notchback":
		return Hatchback
	case "Motorcycle - Standard":
		return StandardMotorcycle
	case "Sport Utility Vehicle (SUV)/Multi Purpose Vehicle (MPV)":
		return SportUtilityVehicle
	case "Crossover Utility Vehicle (CUV)":
		return CrossoverUtilityVehicle
	case "Van":
		return Van
	case "Roadster":
		return Roadster
	case "Truck":
		return Truck
	case "Motorcycle - Scooter":
		return Scooter
	case "Sedan/Saloon":
		return Sedan
	case "Wagon":
		return Wagon
	case "Bus":
		return Bus
	case "Pickup":
		return Pickup
	case "Trailer":
		return Trailer
	case "Truck - Tractor":
		return TractorTruck
	case "Streetcar / Trolley":
		return Streetcar
	case "Off-road Vehicle - All Terrain Vehicle (ATV) (Motorcycle-style)":
		return AllTerrainVehicle
	case "Bus - School Bus":
		return SchoolBus
	case "Racing Car":
		return RacingCar
	case "Motorcycle - Sport":
		return SportMotorcycle
	case "Motorcycle - Touring / Sport Touring":
		return TouringMotorcycle
	case "Motorcycle - Cruiser":
		return CruiserMotorcycle
	case "Motorcycle - Trike":
		return TrikeMotorcycle
	case "Off-road Vehicle - Dirt Bike / Off-Road":
		return DirtBike
	case "Motorcycle - Dual Sport / Adventure / Supermoto / On/Off-road":
		return DualSportMotorcycle
	case "Off-road Vehicle - Enduro (Off-road long distance racing)":
		return EnduroVehicle
	case "Motorcycle - Small / Minibike":
		return MiniBikeMotorcycle
	case "Off-road Vehicle - Go Kart":
		return GoKart
	case "Motorcycle - Side Car":
		return SideCarMotorcycle
	case "Motorcycle - Custom":
		return CustomMotorcycle
	case "Cargo Van":
		return CargoVan
	case "Snowmobile":
		return Snowmobile
	case "Motorcycle - Street":
		return StreetMotorcycle
	case "Motorcycle - Enclosed Three Wheeled / Enclosed Autocycle":
		return EnclosedThreeWheelMotorcycle
	case "Motorcycle - Unenclosed Three Wheeled / Open Autocycle":
		return UnenclosedThreeWheelMotorcycle
	case "Motorcycle - Moped":
		return Moped
	case "Off-road Vehicle - Recreational Off-Road Vehicle (ROV)":
		return RecreationalOffRoadVehicle
	case "Motorhome":
		return Motorhome
	case "Motorcycle - Cross Country":
		return CrossCountryMotorcycle
	case "Motorcycle - Underbone":
		return UnderboneMotorcycle
	case "Step Van / Walk-in Van":
		return StepVan
	case "Off-road Vehicle - Motocross (Off-road short distance, closed track racing) ":
		return MotocrossVehicle
	case "Motorcycle - Competition":
		return CompetitionMotorcycle
	case "Limousine":
		return Limousine
	case "Sport Utility Truck (SUT)":
		return SportUtilityVehicle
	case "Golf Cart":
		return GolfCart
	default:
		return UnkownVehicleType
	}
}
