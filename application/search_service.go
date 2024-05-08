package application

import (
	"github.com/mmmajder/devops-search-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchService struct {
	store domain.SearchStore
}

func NewSearchService(store domain.SearchStore) *SearchService {
	return &SearchService{
		store: store,
	}
}

func (service *SearchService) Get(id primitive.ObjectID) (*domain.Search, error) {
	return service.store.Get(id)
}

func (service *SearchService) GetAll() ([]*domain.Search, error) {
	return service.store.GetAll()
}
