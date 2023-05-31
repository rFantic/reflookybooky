//go:generate go run github.com/99designs/gqlgen
package main

import (
	"flookybooky/middleware"
	"flookybooky/pb"
	"flookybooky/resolver"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
}

func main() {
	client := servicesConn()
	h := handler.NewDefaultServer(resolver.NewSchema(client))
	p := playground.Handler("App", "/graphql")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(
		middleware.CorsMiddleware(),
		// middleware.RequestCtxMiddleware(),
		middleware.CookieMiddleware(),
	)

	// Create a new GraphQL
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.OPTIONS("/graphql", func(c *gin.Context) {
		c.Status(200)
	})

	// Enable playground for query/testing
	r.GET("/", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})

	log.Fatal(r.Run(":8081"))
}

func servicesConn() resolver.Client {
	userConn, err := grpc.Dial("user:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient := pb.NewUserServiceClient(userConn)
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
	bookingConn, err := grpc.Dial("booking:2220", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	bookingClient := pb.NewBookingServiceClient(bookingConn)
	return resolver.Client{
		UserClient:     userClient,
		CustomerClient: customerClient,
		FlightClient:   flightClient,
		BookingClient:  bookingClient,
	}
}
