package mongo_od

type ODDataSummary struct {
	TravelDate        string              `json:"TravelDate" bson:"TravelDate"`
	TravelHour        string              `json:"TravelHour" bson:"TravelHour"`
	ODMatchingSummary MatchingSummaryData `json:"ODMatchingSummary" bson:"ODMatchingSummary"`
	ODDetails         []ODDetail          `json:"ODDetails" bson:"ODDetails"`
}
