package mongo_od

import (
	"log"

	"github.com/bemsgd/od_common/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Holiday struct {
	Date         string `bson:"Date" json:"Date"`
	Description  string `bson:"Description" json:"Description"`
	FreeTollFlag string `bson:"FreeTollFlag" json:"FreeTollFlag"`
	CreateBy     string `bson:"CreateBy" json:"CreateBy"`
}

func (adb *MongoRepository) IsHolidayDate(date string) bool {
	ctx := adb.Context
	collection := adb.Collection
	var h Holiday
	filter := bson.M{"Date": date}

	err := collection.FindOne(ctx, filter).Decode(&h)
	if err != nil {
		log.Print(err)
	}
	return h.Date != ""
}

func (adb *MongoRepository) IsFreeTollDay(date string) bool {
	ctx := adb.Context
	collection := adb.Collection
	var h Holiday
	filter := bson.M{"Date": date}

	err := collection.FindOne(ctx, filter).Decode(&h)
	if err != nil {
		log.Print(err)
	}

	return h.FreeTollFlag == "Y"
}

func (adb *MongoRepository) GetAllHoliday() []Holiday {
	ctx := adb.Context
	collection := adb.Collection
	var holidays []Holiday
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cur.All(ctx, &holidays)
	if err != nil {
		log.Fatal(err)
	}
	return holidays
}

func (adb *MongoRepository) AddHoliday(h Holiday) (*mongo.InsertOneResult, error) {
	result, err := adb.Collection.InsertOne(adb.Context, h)
	util.LogOnError("error cannot insert holiday data : ", err)
	return result, err
}

func (adb *MongoRepository) RemoveHoliday(h Holiday) (*mongo.DeleteResult, error) {
	result, err := adb.Collection.DeleteOne(adb.Context, h)
	util.LogOnError("error cannot delete holiday data : ", err)
	return result, err
}
