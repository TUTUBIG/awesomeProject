package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
)

var mongoCollection *mongo.Collection

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://218.30.91.134:27069/?connect=direct"))
	if err != nil {
		log.Fatal("e", err)
	}
	defer client.Disconnect(context.Background())

	mongoCollection = client.Database("dccn-uaa").Collection("email")

	fsn, err := os.Create(`./b.csv`)
	if err != nil {
		panic(err)
	}
	defer fsn.Close()
	w := csv.NewWriter(fsn)
	fs, err := os.Open(`./a.csv`)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		uid := row[4]
		email := getEmail(uid)
		row1 := make([]string, 0, len(row)+1)
		row1 = row[:]
		row1[len(row1)-1] = email
		if err := w.Write(row); err != nil {
			panic(err)
		}
		fmt.Println("done", row1[len(row1)-1])
	}
	w.Flush()
	return
}

func getEmail(uid string) string {
	r := make(map[string]interface{})
	if err := mongoCollection.FindOne(context.Background(), bson.D{bson.E{Key: "uid", Value: uid}}).Decode(&r); err != nil {
		panic(err)
	}
	return r["email"].(string)
}
