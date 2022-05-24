package mongoDB

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Table struct{}

var collection *mongo.Collection // collection 话柄

func AddOne(t *Table) {
	objId, err := collection.InsertOne(context.TODO(), &t)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Println("插入成功", objId)
}

func Del(m bson.M) {
	deleteResult, err := collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
}

func EditOne(t *Table, m bson.M) {
	update := bson.M{"$set": t}
	updateResult, err := collection.UpdateOne(context.Background(), m, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

func Update(t *Table, m bson.M) {
	update := bson.M{"$set": t}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

func Sectle(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var t Table
		if err = cur.Decode(&t); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{xx}: ", t)
	}
	_ = cur.Close(context.Background())
}

func GetOne(m bson.M) {
	var one Table
	err := collection.FindOne(context.Background(), m).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.FindOne: ", one)
}

func GetList(m bson.M) {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*Table
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	for _, one := range all {
		// log.Println("Id:", one.Id, " - name:", one.Name, " - level:", one.Level)
		log.Println(one)
	}
}

func Count() {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("collection.CountDocuments:", count)
}
