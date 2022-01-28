package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBDatabase *mongo.Database

func InitDatabase() {
	// TODO: add config.yaml
	co := options.Client().ApplyURI("mongodb://user:123456@localhost:27017/updownloader")

	client, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DBDatabase = client.Database("updownloader")

	log.Println("Connected to MongoDB!")
}

// TODO: 这里要用指针吗
func InsertRecord(record Record) {
	// TODO: 保存成全局变量？
	collection := DBDatabase.Collection("Record")
	_, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		log.Fatal(err)
	}
}

func QueryRecordByCode(code string) (Record, bool) {
	collection := DBDatabase.Collection("Record")
	findOneOpts := options.FindOne()
	singleResult := collection.FindOne(context.TODO(), bson.M{"code": code}, findOneOpts)
	var record Record
	if err := singleResult.Decode(&record); err == nil {
		return record, true
	}
	return Record{}, false
}
