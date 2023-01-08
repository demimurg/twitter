package tests

import "github.com/demimurg/twitter/pkg/proto"

// TestAuth for basic operations, subtests can't be run separate
func (s *endToEndTestSuite) TestAuth() {
	var (
		gretaID      int64
		gretaProfile = &proto.UserProfile{
			FullName:    "Greta Thunberg",
			Email:       "smalldickenergy@getalife.com",
			Caption:     "How dare you are?🤬",
			DateOfBirth: date(2003, 01, 03),
		}
	)

	s.Run("register greta thunberg", func() {
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{User: gretaProfile})
		s.NoError(err)
		gretaID = resp.UserId
	})

	s.Run("can't register second time", func() {
		s.T().Skip("should fix repeatable registration")
		_, err := s.cli.Register(ctx, &proto.RegisterRequest{User: gretaProfile})
		s.Error(err)
		s.T().Log("error on second time registration:", err)
	})

	s.Run("greta trying to login", func() {
		resp, err := s.cli.Login(ctx, &proto.LoginRequest{
			Email: gretaProfile.Email,
		})
		s.Require().NoError(err)

		s.Equal(gretaID, resp.UserId)
		s.Equal(gretaProfile.FullName, resp.UserProfile.FullName)
		s.Equal(gretaProfile.Email, resp.UserProfile.Email)
		s.Equal(gretaProfile.Caption, resp.UserProfile.Caption)
		s.Equal(gretaProfile.DateOfBirth.AsTime(), resp.UserProfile.DateOfBirth.AsTime())
	})
}
