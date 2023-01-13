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

func (t *twitter) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	user, err := t.up.Register(
		ctx, req.User.FullName, req.User.Email,
		req.User.Caption, req.User.DateOfBirth.AsTime(),
	)
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			err = status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &proto.RegisterResponse{UserId: int64(user.ID)}, nil
}

func (t *twitter) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := t.up.Login(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		UserId:      int64(user.ID),
		UserProfile: convertToUserProfile(user),
	}, nil
}

func (t *twitter) UpdateCaption(ctx context.Context, req *proto.UpdateCaptionRequest) (*proto.UpdateCaptionResponse, error) {
    err := t.up.UpdateCaption(ctx, int(req.UserId), req.NewCaption)
    if err != nil {
        return nil, err
    }

    return &proto.UpdateCaptionResponse{}, nil
}

func convertToUserProfile(user *entity.User) *proto.UserProfile {
	return &proto.UserProfile{
		Email:       user.Email,
		FullName:    user.FullName,
		Caption:     user.Caption,
		DateOfBirth: timestamppb.New(user.BirthDate),
	}
}
