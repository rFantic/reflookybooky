package resolver

import (
	"context"
	"flookybooky/gql_generated"
	"flookybooky/internal/util"
	"flookybooky/model"
	"flookybooky/pb"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Client struct {
	UserClient     pb.UserServiceClient
	CustomerClient pb.CustomerServiceClient
	FlightClient   pb.FlightServiceClient
	BookingClient  pb.BookingServiceClient
}

type Resolver struct{ client Client }

func NewSchema(client Client) graphql.ExecutableSchema {
	var d = gql_generated.DirectiveRoot{
		HasRoles: func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*model.Role) (interface{}, error) {
			c, _ := ctx.Value(util.ContextKey{}).(*gin.Context)
			tokenString, err := c.Cookie("Authentication")
			if err != nil {
				return nil, fmt.Errorf("missing authentication cookie")
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return util.Secretkey, nil
			})
			if err != nil {
				return nil, fmt.Errorf("can't parse jwt: %w", err)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					return nil, fmt.Errorf("token expired")
				}
				id := claims["sub"].(string)
				res, err := client.UserClient.GetUser(ctx, &pb.UUID{Id: id})
				if err != nil {
					return nil, fmt.Errorf("claims user not found")
				}
				ctxRole := model.Role(res.Role)
				if !Contains(roles, ctxRole) {
					return nil, fmt.Errorf("current role not qualified")
				}
			} else {
				return nil, err
			}
			return next(ctx)
		},
	}

	return gql_generated.NewExecutableSchema(gql_generated.Config{
		Resolvers:  &Resolver{client},
		Directives: d,
	})
}

func Contains(roles []*model.Role, role model.Role) bool {
	for _, r := range roles {
		if r.String() == role.String() {
			return true
		}
	}
	return false
}
