package reposirtory

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"datastore/domain"
	"datastore/mongo"
)

type ObjectRepository struct {
	mClient mongo.IMongoClient
}

// NewObjectRepository Connecting to mongo
func NewObjectRepository(mongoClient mongo.IMongoClient) domain.ObjectDB {
	return &ObjectRepository{
		mClient: mongoClient,
	}
}

// Store stores the object into mongo
func (o ObjectRepository) Store(ctx context.Context, object domain.Object) error {
	return o.mClient.InsertOne(ctx, object)
}

// GetObjectByID filters objects from mongo by using id
func (o ObjectRepository) GetObjectByID(ctx context.Context, id string) (domain.Object, error) {
	filter := bson.M{"id": id}

	foundObject, err := o.mClient.Find(ctx, filter)
	if err != nil {
		return domain.Object{}, err
	}

	return foundObject, nil
}

// GetObjectByName filters objects from mongo by using name
func (o ObjectRepository) GetObjectByName(ctx context.Context, name string) (domain.Object, error) {
	filter := bson.M{"name": name}

	foundObject, err := o.mClient.Find(ctx, filter)
	if err != nil {
		return domain.Object{}, err
	}

	return foundObject, nil
}

// ListObjects filters objects from mongo by using kind
func (o ObjectRepository) ListObjects(ctx context.Context, kind string) ([]domain.Object, error) {
	filter := bson.M{"kind": kind}

	objects, err := o.mClient.FindAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

// DeleteObject deletes the object from mongo by id
func (o ObjectRepository) DeleteObject(ctx context.Context, id string) error {
	filter := bson.M{"id": id}

	err := o.mClient.Delete(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
