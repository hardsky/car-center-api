package models

type Engine struct {
	ID               int64
	Capacity         uint16 `json:"capacity"`
	NumCylinders     uint8  `json:"numCylinders"`
	MaxRpm           uint16 `json:"maxRpm"`
	ManufacturerCode string `json:"manufacturerCode"`
}

type FuelFigure struct {
	ID               int64
	Speed            uint16  `json:"speed"`
	Mpg              float64 `json:"mpg"`
	UsageDescription string  `json:"usageDescription"`
}

type PerformanceFigure struct {
	ID           int64
	OctaneRating uint16 `json:"octaneRating"`

	AccelerationID int64
	Acceleration   *Acceleration `json:"acceleration"`
}

type Acceleration struct {
	ID      int64
	Mph     uint16
	Seconds float64
}

type Car struct {
	SerialNumber   uint64 `json:"serialNumber" sql:",pk"`
	OwnerName      string `json:"ownerName"`
	ModelYear      uint64 `json:"modelYear"`
	Code           string `json:"code"`
	VehicleCode    string `json:"vehicleCode"`
	Manufacturer   string `json:"manufacturer"`
	Model          string `json:"model"`
	ActivationCode string `json:"activationCode"`

	EngineID int64
	Engine   *Engine `json:"engine"`

	FuelFigureID int64
	FuelFigure   *FuelFigure `json:"fuelFigures"`

	PerformanceFigureID int64
	PerformanceFigure   *PerformanceFigure `json:"performanceFigures"`
}
