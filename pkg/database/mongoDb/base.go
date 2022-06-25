package mongoDB

import (
	"app-log/app/repo"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoLogRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(c *mongo.Collection) repo.ILogRepo {
	return &mongoLogRepository{collection: c}
}

func (m *mongoLogRepository) List() {
	fmt.Println("i am mongo db list")
}

func (m *mongoLogRepository) Find() {
	fmt.Println("i am mongo db find")
}

func (m *mongoLogRepository) Count() {
	// res, _ := m.collection.Find(context.Background(), filter)
	res, _ := m.collection.CountDocuments(context.Background(), bson.D{})
	fmt.Println(res)
}
