package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"common/auth"
	"common/connection"
	"common/hashing"
	commonTesting "common/testing"

	"account/model"
	"account/option"
	proto "account/proto/account"
	"account/repository"
)

type authHandlerSuite struct {
	commonTesting.Suite
	handler     proto.AuthHandler
	userHandler proto.UserHandler
}

func (suite *authHandlerSuite) SetupSuite() {
	r := suite.Require()

	suite.Init()

	suite.DeleteDatabaseFile("account")

	db, err := connection.GetDatabaseConnection("account")
	r.Nilf(err, "error getting database connection: %v", err)

	suite.MigrateModels(
		"account",
		&model.User{},
		&model.UserToken{},
	)

	authStrategy, err := auth.NewAuthStrategy()
	if err != nil {
		panic(err)
	}

	hashingStrategy := hashing.NewHashingStrategy(hashing.BCryptHashingStrategy, nil)

	userRepository := repository.NewUserRepository(db)
	userTokenRepository := repository.NewUserTokenRepository(db)

	options := option.NewOptions(
		option.WithUserRepository(userRepository),
		option.WithUserTokenRepository(userTokenRepository),
		option.WithHashingStrategy(hashingStrategy),
		option.WithAuthStrategy(authStrategy),
	)

	suite.handler = NewAuthHandler(options)
	suite.userHandler = NewUserHandler(options)
}

func (suite *authHandlerSuite) SetupTest() {
	suite.TruncateTables(
		"account",
		"users",
		"user_tokens",
	)
}

func (suite *authHandlerSuite) TearDownSuite() {
	suite.DeleteDatabaseFile("account")
}

func (suite *authHandlerSuite) TestCreateToken() {
	r := suite.Require()

	ctx := context.TODO()

	_, err := createDummyUser(
		suite.userHandler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)
	r.Nilf(err, "error creating user: %v", err)

	res, err := createDummyUserToken(
		suite.handler,
		ctx,
		"buddy@example.com",
		"123456",
	)
	r.Nilf(err, "error creating token: %v", err)

	r.NotNil(res.Token)
	r.NotNil(res.User)
	r.Equal("buddy@example.com", res.User.Email)
}

func (suite *authHandlerSuite) TestVerifyToken() {
	r := suite.Require()

	ctx := context.TODO()

	_, err := createDummyUser(
		suite.userHandler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)
	r.Nilf(err, "error creating user: %v", err)

	authResponse, err := createDummyUserToken(
		suite.handler,
		ctx,
		"buddy@example.com",
		"123456",
	)
	r.Nilf(err, "error creating token: %v", err)

	req := &proto.VerifyTokenRequest{
		Token: authResponse.Token,
	}
	res := &proto.VerifyTokenResponse{}
	err = suite.handler.VerifyToken(ctx, req, res)
	r.Nilf(err, "error verifying token: %v", err)
}

func (suite *authHandlerSuite) TestRevokeToken() {
	r := suite.Require()

	ctx := context.TODO()

	_, err := createDummyUser(
		suite.userHandler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)
	r.Nilf(err, "error creating user: %v", err)

	authResponse, err := createDummyUserToken(
		suite.handler,
		ctx,
		"buddy@example.com",
		"123456",
	)
	r.Nilf(err, "error creating token: %v", err)

	req := &proto.RevokeTokenRequest{
		Token: authResponse.Token,
	}
	res := &proto.RevokeTokenResponse{}
	err = suite.handler.RevokeToken(ctx, req, res)
	r.Nilf(err, "error revoking token: %v", err)

	verifyReq := &proto.VerifyTokenRequest{
		Token: authResponse.Token,
	}
	verifyRes := &proto.VerifyTokenResponse{}
	err = suite.handler.VerifyToken(ctx, verifyReq, verifyRes)
	r.NotNilf(err, "error should not be nil")
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, &authHandlerSuite{})
}

func createDummyUserToken(
	authHandler proto.AuthHandler,
	ctx context.Context,
	email string,
	password string) (res *proto.AuthResponse, err error) {
	req := &proto.AuthRequest{
		Email:    email,
		Password: password,
	}
	res = &proto.AuthResponse{}
	err = authHandler.CreateToken(ctx, req, res)
	return
}
