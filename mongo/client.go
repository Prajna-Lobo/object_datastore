package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"datastore/domain"
)

type Client struct {
	client     *mongo.Client
	db         string
	collection string
}

// NewClient Connecting to mongo
func NewClient(mongoURI, database, collection string) IMongoClient {
	mClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Please check the mongo user and password in env", err)
	}

	return &Client{
		client:     mClient,
		db:         database,
		collection: collection,
	}
}

// getCollectionName gets collection which is injected while initializing the client
func (c Client) getCollectionName() *mongo.Collection {
	return c.client.Database(c.db).Collection(c.collection)
}

// InsertOne stores the object into mongo
func (c Client) InsertOne(ctx context.Context, object domain.Object) error {
	objectCollection := c.getCollectionName()
	_, err := objectCollection.InsertOne(ctx, object)
	if err != nil {
		return err
	}

	return nil
}

// Find filters objects from mongo by using id
func (c Client) Find(ctx context.Context, filter interface{}) (domain.Object, error) {
	objectCollection := c.getCollectionName()

	result := objectCollection.FindOne(ctx, filter, nil)

	var mObject domain.Object
	err := result.Decode(&mObject)
	if err != nil {
		return mObject, err
	}

	return mObject, nil
}

// FindAll filters objects from mongo by using kind
func (c Client) FindAll(ctx context.Context, filter interface{}) ([]domain.Object, error) {
	objectCollection := c.getCollectionName()

	result, err := objectCollection.Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	var objects []domain.Object
	err = result.All(ctx, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

// Delete deletes the object from mongo by id
func (c Client) Delete(ctx context.Context, filter interface{}) error {
	objectCollection := c.getCollectionName()

	_, err := objectCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
