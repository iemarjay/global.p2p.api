package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDatabase struct {
	connection *mongo.Database
	collection *mongo.Collection
}

func (db *MongoDatabase) Collection() *mongo.Collection {
	return db.collection
}

func (db *MongoDatabase) Connect(uri string, dbName string) *MongoDatabase {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI(uri)
	clientOptions = clientOptions.SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db.connection = client.Database(dbName)

	return db
}

func (db *MongoDatabase) Table(name string) Database {
	db.collection = db.connection.Collection(name)

	return db
}

func (db MongoDatabase) Insert(data interface{}) (*mongo.SingleResult, error) {
	ior, err := db.collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}

	insertedID := ior.InsertedID.(primitive.ObjectID)

	return db.findOneByID(insertedID), err
}

func (db *MongoDatabase) FindOneByID(id string) *mongo.SingleResult {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return db.findOneByID(objectId)
}

func (db *MongoDatabase) findOneByID(id primitive.ObjectID) *mongo.SingleResult {
	result := db.collection.FindOne(context.Background(), bson.M{"_id": id})
	return result
}

func (db *MongoDatabase) FindOne(filter interface{}) *mongo.SingleResult {
	result := db.collection.FindOne(context.Background(), filter)
	return result
}

func NewMongo(uri string, dbName string) *MongoDatabase {
	m := &MongoDatabase{}
	return m.Connect(uri, dbName)
}