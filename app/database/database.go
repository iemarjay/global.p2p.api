package database

import "go.mongodb.org/mongo-driver/bson"

type DbDecoder interface {
	Decode(v interface{}) error
}

type Database interface {
	Insert(data interface{}) (DbDecoder, error)
	FindOneByID(id string) DbDecoder
	FindOne(filter interface{}) DbDecoder
	UpdateOneById(id string, filter interface{}) (DbDecoder, error)
	UpdateOne(filter bson.M, data interface{}) (DbDecoder, error)
	Table(name string) Database
}