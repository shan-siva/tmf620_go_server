// Copyright 2018 Kuei-chun Chen. All rights reserved.

package mongo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "argos"
const collectionName = "cars"
const collectionFavorites = "favorites"
const collectionExamples = "examples"

var once sync.Once

// GetMongoClient returns *mongo.Client
func GetMongoClient(uri string) (*mongo.Client, error) {
	return getMongoClientByURI(uri)
}

func GetMongoClientDefault() (*mongo.Client, error) {
	return getMongoClient()
}

func getMongoClient() (*mongo.Client, error) {
	uri := "mongodb+srv://151320:admin124@shancluster.nkolnfm.mongodb.net/?retryWrites=true&w=majority"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	fmt.Printf("URI is %s \n", uri)
	return getMongoClientByURI(uri)
}

func getMongoClientByURI(uri string) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		return client, err
	}
	if err != nil {
		fmt.Printf("Error conencting to Mongo : %s \n", err)
	} else {
		fmt.Printf("Mongo Connected \n\n")
	}
	client.Ping(context.Background(), nil)
	return client, err
}

// MongoPipeline gets aggregation pipeline from a string
func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}

func toInt64(num interface{}) int64 {
	f := fmt.Sprintf("%v", num)
	x, _ := strconv.ParseFloat(f, 64)
	return int64(x)
}

func InsertOne(coll string, document interface{}) *mongo.InsertOneResult {
	var result *mongo.InsertOneResult
	var collection *mongo.Collection
	var ctx = context.Background()

	client, err := GetMongoClientDefault()

	collection = client.Database("tmf").Collection(coll)
	fmt.Printf("collection Success %s\n", collection.Name())

	if result, err = collection.InsertOne(ctx, document); err != nil {
		fmt.Printf("Error when InsertOne %s \n", err)
	}
	fmt.Printf("Inserted Success %s\n", result.InsertedID)
	return result
}
