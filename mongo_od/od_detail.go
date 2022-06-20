package mongo_od

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bemsgd/od_common/travelling"
)

type ODDetail struct {
	TravelDate             string                     `bson:"TravelDate" json:"TravelDate"`
	TravelHour             string                     `bson:"TravelHour" json:"TravelHour"`
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

func (odd *ODDetail) SetMaxTimeTravellingChain(t travelling.TravellingChain) {
	odd.MaxTimeTravellingChain = t
}

func (odd *ODDetail) SetMinTimeTravellingChain(t travelling.TravellingChain) {
	odd.MinTimeTravellingChain = t
}

func (odd *ODDetail) SetMaxTravellingTime(duration time.Duration) {
	odd.MaxTravellingTime = duration
}

func (odd *ODDetail) SetMinTravellingTime(duration time.Duration) {
	odd.MinTravellingTime = duration
}

func (odd *ODDetail) UpdateTotalTravellingTime(duration time.Duration) {
	odd.TotalTravellingTime += duration
}

func (odd *ODDetail) IncTotalTransaction() {
	odd.TotalTransaction += 1
}

func (odd *ODDetail) GetAvgTravelingTime() time.Duration {
	return odd.TotalTravellingTime / time.Duration(odd.TotalTransaction)
}

func (odd *ODDetail) GetDetail() string {
	return fmt.Sprintf(`TravelDate: %v, Origin: %v, Destination: %v, MaxTravellingTime: %v, MinTravellingTime: %v, AvgTravelingTime: %v, TotalTransaction: %v`,
		odd.TravelDate, odd.Origin, odd.Destination, odd.MaxTravellingTime,
		odd.MinTravellingTime, odd.GetAvgTravelingTime(), odd.TotalTransaction)
}

func (odd ODDetail) GetJSON() string {
	result, _ := json.Marshal(odd)
	return string(result)
}
