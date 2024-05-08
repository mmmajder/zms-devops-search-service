package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchStore interface {
	Get(id primitive.ObjectID) (*Search, error)
	GetAll() ([]*Search, error)
	Insert(hotel *Search) error
	DeleteAll()
}
