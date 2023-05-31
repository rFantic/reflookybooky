package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/gql_generated"
	internal "flookybooky/internal/parser"
	"flookybooky/model"
)

// CreateAirport is the resolver for the createAirport field.
func (r *airportOpsResolver) CreateAirport(ctx context.Context, obj *model.AirportOps, input model.AirportInput) (*model.Airport, error) {
	airportRes, err := r.client.FlightClient.PostAirport(ctx, internal.ParseAirportInputGraphqlToPb(&input))
	if err != nil {
		return nil, err
	}
	return internal.ParseAirportPbToGraphql(airportRes), nil
}

// Airport is the resolver for the airport field.
func (r *mutationResolver) Airport(ctx context.Context) (*model.AirportOps, error) {
	return &model.AirportOps{}, nil
}

// Airport is the resolver for the airport field.
func (r *queryResolver) Airport(ctx context.Context, input *model.Pagination) ([]*model.Airport, error) {
	airportsRes, err := r.client.FlightClient.GetAirports(ctx,
		internal.ParsePaginationGraphqlToPb(input))
	return internal.ParseAirportsPbToGraphql(airportsRes), err
}

// AirportOps returns gql_generated.AirportOpsResolver implementation.
func (r *Resolver) AirportOps() gql_generated.AirportOpsResolver { return &airportOpsResolver{r} }

type airportOpsResolver struct{ *Resolver }
