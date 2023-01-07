package grpcsrv

import (
	"context"
	"errors"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *twitter) Register(ctx context.Context, req *proto.UserProfile) (*proto.RegisterResponse, error) {
	user, err := t.ur.Register(ctx, req.FullName, req.Email, req.Caption, req.DateOfBirth.AsTime())
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			err = status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &proto.RegisterResponse{UserId: int64(user.ID)}, nil
}

func (t *twitter) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := t.ur.Login(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		UserId:      int64(user.ID),
		UserProfile: convertToUserProfile(user),
	}, nil
}

func convertToUserProfile(user *entity.User) *proto.UserProfile {
	return &proto.UserProfile{
		Email:       user.Email,
		FullName:    user.FullName,
		Caption:     user.Caption,
		DateOfBirth: timestamppb.New(user.BirthDate),
	}
}
