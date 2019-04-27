package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"common/connection"
	commonTesting "common/testing"

	"shortener/model"
	"shortener/option"
	proto "shortener/proto/shortener"
	"shortener/repository"
)

type linkHandlerSuite struct {
	commonTesting.Suite
	handler proto.LinkHandler
}

func (suite *linkHandlerSuite) SetupSuite() {
	r := suite.Require()

	suite.Init()

	suite.DeleteDatabaseFile("shortener")

	db, err := connection.GetDatabaseConnection("shortener")
	r.Nilf(err, "error getting database connection: %v", err)

	suite.MigrateModels(
		"shortener",
		&model.Link{},
	)

	linkRepository := repository.NewLinkRepository(db)

	options := option.NewOptions(
		option.WithLinkRepository(linkRepository),
	)

	suite.handler = NewLinkHandler(options)
}

func (suite *linkHandlerSuite) SetupTest() {
	suite.TruncateTables(
		"shortener",
		"links",
	)
}

func (suite *linkHandlerSuite) TearDownSuite() {
	suite.DeleteDatabaseFile("shortener")
}

func (suite *linkHandlerSuite) TestCreateLink() {
	r := suite.Require()

	ctx := context.TODO()
	res, err := createDummyLink(
		suite.handler,
		ctx,
		1,
		"https://google.com",
	)
	r.Nilf(err, "error creating link: %v", err)

	r.NotEmpty(res.Hash)
	r.Zero(res.Visits)
}

func (suite *linkHandlerSuite) TestListLinks() {
	r := suite.Require()

	ctx := context.TODO()
	createResponse, err := createDummyLink(
		suite.handler,
		ctx,
		1,
		"https://google.com",
	)
	r.Nilf(err, "error creating link: %v", err)

	req := &proto.ListLinksRequest{
		UserId: createResponse.UserId,
	}
	res := &proto.ListLinksResponse{}
	err = suite.handler.ListLinks(ctx, req, res)
	r.Nilf(err, "error listing links: %v", err)

	r.Len(res.Links, 1)

	link := res.Links[0]

	r.NotNil(link)
	r.NotEmpty(link.Hash)
	r.Equal("https://google.com", link.Url)
}

func (suite *linkHandlerSuite) TestFindLink() {
	r := suite.Require()

	ctx := context.TODO()
	createResponse, err := createDummyLink(
		suite.handler,
		ctx,
		1,
		"https://google.com",
	)
	r.Nilf(err, "error creating link: %v", err)

	req := &proto.LinkRequest{
		Hash: createResponse.Hash,
	}
	res := &proto.LinkResponse{}
	err = suite.handler.FindLink(ctx, req, res)
	r.Nilf(err, "error finding link: %v", err)

	r.Equal(createResponse.Hash, res.Hash)
	r.Equal(createResponse.Url, res.Url)
}

func (suite *linkHandlerSuite) TestIncreaseVisit() {
	r := suite.Require()

	ctx := context.TODO()
	res, err := createDummyLink(
		suite.handler,
		ctx,
		1,
		"https://google.com",
	)
	r.Nilf(err, "error creating link: %v", err)

	r.Zero(res.Visits)

	req := &proto.LinkRequest{
		Hash: res.Hash,
	}

	res = &proto.LinkResponse{}
	err = suite.handler.IncreaseVisit(ctx, req, res)
	r.Nilf(err, "error increasing visits: %v", err)

	linkResponse := &proto.LinkResponse{}
	err = suite.handler.FindLink(ctx, req, linkResponse)
	r.Nilf(err, "error finding link: %v", err)

	r.NotZero(res.Visits)
	r.Equal(linkResponse.Visits, res.Visits)
}

func TestLinkHandlerSuite(t *testing.T) {
	suite.Run(t, &linkHandlerSuite{})
}

func createDummyLink(
	handler proto.LinkHandler,
	ctx context.Context,
	userID int32,
	url string) (res *proto.LinkResponse, err error) {
	req := &proto.CreateLinkRequest{
		UserId: userID,
		Url:    url,
	}
	res = &proto.LinkResponse{}
	err = handler.CreateLink(ctx, req, res)
	return
}
