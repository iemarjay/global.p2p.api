package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	Insert(data interface{}) (*mongo.SingleResult, error)
	FindOneByID(id primitive.ObjectID) *mongo.SingleResult
	FindOne(filter interface{}) *mongo.SingleResult
	Table(name string) Database
}