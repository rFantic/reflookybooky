package handler

import (
	"context"
	"flookybooky/ent"
	"flookybooky/ent/customer"
	"flookybooky/grpc/customer-grpc/internal"
	"flookybooky/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	client ent.Client
}

func NewCustomerHandler(client ent.Client) (*CustomerHandler, error) {
	return &CustomerHandler{
		client: client,
	}, nil
}

func (h *CustomerHandler) GetCustomer(ctx context.Context, req *pb.UUID) (*pb.Customer, error) {
	neededID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	customerRes, err := h.client.Customer.Query().
		Where(customer.ID(neededID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomerEntToPb(customerRes), nil
}

func (h *CustomerHandler) GetCustomers(ctx context.Context, req *pb.Pagination) (*pb.Customers, error) {
	query := h.client.Customer.Query()
	if req != nil {
		var options []customer.OrderOption
		if req.AscFields != nil {
			options = append(options, ent.Asc(req.AscFields...))
		}
		if req.DesFields != nil {
			options = append(options, ent.Desc(req.DesFields...))
		}
		query.Order(options...)
		if req.Limit != nil {
			query.Limit(int(*req.Limit))
		} else {
			query.Limit(10)
		}
		if req.Offset != nil {
			query.Offset(int(*req.Offset))
		}
	} else {
		query.Limit(10)
	}
	customersRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomersEntToPb(customersRes), nil
}

func (h *CustomerHandler) PostCustomer(ctx context.Context, req *pb.CustomerInput) (*pb.Customer, error) {
	customerRes, err := h.client.Customer.Create().
		SetName(req.Name).
		SetAddress(req.Address).
		SetLicenseID(req.LicenseId).
		SetName(req.Name).
		SetPhoneNumber(req.PhoneNumber).
		SetEmail(req.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseCustomerEntToPb(customerRes), err
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, req *pb.CustomerUpdateInput) (*emptypb.Empty, error) {
	_customerId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	query := h.client.Customer.UpdateOneID(_customerId)
	if req.Email != nil {
		query.SetEmail(*req.Email)
	}
	if req.Address != nil {
		query.SetAddress(*req.Address)
	}
	if req.LicenseId != nil {
		query.SetAddress(*req.LicenseId)
	}
	if req.Name != nil {
		query.SetAddress(*req.Name)
	}
	if req.PhoneNumber != nil {
		query.SetAddress(*req.PhoneNumber)
	}
	err = query.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
