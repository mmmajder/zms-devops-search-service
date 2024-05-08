package persistence

import (
	"context"

	"github.com/mmmajder/devops-search-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "searchdb"
	COLLECTION = "search"
)

type SearchMongoDBStore struct {
	search *mongo.Collection
}

func NewHotelMongoDBStore(client *mongo.Client) domain.SearchStore {
	search := client.Database(DATABASE).Collection(COLLECTION)
	return &SearchMongoDBStore{
		search: search,
	}
}

func (store *SearchMongoDBStore) Get(id primitive.ObjectID) (*domain.Search, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *SearchMongoDBStore) GetAll() ([]*domain.Search, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *SearchMongoDBStore) Insert(search *domain.Search) error {
	result, err := store.search.InsertOne(context.TODO(), search)
	if err != nil {
		return err
	}
	search.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *SearchMongoDBStore) DeleteAll() {
	store.search.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *SearchMongoDBStore) filter(filter interface{}) ([]*domain.Search, error) {
	cursor, err := store.search.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *SearchMongoDBStore) filterOne(filter interface{}) (hotel *domain.Search, err error) {
	result := store.search.FindOne(context.TODO(), filter)
	err = result.Decode(&hotel)
	return
}

func decode(cursor *mongo.Cursor) (searchEntities []*domain.Search, err error) {
	for cursor.Next(context.TODO()) {
		var hotel domain.Search
		err = cursor.Decode(&hotel)
		if err != nil {
			return
		}
		searchEntities = append(searchEntities, &hotel)
	}
	err = cursor.Err()
	return
}
