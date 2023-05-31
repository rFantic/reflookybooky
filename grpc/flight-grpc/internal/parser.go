package internal

import (
	"flookybooky/ent"
	"flookybooky/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ParseAirportEntToPb(in *ent.Airport) (out *pb.Airport) {
	out = &pb.Airport{}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
	}
	return out
}

func ParseAirportsEntToPb(in []*ent.Airport) (out *pb.Airports) {
	out = &pb.Airports{
		Airports: make([]*pb.Airport, len(in)),
	}
	for i, a := range in {
		out.Airports[i] = ParseAirportEntToPb(a)
	}
	return out
}

func ParseAirportPbToEnt(in *pb.Airport) (out *ent.Airport, err error) {
	out = &ent.Airport{}
	copier.Copy(&out, in)
	if in != nil {
		out.ID, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ParseFlightEntToPb(in *ent.Flight) (out *pb.Flight) {
	out = &pb.Flight{}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
		out.Origin = &pb.Airport{Id: in.OriginID.String()}
		out.Destination = &pb.Airport{Id: in.DestinartionID.String()}
		out.ArrivalTime = timestamppb.New(in.ArrivalTime)
		out.DepartureTime = timestamppb.New(in.DepartureTime)
	}
	return out
}

func ParseFlightsEntToPb(in []*ent.Flight) (out *pb.Flights) {
	out = &pb.Flights{
		Flights: make([]*pb.Flight, len(in)),
	}
	for i, a := range in {
		out.Flights[i] = ParseFlightEntToPb(a)
	}
	return out
}

// func ParseSeatEntToPb(in *ent.Seat) (out *pb.Seat) {
// 	out = &pb.Seat{}
// 	copier.Copy(&out, in)
// 	out.Id = in.ID.String()
// 	return out
// }

// func ParseSeatsEntToPb(in []*ent.Seat) (out *pb.Seats) {
// 	out = &pb.Seats{
// 		Seats: make([]*pb.Seat, len(in)),
// 	}
// 	for i, a := range in {
// 		out.Seats[i] = ParseSeatEntToPb(a)
// 	}
// 	return out
// }
