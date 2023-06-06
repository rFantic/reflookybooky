package main

import (
	"context"
	"flookybooky/ent"
	"flookybooky/grpc/flight-grpc/handler"
	"flookybooky/pb"
	"net"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func init() {
	viper.SetConfigFile("./config/.env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func main() {
	servicesClient := servicesConn()
	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}
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
	h, err := handler.NewFlightHandler(*client,
		servicesClient.BookingClient)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterFlightServiceServer(s, h)
	s.Serve(listen)
}

type ServicesClient struct {
	BookingClient pb.BookingServiceClient
}

func servicesConn() ServicesClient {
	bookingConn, err := grpc.Dial("booking:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	bookingClient := pb.NewBookingServiceClient(bookingConn)
	return ServicesClient{
		BookingClient: bookingClient,
	}
}
