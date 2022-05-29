package mongoDB

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Collection *mongo.Collection // Collection 话柄
	Model      interface{}       // model
}

// func init() {
// 	var Collection *mongo.Collection // Collection 话柄
// }

func AddOne(mo *Mongo) {
	objId, err := mo.Collection.InsertOne(context.TODO(), &mo.Model)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Println("插入成功", objId)
}

func Del(m bson.M, mo *Mongo) {
	deleteResult, err := mo.Collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Collection.DeleteOne:", deleteResult)
}

func EditOne(m bson.M, mo *Mongo) {
	update := bson.M{"$set": mo.Model}
	updateResult, err := mo.Collection.UpdateOne(context.Background(), m, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Collection.UpdateOne:", updateResult)
}

func Update(m bson.M, mo *Mongo) {
	update := bson.M{"$set": mo.Model}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := mo.Collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Collection.UpdateOne:", updateResult)
}

func Sectle(m bson.M, mo *Mongo) {
	cur, err := mo.Collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var t struct{}
		if err = cur.Decode(&t); err != nil {
			log.Fatal(err)
		}
		log.Println("Collection.Find name=primitive.Regex{xx}: ", t)
	}
	_ = cur.Close(context.Background())
}

func GetOne(m bson.M, mo *Mongo) {
	// var one Mongo
	err := mo.Collection.FindOne(context.Background(), m).Decode(&mo.Model)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Collection.FindOne: ", mo.Model)
}

func GetList(m bson.M, mo *Mongo) {
	cur, err := mo.Collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []struct{}
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())

	log.Println("Collection.Find curl.All: ", all)
	for _, one := range all {
		// log.Println("Id:", one.Id, " - name:", one.Name, " - level:", one.Level)
		log.Println(one)
	}
}

func Count(mo *Mongo) {
	// fmt.Printf("%v", mo.Collection)
	count, err := mo.Collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("Collection.CountDocuments:", count)
}
