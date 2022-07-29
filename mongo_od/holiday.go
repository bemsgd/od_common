package mongo_od

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Holiday struct {
	Date        string `bson:"Date" json:"Date"`
	Description string `bson:"Description" json:"Description"`
	CreateBy    string `bson:"CreateBy" json:"CreateBy"`
}

func (adb *MongoRepository) IsHolidayDate(date string) bool {
	ctx := adb.Context
	collection := adb.Collection
	var h Holiday
	filter := bson.M{"Date": date}

	err := collection.FindOne(ctx, filter).Decode(&h)
	if err != nil {
		log.Fatal(err)
	}
	return h.Date != ""
}
