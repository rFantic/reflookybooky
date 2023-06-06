package internal

import (
	"flookybooky/ent"
	"flookybooky/pb"

	"github.com/jinzhu/copier"
)

func ParseCustomerEntToPb(in *ent.Customer) (out *pb.Customer) {
	out = &pb.Customer{}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
		out.LicenseId = in.LicenseID
		out.CreatedAt = in.CreatedAt.String()
		out.UpdatedAt = in.CreatedAt.String()
	}
	return out
}

func ParseCustomersEntToPb(in []*ent.Customer) (out *pb.Customers) {
	out = &pb.Customers{
		Customers: make([]*pb.Customer, len(in)),
	}
	for i, a := range in {
		out.Customers[i] = ParseCustomerEntToPb(a)
	}
	return out
}
