package main

import (
	"context"
	"flookybooky/ent"
	"flookybooky/grpc/booking-grpc/handler"
	"flookybooky/pb"
	"net"

	"entgo.io/ent/dialect"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile("./config/.env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func main() {
	servicesClient := servicesConn()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := logger.Sugar()
	POSTGRES_URI := string(viper.GetString("POSTGRES_URI"))
	client, err := ent.Open(dialect.Postgres, POSTGRES_URI)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	listen, err := net.Listen("tcp", ":2220")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	h, err := handler.NewBookingHandler(*client,
		servicesClient.CustomerClient,
		servicesClient.FlightClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterBookingServiceServer(s, h)
	s.Serve(listen)
}

type ServicesClient struct {
	CustomerClient pb.CustomerServiceClient
	FlightClient   pb.FlightServiceClient
}

func servicesConn() ServicesClient {
	customerConn, err := grpc.Dial("customer:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	customerClient := pb.NewCustomerServiceClient(customerConn)
	flightConn, err := grpc.Dial("flight:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	flightClient := pb.NewFlightServiceClient(flightConn)
	return ServicesClient{
		CustomerClient: customerClient,
		FlightClient:   flightClient,
	}
}
