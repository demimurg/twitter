package grpcsrv

import (
	"context"
	"errors"

	"github.com/demimurg/twitter/internal/entity"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/log"
	"github.com/demimurg/twitter/pkg/proto"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (t *twitter) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	user, err := t.up.Register(
		ctx, req.User.FullName, req.User.Email, req.Password,
		req.User.Caption, req.User.DateOfBirth.AsTime(),
	)
	if err != nil {
		if errors.Is(err, usecase.ErrFakeEmail) {
			err = status.Error(codes.PermissionDenied, err.Error())
		} else if errors.Is(err, usecase.ErrUserExists) {
			err = status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, err
	}
	return &proto.RegisterResponse{
		UserId: int64(user.ID),
		Jwt:    t.createJWT(user),
	}, nil
}

func (t *twitter) createJWT(user *entity.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
	jwtToken, err := token.SignedString([]byte(t.secret))
	if err != nil {
		log.Error(context.Background(), "can't sign jwt token!", "error", err)
	}
	return jwtToken
}

func (t *twitter) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := t.up.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrWrongCredentials) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return &proto.LoginResponse{
		UserId:      int64(user.ID),
		UserProfile: convertToUserProfile(user),
		Jwt:         t.createJWT(user),
	}, nil
}

func (t *twitter) UpdateCaption(ctx context.Context, req *proto.UpdateCaptionRequest) (*proto.UpdateCaptionResponse, error) {
	err := t.up.UpdateCaption(ctx, getUserID(ctx), req.NewCaption)
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
