package internal

import (
	"flookybooky/ent"
	"flookybooky/pb"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func ParseTicketEntToPb(in *ent.Ticket) (out *pb.Ticket) {
	if in == nil {
		return nil
	}
	out = &pb.Ticket{
		Id:                 in.ID.String(),
		PassengerLicenseId: in.PassengerLicenseID,
		CreatedAt:          in.CreatedAt.String(),
		UpdatedAt:          in.CreatedAt.String(),
	}

	copier.Copy(&out, in)
	return out
}

func ParseTicketsEntToPb(in []*ent.Ticket) (out *pb.Tickets) {
	if in == nil {
		return nil
	}
	out = &pb.Tickets{
		Tickets: make([]*pb.Ticket, len(in)),
	}
	copier.Copy(&out, in)
	for i, a := range in {
		out.Tickets[i] = ParseTicketEntToPb(a)
	}
	return out
}

func ParseBookingEntToPb(in *ent.Booking) (out *pb.Booking) {
	if in == nil {
		return nil
	}
	out = &pb.Booking{
		Id: in.ID.String(),
		GoingFlight: &pb.Flight{
			Id: in.GoingFlightID.String(),
		},
		CreatedAt: in.CreatedAt.String(),
		UpdatedAt: in.CreatedAt.String(),
	}
	copier.Copy(&out, in)
	out.Customer = &pb.Customer{
		Id: in.CustomerID.String(),
	}
	if in.ReturnFlightID != uuid.Nil {
		out.ReturnFlight = &pb.Flight{
			Id: in.ReturnFlightID.String(),
		}
	}
	return out
}

func ParseBookingsEntToPb(in []*ent.Booking) (out *pb.Bookings) {
	out = &pb.Bookings{
		Bookings: make([]*pb.Booking, len(in)),
	}
	for i, a := range in {
		out.Bookings[i] = ParseBookingEntToPb(a)
	}
	return out
}
