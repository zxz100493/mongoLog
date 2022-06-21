package mongoDb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoLogRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(c *mongo.Collection) *mongoLogRepository {
	return &mongoLogRepository{c}
}

func (m *mongoLogRepository) List() {
	fmt.Println("i am mongo list")
}

func (m *mongoLogRepository) Find() {
	fmt.Println("i am mongo find")
}
