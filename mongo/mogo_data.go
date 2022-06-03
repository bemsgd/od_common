package mongo

import (
	"encoding/json"
	"fmt"
)

type MongoData interface {
	GetJSON() string
}

const (
	MatchingSummary = "MatchingSummary"
)

type MatchingSummaryData struct {
	SummaryType        string `bson:"SummaryType" json:"SummaryType"`
	TravelDate         string `bson:"TravelDate" json:"TravelDate"`
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
	return fmt.Sprintf("TravelDate: %v, UnMatchDestination: %v, MatchDestion: %v", ods.TravelDate, ods.UnMatchDestination, ods.MatchDestion)
}
func (ods MatchingSummaryData) GetJSON() string {
	result, _ := json.Marshal(ods)
	return string(result)
}
