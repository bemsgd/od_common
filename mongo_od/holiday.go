package mongo_od

type Holiday struct {
	Date        string `bson:"Date" json:"Date"`
	Description string `bson:"Description" json:"Description"`
	CreateBy    string `bson:"CreateBy" json:"CreateBy"`
}
