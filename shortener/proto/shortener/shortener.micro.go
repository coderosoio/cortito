// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/shortener/shortener.proto

/*
Package shortener is a generated protocol buffer package.

It is generated from these files:
	proto/shortener/shortener.proto

It has these top-level messages:
	CreateLinkRequest
	ListLinksRequest
	ListLinksResponse
	LinkRequest
	LinkResponse
*/
package shortener

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"

	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors.js if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors.js if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Link service

type LinkService interface {
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...client.CallOption) (*LinkResponse, error)
	ListLinks(ctx context.Context, in *ListLinksRequest, opts ...client.CallOption) (*ListLinksResponse, error)
	FindLink(ctx context.Context, in *LinkRequest, opts ...client.CallOption) (*LinkResponse, error)
	IncreaseVisit(ctx context.Context, in *LinkRequest, opts ...client.CallOption) (*LinkResponse, error)
}

type linkService struct {
	c    client.Client
	name string
}

func NewLinkService(name string, c client.Client) LinkService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "shortener"
	}
	return &linkService{
		c:    c,
		name: name,
	}
}

func (c *linkService) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...client.CallOption) (*LinkResponse, error) {
	req := c.c.NewRequest(c.name, "Link.CreateLink", in)
	out := new(LinkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) ListLinks(ctx context.Context, in *ListLinksRequest, opts ...client.CallOption) (*ListLinksResponse, error) {
	req := c.c.NewRequest(c.name, "Link.ListLinks", in)
	out := new(ListLinksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) FindLink(ctx context.Context, in *LinkRequest, opts ...client.CallOption) (*LinkResponse, error) {
	req := c.c.NewRequest(c.name, "Link.FindLink", in)
	out := new(LinkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkService) IncreaseVisit(ctx context.Context, in *LinkRequest, opts ...client.CallOption) (*LinkResponse, error) {
	req := c.c.NewRequest(c.name, "Link.IncreaseVisit", in)
	out := new(LinkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Link service

type LinkHandler interface {
	CreateLink(context.Context, *CreateLinkRequest, *LinkResponse) error
	ListLinks(context.Context, *ListLinksRequest, *ListLinksResponse) error
	FindLink(context.Context, *LinkRequest, *LinkResponse) error
	IncreaseVisit(context.Context, *LinkRequest, *LinkResponse) error
}

func RegisterLinkHandler(s server.Server, hdlr LinkHandler, opts ...server.HandlerOption) error {
	type link interface {
		CreateLink(ctx context.Context, in *CreateLinkRequest, out *LinkResponse) error
		ListLinks(ctx context.Context, in *ListLinksRequest, out *ListLinksResponse) error
		FindLink(ctx context.Context, in *LinkRequest, out *LinkResponse) error
		IncreaseVisit(ctx context.Context, in *LinkRequest, out *LinkResponse) error
	}
	type Link struct {
		link
	}
	h := &linkHandler{hdlr}
	return s.Handle(s.NewHandler(&Link{h}, opts...))
}

type linkHandler struct {
	LinkHandler
}

func (h *linkHandler) CreateLink(ctx context.Context, in *CreateLinkRequest, out *LinkResponse) error {
	return h.LinkHandler.CreateLink(ctx, in, out)
}

func (h *linkHandler) ListLinks(ctx context.Context, in *ListLinksRequest, out *ListLinksResponse) error {
	return h.LinkHandler.ListLinks(ctx, in, out)
}

func (h *linkHandler) FindLink(ctx context.Context, in *LinkRequest, out *LinkResponse) error {
	return h.LinkHandler.FindLink(ctx, in, out)
}

func (h *linkHandler) IncreaseVisit(ctx context.Context, in *LinkRequest, out *LinkResponse) error {
	return h.LinkHandler.IncreaseVisit(ctx, in, out)
}
