package api

import (
	"context"
	"encoding/json"
	"fmt"
	booking "github.com/mmmajder/devops-booking-service/proto"
	"github.com/mmmajder/devops-search-service/infrastructure/services"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mmmajder/devops-search-service/application"
)

type SearchHandler struct {
	service              *application.SearchService
	bookingClientAddress string
}

func NewSearchHandler(service *application.SearchService, bookingClientAddress string) *SearchHandler {
	return &SearchHandler{
		service:              service,
		bookingClientAddress: bookingClientAddress,
	}
}

func (handler *SearchHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/search/proba", handler.GetAll)
	err = mux.HandlePath("GET", "/search", handler.GetAll)
	err = mux.HandlePath("GET", "/", handler.GetAll)
	if err != nil {
		panic(err)
	}
}

func (handler *SearchHandler) GetAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hotel, err := handler.service.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bookingInfo, err := handler.getAllHotelsFromBooking()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		for _, item := range bookingInfo.GetHotels() {
			fmt.Println(item.Name)
		}
	}

	response, err := json.Marshal(hotel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *SearchHandler) getAllHotelsFromBooking() (*booking.GetAllResponse, error) {
	bookingClient := services.NewBookingClient(handler.bookingClientAddress)
	return bookingClient.GetAll(context.TODO(), &booking.GetAllRequest{})
}
