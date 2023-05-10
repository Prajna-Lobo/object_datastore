package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"datastore/config"
	"datastore/domain"
)

const TestDb = "test_datastore"
const TestCollection = "test_objects"

func GetMongoTestURI() string {
	testConfig := config.LoadConfiguration("../config/config.json")
	return "mongodb://" + os.Getenv("MONGO_USER") + ":" +
		os.Getenv("MONGO_PASS") + "@" + testConfig.Mongo.Host + ":" + testConfig.Mongo.Port
}

func TestDbConnection() *mongo.Client {
	mTestClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(GetMongoTestURI()))
	if err != nil {
		panic(err)
	}

	return mTestClient
}

func AddMultipleObjectIntoTestDb(ctx context.Context, client *mongo.Client) {
	CleanDB(ctx, client)

	mClient := NewClient(GetMongoTestURI(), TestDb, TestCollection)
	obj1 := domain.Object{
		Kind: "human",
		Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
		Name: "John Doe",
	}

	err := mClient.InsertOne(ctx, obj1)
	if err != nil {
		fmt.Println(err)
	}

	obj2 := domain.Object{
		Kind: "human",
		Id:   "870bd42c-a2da-429a-a14a-5cbfb7fa3807",
		Name: "Jane Doe",
	}

	err2 := mClient.InsertOne(ctx, obj2)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func CleanDB(ctx context.Context, mClient *mongo.Client) {
	collection := mClient.Database(TestDb).Collection(TestCollection)
	err := collection.Drop(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
