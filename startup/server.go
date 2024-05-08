package startup

import (

	// "github.com/tamararankovic/microservices_demo/catalogue_service/application"
	// "github.com/tamararankovic/microservices_demo/catalogue_service/domain"
	// "github.com/tamararankovic/microservices_demo/catalogue_service/infrastructure/api"
	// "github.com/tamararankovic/microservices_demo/catalogue_service/infrastructure/persistence"
	// "github.com/tamararankovic/microservices_demo/catalogue_service/startup/config"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mmmajder/devops-search-service/application"
	"github.com/mmmajder/devops-search-service/domain"
	"github.com/mmmajder/devops-search-service/infrastructure/api"
	"github.com/mmmajder/devops-search-service/infrastructure/persistence"
	"github.com/mmmajder/devops-search-service/startup/config"
	"net/http"

	// catalogue "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	searchStore := server.initSearchStore(mongoClient)
	searchService := server.initSearchService(searchStore)

	searchHandler := server.initSearchHandler(searchService)
	searchHandler.Init(server.mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.SearchDBUsername, server.config.SearchDBPassword, server.config.SearchDBHost, server.config.SearchDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initSearchStore(client *mongo.Client) domain.SearchStore {
	store := persistence.NewHotelMongoDBStore(client)
	store.DeleteAll()
	for _, searchEntity := range searches {
		err := store.Insert(searchEntity)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initSearchService(store domain.SearchStore) *application.SearchService {
	return application.NewSearchService(store)
}

func (server *Server) initSearchHandler(service *application.SearchService) *api.SearchHandler {
	bookingEndpoint := fmt.Sprintf("%s:%s", server.config.BookingHost, server.config.BookingPort)
	return api.NewSearchHandler(service, bookingEndpoint)
}
