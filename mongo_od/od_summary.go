package mongo_od

type ODDataSummary struct {
	TravelDate        string              `json:"TravelDate"`
	ODMatchingSummary MatchingSummaryData `json:"ODMatchingSummary"`
	ODDetails         []ODDetail          `json:"ODDetails"`
}
