package handler

import (
	"context"
	"flookybooky/ent"
	"flookybooky/ent/user"
	"flookybooky/grpc/user-grpc/internal"
	"flookybooky/internal/util"
	"flookybooky/pb"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	client ent.Client
}

func NewUserHandler(client ent.Client) (*UserHandler, error) {
	return &UserHandler{
		client: client,
	}, nil
}

func (h *UserHandler) PostUser(ctx context.Context, req *pb.UserInput) (*pb.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, fmt.Errorf("generate hash: %w", err)
	}
	query := h.client.User.Create().
		SetUsername(req.Username).
		SetPassword(string(hash)).
		SetEmail(req.Email).
		SetRole(user.Role(req.Role))
	if req.CustomerId != nil {
		query.SetCustomerID(*req.CustomerId)
	}
	userRes, err := query.Save(ctx)
	return internal.ParseUserEntToPb(userRes), err
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UUID) (*pb.User, error) {
	_uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	userRes, err := h.client.User.Get(ctx, _uuid)
	return internal.ParseUserEntToPb(userRes), err
}

func (h *UserHandler) GetUsers(ctx context.Context, req *pb.Pagination) (*pb.Users, error) {
	query := h.client.User.Query()
	if req != nil {
		var options []user.OrderOption
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
	usersRes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return internal.ParseUsersEntToPb(usersRes), nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.client.User.Query().
		Where(user.Username(req.User.Username)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("wrong user name")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.User.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expireTime,
	})
	tokenString, err := token.SignedString(util.Secretkey)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtToken:   tokenString,
		ExpireTime: expireTime,
	}, err
}

func (h *UserHandler) UpdatePassword(ctx context.Context, req *pb.PasswordUpdateInput) (*emptypb.Empty, error) {
	_userId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_user, err := h.client.User.Get(ctx, _userId)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(req.PreviousPassword))
	if err != nil {
		return nil, fmt.Errorf("wrong old password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		return nil, fmt.Errorf("generate hash: %w", err)
	}

	err = h.client.User.UpdateOneID(_userId).SetPassword(string(hash)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UserUpdateInput) (*emptypb.Empty, error) {
	_userId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	query := h.client.User.UpdateOneID(_userId)
	if req.Email != nil {
		query.SetEmail(*req.Email)
	}
	if req.Role != nil {
		query.SetRole(user.Role(*req.Role))
	}
	err = query.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
