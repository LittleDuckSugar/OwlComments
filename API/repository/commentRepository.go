package repository

import (
	"context"
	"fmt"
	"owlcomments/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbUserPassword string = "22XIfjV32yQZcQAS"
const dbUsername string = "dbUser"

const UriDB string = "mongodb+srv://" + dbUsername + ":" + dbUserPassword + "@test-technique.ukgva.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

var Client *mongo.Client

func init() {
	// Init the db connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(UriDB))

	if err != nil {
		panic(err)
	}

	Client = client
}

// GetTarget return true if target exist, return false instead
func GetCommentByTargetId(targetId string) (bool, []model.Comment) {

	pingDB()

	coll := Client.Database("owlint").Collection("comment")

	cursor, err := coll.Find(context.TODO(), bson.D{{"id", targetId}})
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

	for pass, comment := range results {
		results[pass].Replies = getReplies(comment.Id)
	}

	return true, results
}

// getReplies return replies of a given Id
func getReplies(Id string) []model.Comment {

	pingDB()

	coll := Client.Database("owlint").Collection("comment")

	cursor, err := coll.Find(context.TODO(), bson.D{{"targetId", Id}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return []model.Comment{}
		}
		panic(err)
	}

	var results []model.Comment

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for pos, comment := range results {
		results[pos].Replies = getReplies(comment.Id)
	}

	return results
}

// PostComment post to the db a new comment
func PostComment(comment model.NewComment) {

	pingDB()

	coll := Client.Database("owlint").Collection("comment")
	result, err := coll.InsertOne(context.TODO(), comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// pingDB try to ping the db - if no response then the request will be canceled
func pingDB() {
	// Ping the primary
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

// UpdateFakeTargets return a list of available targetId already available in the db
func UpdateFakeTargets() []string {
	pingDB()

	coll := Client.Database("owlint").Collection("comment")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return []string{}
		}
		panic(err)
	}

	var results []model.Comment

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	updatedFakeTargets := []string{}

	for _, comment := range results {
		updatedFakeTargets = append(updatedFakeTargets, comment.Id)
	}

	return updatedFakeTargets
}
