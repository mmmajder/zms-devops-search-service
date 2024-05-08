package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Search struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
