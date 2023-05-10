package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"datastore/config"
	"datastore/domain"
	"datastore/examples"
	"datastore/mongo"
)

func main() {
	appConfig := config.LoadConfiguration("config/config.json")
	mClient := mongo.NewClient(getMongoURI(appConfig), appConfig.Mongo.Database, appConfig.Mongo.Collection)

	ctx := context.Background()

	err := insertPerson(ctx, mClient)
	if err != nil {
		log.Println(err.Error())
	}

	err2 := insertAnimal(ctx, mClient)
	if err2 != nil {
		log.Println(err.Error())
	}
}

func getMongoURI(config *config.Configuration) string {
	return "mongodb://" + os.Getenv("MONGO_USER") + ":" +
		os.Getenv("MONGO_PASS") + "@" + config.Mongo.Host + ":" + config.Mongo.Port
}

func objectMapper[T domain.IObject](object T) domain.Object {
	return domain.Object{
		Kind: reflect.TypeOf(object).String(),
		Id:   object.GetID(),
		Name: object.GetName(),
	}
}

func insertPerson(ctx context.Context, mClient mongo.IMongoClient) error {

	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, "2000-06-19")
	if err != nil {
		fmt.Println(err)

		return err
	}

	person := examples.NewPerson(
		"John Doe",
		"34d6d8c5-e4a3-4f83-b4f3-d2c3f2b33f27",
		"Doe",
		"19 Aug",
		birthDate)

	object := objectMapper(person)

	mErr := mClient.InsertOne(ctx, object)
	if mErr != nil {
		fmt.Println(mErr)

		return mErr
	}

	fmt.Println("---person---")
	fmt.Println("ID", person.GetID())
	fmt.Println("Name", person.GetName())
	fmt.Println("Kind", person.GetKind())

	return nil
}

func insertAnimal(ctx context.Context, mClient mongo.IMongoClient) error {
	animal := examples.NewAnimal(
		"dog",
		"08adc480-a355-4760-89d4-d0245e62b77b",
		"896478bf-a66b-4fb6-a226-2efaf157a517",
	)

	animalObject := objectMapper(animal)

	err := mClient.InsertOne(ctx, animalObject)
	if err != nil {
		fmt.Println(err)

		return err
	}

	fmt.Println("---animal---")
	fmt.Println("ID", animal.GetID())
	fmt.Println("Name", animal.GetName())
	fmt.Println("Kind", animal.GetKind())

	return nil
}
