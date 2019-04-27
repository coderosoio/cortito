package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"common/connection"
	"common/hashing"
	commonTesting "common/testing"

	proto "account/proto/account"

	"account/model"
	"account/option"
	"account/repository"
)

type userHandlerSuite struct {
	commonTesting.Suite
	handler proto.UserHandler
}

func (suite *userHandlerSuite) SetupSuite() {
	r := suite.Require()

	suite.Init()

	suite.DeleteDatabaseFile("account")

	db, err := connection.GetDatabaseConnection("account")
	r.Nilf(err, "error getting database connection: %v", err)

	suite.MigrateModels(
		"account",
		&model.User{},
	)

	hashingStrategy := hashing.NewHashingStrategy(hashing.BCryptHashingStrategy, nil)

	userRepository := repository.NewUserRepository(db)

	options := option.NewOptions(
		option.WithUserRepository(userRepository),
		option.WithHashingStrategy(hashingStrategy),
	)

	suite.handler = NewUserHandler(options)
}

func (suite *userHandlerSuite) SetupTest() {
	suite.TruncateTables(
		"account",
		"users",
	)
}

func (suite *userHandlerSuite) TearDownSuite() {
	suite.DeleteDatabaseFile("account")
}

func (suite *userHandlerSuite) TestCreateUser() {
	r := suite.Require()

	ctx := context.TODO()

	res, err := createDummyUser(
		suite.handler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)

	r.Nilf(err, "error creating user: %v", err)

	r.NotZero(res.Id)
	r.Equal("Buddy Tester", res.Name)
	r.Equal("buddy@example.com", res.Email)
}

func (suite *userHandlerSuite) TestUpdateUser() {
	r := suite.Require()

	ctx := context.TODO()

	res, err := createDummyUser(
		suite.handler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)
	r.Nilf(err, "error creating user: %v", err)

	newName := "Another Name"
	id := res.Id

	req := &proto.UserRequest{
		Id:   id,
		Name: newName,
	}
	err = suite.handler.UpdateUser(ctx, req, res)
	r.Nilf(err, "error updating user: %v", err)

	r.NotZero(res.Id)
	r.Equal(res.Id, id)
	r.Equal(res.Name, newName)
	r.NotNil(res.Email)
}

func (suite *userHandlerSuite) TestFindUser() {
	r := suite.Require()

	ctx := context.TODO()

	res, err := createDummyUser(
		suite.handler,
		ctx,
		"Buddy Tester",
		"buddy@example.com",
		"123456",
		"123456",
	)
	r.Nilf(err, "error creating user: %v", err)

	id := res.Id

	req := &proto.UserRequest{
		Id: id,
	}
	res = &proto.UserResponse{}
	err = suite.handler.FindUser(ctx, req, res)
	r.Nilf(err, "error finding user: %v", err)

	r.NotZero(res.Id)
	r.Equal(res.Id, id)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, &userHandlerSuite{})
}

func createDummyUser(
	userHandler proto.UserHandler,
	ctx context.Context,
	name string,
	email string,
	password string,
	passwordConfirmation string) (res *proto.UserResponse, err error) {
	req := &proto.UserRequest{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}
	res = &proto.UserResponse{}

	err = userHandler.CreateUser(ctx, req, res)
	return
}
