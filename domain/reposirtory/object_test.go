package reposirtory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"

	"datastore/domain"
	"datastore/domain/reposirtory"
	"datastore/mongo/mocks"
)

func TestClient_Store(t *testing.T) {
	cases := map[string]struct {
		kind        string
		id          string
		name        string
		expectedErr error
	}{
		"Successfully store the objects in mongo": {
			kind:        "human",
			name:        "john doe",
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: nil,
		},
		"Failed to store the objects in mongo": {
			kind:        "human",
			name:        "john doe",
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: errors.New("something went wrong"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			mongoMockDB := new(mocks.IMongoClient)

			object := domain.Object{
				Kind: tc.kind,
				Id:   tc.id,
				Name: tc.name,
			}

			mongoMockDB.On(
				"InsertOne",
				mock.Anything,
				object,
			).Return(tc.expectedErr)

			repository := reposirtory.NewObjectRepository(mongoMockDB)
			err := repository.Store(context.Background(), object)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_GetObjectByID(t *testing.T) {
	cases := map[string]struct {
		id             string
		expectedErr    error
		expectedObject domain.Object
	}{
		"Successfully get the objects by id": {
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: nil,
			expectedObject: domain.Object{
				Kind: "human",
				Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
				Name: "Priya",
			},
		},
		"Failed to get objects using id": {
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: errors.New("something went wrong"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			mongoMockDB := new(mocks.IMongoClient)

			filter := bson.M{"id": tc.id}

			mongoMockDB.On(
				"Find",
				mock.Anything,
				filter,
			).Return(tc.expectedObject, tc.expectedErr)

			repository := reposirtory.NewObjectRepository(mongoMockDB)
			object, err := repository.GetObjectByID(context.Background(), tc.id)

			assert.Equal(t, tc.expectedObject, object)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_GetObjectByName(t *testing.T) {
	cases := map[string]struct {
		name           string
		expectedErr    error
		expectedObject domain.Object
		inserts        []string
	}{
		"Successfully get the objects by name": {
			name:        "John Doe",
			expectedErr: nil,
			expectedObject: domain.Object{
				Kind: "human",
				Id:   "f40e3be5-7b56-434f-b4e3-4d362277f307",
				Name: "John Doe",
			},
		},
		"Failed to get objects from mongo using name": {
			name:        "Anonymous",
			expectedErr: errors.New("not found"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			mongoMockDB := new(mocks.IMongoClient)

			filter := bson.M{"name": tc.name}

			mongoMockDB.On(
				"Find",
				mock.Anything,
				filter,
			).Return(tc.expectedObject, tc.expectedErr)

			repository := reposirtory.NewObjectRepository(mongoMockDB)
			object, err := repository.GetObjectByName(context.Background(), tc.name)

			assert.Equal(t, tc.expectedObject, object)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_ListObjects(t *testing.T) {
	cases := map[string]struct {
		kind            string
		expectedErr     error
		expectedObjects []domain.Object
	}{
		"Successfully get list of objects by kind": {
			kind:        "human",
			expectedErr: nil,
			expectedObjects: []domain.Object{
				{
					Kind: "human",
					Id:   "870bd42c-a2da-429a-a14a-5cbfb7fa3807",
					Name: "John Doe",
				},
				{
					Kind: "human",
					Id:   "d4a442c3-610d-473c-91e8-34c0a7ab2a1f",
					Name: "Jane Doe",
				},
			},
		},
		"Failed to get list of objects by kind": {
			kind:            "robots",
			expectedErr:     errors.New("not found"),
			expectedObjects: nil,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			mongoMockDB := new(mocks.IMongoClient)

			filter := bson.M{"kind": tc.kind}

			mongoMockDB.On(
				"FindAll",
				mock.Anything,
				filter,
			).Return(tc.expectedObjects, tc.expectedErr)

			repository := reposirtory.NewObjectRepository(mongoMockDB)
			objects, err := repository.ListObjects(context.Background(), tc.kind)

			assert.Equal(t, tc.expectedObjects, objects)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestClient_DeleteObject(t *testing.T) {
	cases := map[string]struct {
		id          string
		expectedErr error
	}{
		"Successfully delete the object by id": {
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: nil,
		},
		"Failed to delete the object by id": {
			id:          "f40e3be5-7b56-434f-b4e3-4d362277f307",
			expectedErr: errors.New("not found"),
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			mongoMockDB := new(mocks.IMongoClient)

			filter := bson.M{"id": tc.id}
			mongoMockDB.On(
				"Delete",
				mock.Anything,
				filter,
			).Return(tc.expectedErr)

			repository := reposirtory.NewObjectRepository(mongoMockDB)
			err := repository.DeleteObject(context.Background(), tc.id)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
