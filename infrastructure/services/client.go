package services

import (
	booking "github.com/mmmajder/devops-booking-service/proto"
	//"github.com/mmmajder/devops-search-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewBookingClient(address string) booking.BookingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return booking.NewBookingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
