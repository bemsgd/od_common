package mongo_od

import (
	"encoding/json"
	"fmt"
)

type SummaryOD interface {
	GetSummary() string
}

const (
	MatchingSummary = "MatchingSummary"
)

type MatchingSummaryData struct {
	TravelDate         string `bson:"TravelDate" json:"TravelDate"`
	TravelHour         string `bson:"TravelHour" json:"TravelHour"`
	UnMatchDestination int    `bson:"UnMatchDestination" json:"UnMatchDestination"`
	MatchDestion       int    `bson:"MatchDestion" json:"MatchDestion"`
}

func (ods *MatchingSummaryData) UpdateUnMatch(unmatch int) {
	ods.UnMatchDestination += unmatch
}

func (ods *MatchingSummaryData) UpdateMatch(match int) {
	ods.MatchDestion += match
}

func (ods *MatchingSummaryData) GetDetail() string {
	return fmt.Sprintf("TravelDate: %v, TravelHour: %v, UnMatchDestination: %v, MatchDestion: %v", ods.TravelDate, ods.TravelHour, ods.UnMatchDestination, ods.MatchDestion)
}
func (ods MatchingSummaryData) GetJSON() string {
	result, _ := json.Marshal(ods)
	return string(result)
}
