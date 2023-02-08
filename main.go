/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//

	sw "tmf620/go"

	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

	"os"

	"github.com/go-http-utils/logger"
)

// Global var
var client *mongo.Client

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()
	/*
		client, err := mgo.GetMongoClientDefault()
		if err != nil {
			log.Fatal(err)
		}
		client.Ping(context.Background(), nil)
	*/
	/*
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://151320:admin124@shancluster.nkolnfm.mongodb.net/?retryWrites=true&w=majority").
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		coll := client.Database("tmf").Collection("Category")
		catalog := sw.CatalogCreate{CatalogType: "Atonement", Description: "Ian McEwan"}
		result, err := coll.InsertOne(context.TODO(), catalog)
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	*/
	//log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", logger.Handler(router, os.Stdout, logger.CombineLoggerType))
}
