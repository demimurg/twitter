//go:build e2e

package tests

import (
	"github.com/demimurg/twitter/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestAuth for basic operations, subtests can't be run separate
func (s *endToEndTestSuite) TestAuth() {
	var (
		gretaID       int64
		gretaJWT      string
		gretaPassword = "pony-girl2003"
		gretaProfile  = &proto.UserProfile{
			FullName:    "Greta Thunberg",
			Email:       "smalldickenergy@getalife.com",
			Caption:     "How dare you are?🤬",
			DateOfBirth: date(2003, 01, 03),
		}
	)

	s.Run("register greta thunberg", func() {
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{
			User: gretaProfile, Password: gretaPassword,
		})
		s.NoError(err)
		s.Require().NotNil(resp)
		gretaID = resp.UserId
		gretaJWT = resp.Jwt
	})

	s.Run("can't register second time", func() {
		_, err := s.cli.Register(ctx, &proto.RegisterRequest{User: gretaProfile, Password: gretaPassword})
		s.Error(err)
		s.T().Log("error on second time registration:", err)
	})

	s.Run("greta trying to login", func() {
		resp, err := s.cli.Login(ctx, &proto.LoginRequest{
			Email: gretaProfile.Email, Password: gretaPassword,
		})
		s.Require().NoError(err)

		s.Equal(gretaID, resp.UserId, "check greta id")
		s.Equal(gretaJWT, resp.Jwt, "check jwt token")
		s.EqualProto(gretaProfile, resp.UserProfile, "check greta profile")
	})

	s.Run("greta trying to login with wrong password", func() {
		resp, err := s.cli.Login(ctx, &proto.LoginRequest{
			Email: gretaProfile.Email, Password: gretaPassword + "typo",
		})
		s.Assert().Nil(resp)
		s.Require().Error(err)
		s.Require().True(status.Code(err) == codes.NotFound, "must return not found")
	})

	s.Run("greta set new caption for profile", func() {
		ctx := withToken(ctx, gretaJWT)
		callForChaos := "We can no longer save the world by playing by the rules"
		_, err := s.cli.UpdateCaption(ctx, &proto.UpdateCaptionRequest{
			NewCaption: callForChaos,
		})
		s.Require().NoError(err)

		resp, err := s.cli.Login(ctx, &proto.LoginRequest{
			Email: gretaProfile.Email, Password: gretaPassword,
		})
		s.Require().NoError(err)
		s.Equal(callForChaos, resp.UserProfile.Caption)
	})

	s.Run("donald trump can't register, no more", func() {
		_, err := s.cli.Register(ctx, &proto.RegisterRequest{
			User: &proto.UserProfile{
				FullName:    "Donald Trump",
				Email:       "donald.trump@inpresident.com",
				Caption:     "Meet me at the Capitol",
				DateOfBirth: date(1946, 06, 14),
			},
			Password: "guilty-pleasure",
		})
		s.True(status.Code(err) == codes.PermissionDenied, "permission must be denied for trump")
	})
}
