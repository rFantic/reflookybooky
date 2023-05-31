package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"flookybooky/gql_generated"
	internal "flookybooky/internal/parser"
	"flookybooky/model"
	"flookybooky/pb"
	"fmt"
)

// Origin is the resolver for the origin field.
func (r *flightResolver) Origin(ctx context.Context, obj *model.Flight) (*model.Airport, error) {
	var out *model.Airport
	if obj.Origin != nil {
		originRes, err := r.client.FlightClient.GetAirport(ctx, &pb.UUID{Id: obj.Origin.ID})
		if err != nil {
			return nil, err
		}
		out = internal.ParseAirportPbToGraphql(originRes)
	}
	return out, nil
}

// Destination is the resolver for the destination field.
func (r *flightResolver) Destination(ctx context.Context, obj *model.Flight) (*model.Airport, error) {
	var out *model.Airport
	if obj.Destination != nil {
		destinationRes, err := r.client.FlightClient.GetAirport(ctx, &pb.UUID{Id: obj.Destination.ID})
		if err != nil {
			return nil, err
		}
		out = internal.ParseAirportPbToGraphql(destinationRes)
	}
	return out, nil
}

// CreateFlight is the resolver for the createFlight field.
func (r *flightOpsResolver) CreateFlight(ctx context.Context, obj *model.FlightOps, input model.FlightInput) (*model.Flight, error) {
	flightReq, err := internal.ParseFlightInputGraphqlToPb(&input)
	if err != nil {
		return nil, err
	}
	flightRes, err := r.client.FlightClient.PostFlight(ctx, flightReq)
	return internal.ParseFlightPbToGraphql(flightRes), err
}

// UpdateFlight is the resolver for the updateFlight field.
func (r *flightOpsResolver) UpdateFlight(ctx context.Context, obj *model.FlightOps, input model.FlightUpdateInput) (bool, error) {
	updateInput, err := internal.ParseFlightUpdateInputGraphqlToPb(&input)
	if err != nil {
		return false, err
	}
	_, err = r.client.FlightClient.UpdateFlight(ctx, updateInput)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CancelFlight is the resolver for the cancelFlight field.
func (r *flightOpsResolver) CancelFlight(ctx context.Context, obj *model.FlightOps, input *model.FlightCancelInput) (bool, error) {
	if input != nil {
		_, err := r.client.FlightClient.CancelFlight(ctx, &pb.UUID{Id: input.ID})
		if err != nil {
			return false, err
		}
		_, err = r.client.BookingClient.CancelBookingOfFlight(ctx, &pb.UUID{Id: input.ID})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, fmt.Errorf("missing input")
}

// Flight is the resolver for the flight field.
func (r *mutationResolver) Flight(ctx context.Context) (*model.FlightOps, error) {
	return &model.FlightOps{}, nil
}

// Flight is the resolver for the flight field.
func (r *queryResolver) Flight(ctx context.Context, input *model.Pagination) ([]*model.Flight, error) {
	flightsRes, err := r.client.FlightClient.GetFlights(ctx,
		internal.ParsePaginationGraphqlToPb(input))
	return internal.ParseFlightsPbToGraphql(flightsRes), err
}

// SearchFlight is the resolver for the searchFlight field.
func (r *queryResolver) SearchFlight(ctx context.Context, input *model.FlightSearchInput) ([]*model.Flight, error) {
	flightsRes, err := r.client.FlightClient.SearchFlight(ctx, internal.ParseFlightSearchInputGraphqlToPb(input))
	return internal.ParseFlightsPbToGraphql(flightsRes), err
}

// Flight returns gql_generated.FlightResolver implementation.
func (r *Resolver) Flight() gql_generated.FlightResolver { return &flightResolver{r} }

// FlightOps returns gql_generated.FlightOpsResolver implementation.
func (r *Resolver) FlightOps() gql_generated.FlightOpsResolver { return &flightOpsResolver{r} }

type flightResolver struct{ *Resolver }
type flightOpsResolver struct{ *Resolver }