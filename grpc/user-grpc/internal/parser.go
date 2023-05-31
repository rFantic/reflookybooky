package internal

import (
	"flookybooky/ent"
	"flookybooky/pb"

	"github.com/jinzhu/copier"
)

func ParseUserEntToPb(in *ent.User) (out *pb.User) {
	out = &pb.User{}
	copier.Copy(&out, in)
	if in != nil {
		out.Id = in.ID.String()
		if in.CustomerID != nil {
			out.Customer = &pb.Customer{
				Id: *in.CustomerID,
			}
		}
	}
	return out
}

func ParseUsersEntToPb(in []*ent.User) (out *pb.Users) {
	out = &pb.Users{
		Users: make([]*pb.User, len(in)),
	}
	for i, a := range in {
		out.Users[i] = ParseUserEntToPb(a)
	}
	return out
}
