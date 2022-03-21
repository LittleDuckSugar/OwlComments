package repository

import (
	"context"
	"fmt"
	"owlcomments/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbUserPassword string = "22XIfjV32yQZcQAS"
const dbUsername string = "dbUser"

const UriDB string = "mongodb+srv://" + dbUsername + ":" + dbUserPassword + "@test-technique.ukgva.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

var Client *mongo.Client

// Should control db from here

// GetTarget return true if target exist, return false instead
func GetByTarget(targetId string) (bool, model.Comment) {

	// Ping the primary
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	coll := Client.Database("owlint").Collection("comment")

	var result model.Comment
	err := coll.FindOne(context.TODO(), bson.D{{"id", targetId}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return false, result
		}
		panic(err)
	}

	return true, result
}

// GetTarget return true if target exist, return false instead
func GetReplies(Id string) (bool, []model.Comment) {

	// Ping the primary
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	coll := Client.Database("owlint").Collection("comment")

	cursor, err := coll.Find(context.TODO(), bson.D{{"targetId", Id}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return false, []model.Comment{}
		}
		panic(err)
	}

	var results []model.Comment

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return true, results
}
