package mongoDB

import (
	"app-log/app/repo"
	"fmt"

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
