package miner

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"

	"log"
)

type MinerPenalty struct {
	Caller, Recevier, Reason string
	Burns, Currepoch         int64
}

func connet_mongodb() *mongo.Client {
	clientoptions := options.Client().ApplyURI("mongodb://root:jindi2021@10.32.2.2:27017")
	var ctx = context.TODO()
	cliecnt, err := mongo.Connect(ctx, clientoptions)

	if err != nil {
		log.Println("connect the mongodb filed", err.Error())
	}

	err = cliecnt.Ping(ctx, nil)

	if err != nil {
		log.Println("connect the mongodb filed", err.Error())
	}

	log.Println("Connected to MongoDB!")

	return cliecnt
}

func InsertmanyMinerPenplty(client *mongo.Client, data []interface{}) {
	collection := client.Database("MinerPenalty").Collection("penalties")

	insermany, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Println("insert failed", err.Error())
	}

	log.Println("insert successful!!!!!", insermany.InsertedIDs)

}

func InsertoneMinerPenplty(client *mongo.Client, data MinerPenalty) {
	collection := client.Database("MinerPenalty").Collection("Test")

	inserone, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("insert failed", err.Error())
	}

	log.Println("insert successful!!!!!", inserone.InsertedID)

}
