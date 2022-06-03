package mongo

import (
	"context"
	"fmt"

	"github.com/bemsgd/od_common/config"
	"github.com/bemsgd/od_common/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var appConfig = config.GetYamlAppConfig()

type MongoRepository struct {
	Context    context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

func ConnectMongoDB(dbUrl, dbName, collectionName string) MongoRepository {
	ctx := context.TODO()
	appDbUrl := fmt.Sprintf("%v%v?retryWrites=true&w=majority", dbUrl, dbName)
	clientOptions := options.Client().ApplyURI(appDbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	util.DieOnError("Error while connect to db", err)

	err = client.Ping(ctx, nil)
	util.DieOnError("Error while connect to db", err)

	collection := client.Database(dbName).Collection(collectionName)
	return MongoRepository{
		Collection: collection,
		Client:     client,
		Context:    ctx,
	}
}

func (tc *MongoRepository) DisconnectMongoDB() {
	tc.Client.Disconnect(tc.Context)
}

func (tc *MongoRepository) InsertDatas(datas []MongoData) (*mongo.InsertManyResult, error) {
	opts := options.InsertMany().SetOrdered(false)
	var docs []interface{}
	for _, f := range datas {
		docs = append(docs, f)
	}

	result, err := tc.Collection.InsertMany(tc.Context, docs, opts)
	util.LogOnError("error cannot insert []data : ", err)
	return result, err
}

func (tdb *MongoRepository) InsertData(data MongoData) (*mongo.InsertOneResult, error) {
	result, err := tdb.Collection.InsertOne(tdb.Context, data)
	util.LogOnError("error cannot insert data : ", err)
	return result, err
}
