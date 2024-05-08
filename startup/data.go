package startup

import (
	"github.com/mmmajder/devops-search-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var searches = []*domain.Search{
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f82"),
		Name: "name",
	},
	{
		Id:   getObjectId("623b0cc3a34d25d8567f9f83"),
		Name: "name2",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
