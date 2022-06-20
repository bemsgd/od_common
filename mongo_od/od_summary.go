package mongo_od

type ODDataSummary struct {
	TravelDate        string              `json:"TravelDate" bson:"TravelDate"`
	ODMatchingSummary MatchingSummaryData `json:"ODMatchingSummary" bson:"ODMatchingSummary"`
	ODDetails         []ODDetail          `json:"ODDetails" bson:"ODDetails"`
}
