package model

type Test struct {
	Id    string `bson:"_id"`
	Name  string `bson:"name"`
	Level int    `bson:"level"`
}
