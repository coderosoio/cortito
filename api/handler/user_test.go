package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/suite"

	commonTesting "common/testing"

	accountProto "account/proto/account"
	"api/option"
)

const (
	userCreatePath = "/api/account/users/"
	userUpdatePath = "/api/account/users/%d"
	userFindPath   = userUpdatePath
)

type userHandlerSuite struct {
	commonTesting.HTTPSuite
}

func (suite *userHandlerSuite) SetupSuite() {
	suite.Init()

	accountService := suite.Config.Service["account"]

	userService := accountProto.NewUserService(
		accountService.URL(),
		suite.ServiceClient,
	)

	options := option.NewOptions(
		option.WithUserService(userService),
	)
	suite.Router = NewRouter(options)
}

func (suite *userHandlerSuite) TestCreateUser() {
	r := suite.Require()

	existingUsers := suite.createUsers(100)
	for _, user := range existingUsers {
		password := gofakeit.Password(true, true, true, true, false, 6)
		userRequest := &accountProto.UserRequest{
			Name:                 user.Name,
			Email:                user.Email,
			Password:             password,
			PasswordConfirmation: password,
		}
		suite.T().Logf("Creating user again: %+v", userRequest)
		w, err := suite.Request(userCreatePath, http.MethodPost, userRequest, nil)
		r.Nilf(err, "error making request: %v", err)

		r.Equal(w.Code, http.StatusUnprocessableEntity)

		suite.T().Logf("response: %+v", w)
	}
}

func (suite *userHandlerSuite) TestUpdateUser() {
	r := suite.Require()

	existingUsers := suite.createUsers(100)
	for _, user := range existingUsers {
		person := gofakeit.Person()
		newName := fmt.Sprintf("%s %s", person.FirstName, person.LastName)

		userRequest := accountProto.UserRequest{
			Name: newName,
		}
		path := fmt.Sprintf(userUpdatePath, user.Id)
		suite.T().Logf("Updating user name %s => %s", user.Name, newName)
		w, err := suite.Request(path, http.MethodPut, userRequest, nil)
		r.Nilf(err, "error updating user: %v", err)
		r.Equal(w.Code, http.StatusOK)

		userResponse := &accountProto.UserResponse{}
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		r.Nilf(err, "error parsing response: %v", err)

		r.Equal(newName, userResponse.Name)
	}
}

func (suite *userHandlerSuite) TestFindUser() {
	r := suite.Require()

	existingUsers := suite.createUsers(100)
	for _, user := range existingUsers {
		path := fmt.Sprintf(userFindPath, user.Id)
		w, err := suite.Request(path, http.MethodGet, nil, nil)
		r.Nilf(err, "error finding user: %v", err)
		r.Equal(w.Code, http.StatusOK)

		userResponse := &accountProto.UserResponse{}
		err = json.Unmarshal(w.Body.Bytes(), &userResponse)
		r.Nilf(err, "error parsing response: %v", err)

		r.Equal(user.Id, userResponse.Id)
		r.Equal(user.Name, userResponse.Name)
	}
}

func (suite *userHandlerSuite) createUsers(amount int) []*accountProto.UserResponse {
	existingUsers := make([]*accountProto.UserResponse, amount)
	for i := 0; i < amount; i++ {
		person := gofakeit.Person()
		password := gofakeit.Password(true, true, true, true, false, 6)

		userRequest := &accountProto.UserRequest{
			Name:                 fmt.Sprintf("%s %s", person.FirstName, person.LastName),
			Email:                person.Contact.Email,
			Password:             password,
			PasswordConfirmation: password,
		}
		userResponse := suite.createUser(userRequest)
		existingUsers[i] = userResponse
	}
	return existingUsers
}

func (suite *userHandlerSuite) createUser(req *accountProto.UserRequest) (res *accountProto.UserResponse) {
	r := suite.Require()

	suite.T().Logf("Creating user: %+v", req)

	w, err := suite.Request(userCreatePath, http.MethodPost, req, nil)
	r.Nilf(err, "error creating user: %v", err)
	r.Equal(w.Code, http.StatusCreated)

	res = &accountProto.UserResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &res)
	r.Nilf(err, "error parsing response: %v", err)

	return res
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, &userHandlerSuite{})
}
