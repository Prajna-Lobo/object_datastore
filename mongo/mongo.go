package mongo

import (
	"context"

	"datastore/domain"
)

type IMongoClient interface {
	InsertOne(ctx context.Context, object domain.Object) error
	Find(ctx context.Context, filter interface{}) (domain.Object, error)
	FindAll(ctx context.Context, filter interface{}) ([]domain.Object, error)
	Delete(ctx context.Context, filter interface{}) error
}
