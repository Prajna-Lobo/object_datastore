package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"datastore/domain"
	"datastore/mongo"
)

var mongoTestDbURI = mongo.GetMongoTestURI()

func TestClient_InsertOne(t *testing.T) {
	cases := map[string]struct {
		object      domain.Object
		expectedErr error
	}{
		"Successfully insert object into mongo db": {
			object: domain.Object{
				Kind: "human",
				Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
				Name: "John Doe",
			},
			expectedErr: nil,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			client := mongo.NewClient(mongoTestDbURI, mongo.TestDb, mongo.TestCollection)
			err := client.InsertOne(context.Background(), tc.object)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_Find(t *testing.T) {
	cases := map[string]struct {
		filter         interface{}
		expectedObject domain.Object
		expectedErr    error
	}{
		"Successfully fetch objects from mongo by id": {
			filter: bson.M{"id": "f40e3be5-7b56-434f-b4e3-4d362277f307"},
			expectedObject: domain.Object{
				Kind: "human",
				Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
				Name: "John Doe",
			},
			expectedErr: nil,
		},
		"Successfully fetch objects from mongo by name": {
			filter: bson.M{"name": "John Doe"},
			expectedObject: domain.Object{
				Kind: "human",
				Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
				Name: "John Doe",
			},
			expectedErr: nil,
		},
	}

	mTestClient := mongo.TestDbConnection()
	mongo.AddMultipleObjectIntoTestDb(context.Background(), mTestClient)

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {

			client := mongo.NewClient(mongoTestDbURI, mongo.TestDb, mongo.TestCollection)
			object, err := client.Find(context.Background(), tc.filter)

			assert.Equal(t, tc.expectedObject, object)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_FindAll(t *testing.T) {
	cases := map[string]struct {
		filter          interface{}
		expectedObjects []domain.Object
		expectedErr     error
	}{
		"Successfully fetch all objects from mongo by kind": {
			filter: bson.M{"kind": "human"},
			expectedObjects: []domain.Object{
				{
					Kind: "human",
					Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
					Name: "John Doe",
				},
				{
					Kind: "human",
					Id:   "870bd42c-a2da-429a-a14a-5cbfb7fa3807",
					Name: "Jane Doe",
				},
			},
			expectedErr: nil,
		},
	}

	mTestClient := mongo.TestDbConnection()
	mongo.AddMultipleObjectIntoTestDb(context.Background(), mTestClient)

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {

			client := mongo.NewClient(mongoTestDbURI, mongo.TestDb, mongo.TestCollection)
			objects, err := client.FindAll(context.Background(), tc.filter)

			assert.Equal(t, tc.expectedObjects, objects)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_Delete(t *testing.T) {
	cases := map[string]struct {
		filter      interface{}
		expectedErr error
	}{
		"Successfully delete the object from mongo": {
			filter:      bson.M{"id": "f40e3be5-7b56-434f-b4e3-4d362277f307"},
			expectedErr: nil,
		},
	}

	ctx := context.Background()

	mTestClient := mongo.TestDbConnection()
	mongo.AddMultipleObjectIntoTestDb(ctx, mTestClient)

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			client := mongo.NewClient(mongoTestDbURI, mongo.TestDb, mongo.TestCollection)
			err := client.Delete(ctx, tc.filter)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
