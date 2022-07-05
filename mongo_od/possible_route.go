package mongo_od

import (
	"context"
	"log"

	"github.com/bemsgd/od_common/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PossibleRoute struct {
	Entry  string `bson:"entry" json:"entry"`
	Exits  []Exit `bson:"exits" json:"exits"`
	DestDb string `bson:"destDb" json:"destDb"`
}

type Exit struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"exit" json:"exit"`
	Distance float64 `bson:"distance" json:"distance"`
}

type ApplicationDbService struct {
	Context    context.Context
	Collection *mongo.Collection
	Client     *mongo.Client
}

func (adb *MongoRepository) GetPossibleRoute(entryId string) PossibleRoute {
	ctx := adb.Context
	collection := adb.Collection
	var possibleRoute PossibleRoute
	filter := bson.M{"entry": entryId}

	err := collection.FindOne(ctx, filter).Decode(&possibleRoute)
	if err != nil {
		log.Fatal(err)
	}
	return possibleRoute
}

func (adb *MongoRepository) GetPossibleRoutes() map[string]PossibleRoute {
	ctx := adb.Context
	collection := adb.Collection
	var possibleRoutes []PossibleRoute
	possibleRoutesResult := make(map[string]PossibleRoute)
	filter := bson.M{}

	cur, err := collection.Find(ctx, filter)
	util.DieOnError("error find possible route", err)
	err = cur.All(ctx, &possibleRoutes)
	util.DieOnError("error scan possible route", err)

	for _, route := range possibleRoutes {
		possibleRoutesResult[route.Entry] = route
	}
	return possibleRoutesResult
}
