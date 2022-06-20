package travelling

import "time"

type TravellingChain struct {
	TotalPoint       int                     `bson:"TotalPoint" json:"TotalPoint"`
	TravellingDate   time.Time               `bson:"TravellingDate" json:"TravellingDate"`
	TravellingPoints []TravellingPointDetail `bson:"TravellingPoints" json:"TravellingPoints"`
}
type TravellingPointDetail struct {
	PointIndex       int              `bson:"PointIndex" json:"PoinPointIndextIndex"`
	PlacCode         int              `bson:"PlaceCode" json:"PlaceCode"`
	PlaceName        string           `bson:"PlaceName" json:"PlaceName"`
	LaneNo           int              `bson:"LaneNo" json:"LaneNo"`
	Fare             float64          `bson:"Fare" json:"Fare"`
	VehicleClass     int              `bson:"VehicleClass" json:"VehicleClass"`
	PaymentType      string           `bson:"PaymentType" json:"PaymentType"`
	DateTime         time.Time        `bson:"DateTime" json:"DateTime"`
	LicensePlateData LicensePlateData `bson:"LicensePlateData" json:"LicensePlateData"`
	PlateProvince    string           `bson:"PlateProvince" json:"PlateProvince"`
	VehicleColor     VehicleColor     `bson:"VehicleColor" json:"VehicleColor"`
	VehicleBrand     VehicleBrand     `bson:"VehicleBrand" json:"VehicleBrand"`
	VehicleBodyType  VehicleBodyType  `bson:"VehicleBodyType" json:"VehicleBodyType"`
	VehicleYear      VehicleYear      `bson:"VehicleYear" json:"VehicleYear"`
	VehicleModel     VehicleModel     `bson:"VehicleModel" json:"VehicleModel"`
	ImageUrl         string           `bson:"ImageUrl" json:"ImageUrl"`
}

type LicensePlateData struct {
	LicensePlate string  `bson:"LicensePlate" json:"LicensePlate"`
	Confidence   float64 `bson:"Confidence" json:"Confidence"`
}

type VehicleColor struct {
	Color      string  `bson:"Color" json:"Color"`
	Confidence float64 `bson:"Confidence" json:"Confidence"`
}

type VehicleBrand struct {
	Brand      string  `bson:"Brand" json:"Brand"`
	Confidence float64 `bson:"Confidence" json:"Confidence"`
}

type VehicleBodyType struct {
	BodyType   string  `bson:"Brand" json:"Brand"`
	Confidence float64 `bson:"Confidence" json:"Confidence"`
}

type VehicleYear struct {
	Year       string  `bson:"Brand" json:"Brand"`
	Confidence float64 `bson:"Confidence" json:"Confidence"`
}

type VehicleModel struct {
	Model      string  `bson:"Brand" json:"Brand"`
	Confidence float64 `bson:"Confidence" json:"Confidence"`
}
