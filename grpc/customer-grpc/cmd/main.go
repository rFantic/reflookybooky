package main

import (
	"context"
	"flookybooky/ent"
	"flookybooky/grpc/customer-grpc/handler"
	"net"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"flookybooky/pb"
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	logger, err := zap.NewProduction()
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
	h, err := handler.NewCustomerHandler(*client)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)
	s.Serve(listen)
}
