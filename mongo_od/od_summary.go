package mongo_od

import (
	"time"

	"github.com/bemsgd/od_common/travelling"
)

type ODDataSummary struct {
	TravelYear             string                     `json:"TravelYear" bson:"TravelYear"`
	TravelMonth            string                     `json:"TravelMonth" bson:"TravelMonth"`
	TravelDate             string                     `json:"TravelDate" bson:"TravelDate"`
	TravelHour             string                     `json:"TravelHour" bson:"TravelHour"`
	Origin                 string                     `bson:"Origin" json:"Origin"`
	Destination            string                     `bson:"Destination" json:"Destination"`
	VehicleClass           string                     `bson:"VehicleClass" json:"VehicleClass"`
	PaymentType            string                     `bson:"PaymentType" json:"PaymentType"`
	MaxTravellingTime      time.Duration              `bson:"MaxTravellingTime" json:"MaxTravellingTime"`
	MinTravellingTime      time.Duration              `bson:"MinTravellingTime" json:"MinTravellingTime"`
	TotalTravellingTime    time.Duration              `bson:"TotalTravellingTime" json:"TotalTravellingTime"`
	TotalTransaction       int                        `bson:"TotalTransaction" json:"TotalTransaction"`
	MaxTimeTravellingChain travelling.TravellingChain `bson:"MaxTimeTravellingChain" json:"MaxTimeTravellingChain"`
	MinTimeTravellingChain travelling.TravellingChain `bson:"MinTimeTravellingChain" json:"MinTimeTravellingChain"`
}

func (odd *ODDataSummary) SetMaxTimeTravellingChain(t travelling.TravellingChain) {
	odd.MaxTimeTravellingChain = t
}

func (odd *ODDataSummary) SetMinTimeTravellingChain(t travelling.TravellingChain) {
	odd.MinTimeTravellingChain = t
}

func (odd *ODDataSummary) SetMaxTravellingTime(duration time.Duration) {
	odd.MaxTravellingTime = duration
}

func (odd *ODDataSummary) SetMinTravellingTime(duration time.Duration) {
	odd.MinTravellingTime = duration
}

func (odd *ODDataSummary) UpdateTotalTravellingTime(duration time.Duration) {
	odd.TotalTravellingTime += duration
}

func (odd *ODDataSummary) IncTotalTransaction() {
	odd.TotalTransaction += 1
}

func (odd *ODDataSummary) GetAvgTravelingTime() time.Duration {
	return odd.TotalTravellingTime / time.Duration(odd.TotalTransaction)
}
