package handler

import (
	"common/keyvalue"
	"context"
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

	"shortener/model"
	"shortener/option"
	proto "shortener/proto/shortener"
	"shortener/repository"
)

type linkHandler struct {
	linkRepository  *repository.LinkRepository
	keyValueStorage keyvalue.Storage
}

func NewLinkHandler(options *option.Options) proto.LinkHandler {
	return &linkHandler{
		linkRepository:  options.LinkRepository,
		keyValueStorage: options.KeyValueStorage,
	}
}

func (h *linkHandler) CreateLink(ctx context.Context, req *proto.CreateLinkRequest, res *proto.LinkResponse) error {
	if err := validateCreateLink(req); err != nil {
		return err
	}
	hash, err := h.linkRepository.GenerateHash(req.Url)
	if err != nil {
		return err
	}
	link := &model.Link{
		UserID: uint(req.UserId),
		URL:    req.Url,
		Hash:   hash,
	}
	if err := h.linkRepository.CreateLink(link); err != nil {
		return err
	}
	res.Id = int32(link.ID)
	res.Url = link.URL
	res.UserId = int32(link.UserID)
	res.Hash = link.Hash
	res.Visits = 0
	res.CreatedAt = link.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = link.UpdatedAt.Format(time.RFC3339)
	res.LastVisit = ""
	go func() {
		if err := h.keyValueStorage.Set(link.Hash, link.URL); err != nil {
			log.Printf("error saving link hash to key value storage")
		}
	}()
	return nil
}

func (h *linkHandler) ListLinks(ctx context.Context, req *proto.ListLinksRequest, res *proto.ListLinksResponse) error {
	links, err := h.linkRepository.ListLinksByUserID(req.UserId)
	if err != nil {
		return err
	}
	res.Links = make([]*proto.LinkResponse, len(links))
	for i, link := range links {
		lastVisit := ""
		if link.LastVisit != nil {
			lastVisit = link.LastVisit.Format(time.RFC3339)
		}
		res.Links[i] = &proto.LinkResponse{
			Id:        int32(link.ID),
			UserId:    int32(link.UserID),
			Url:       link.URL,
			Hash:      link.Hash,
			Visits:    int32(link.Visits),
			CreatedAt: link.CreatedAt.Format(time.RFC3339),
			UpdatedAt: link.UpdatedAt.Format(time.RFC3339),
			LastVisit: lastVisit,
		}
	}
	return nil
}

func (h *linkHandler) FindLink(ctx context.Context, req *proto.LinkRequest, res *proto.LinkResponse) error {
	link, err := h.linkRepository.FindLinkByHash(req.Hash)
	if err != nil {
		return err
	}
	lastVisit := ""
	if link.LastVisit != nil {
		lastVisit = link.LastVisit.Format(time.RFC3339)
	}
	res.Id = int32(link.ID)
	res.UserId = int32(link.UserID)
	res.Url = link.URL
	res.Hash = link.Hash
	res.Visits = int32(link.Visits)
	res.LastVisit = lastVisit
	res.CreatedAt = link.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = link.UpdatedAt.Format(time.RFC3339)
	return nil
}

func (h *linkHandler) IncreaseVisit(ctx context.Context, req *proto.LinkRequest, res *proto.LinkResponse) error {
	link, err := h.linkRepository.FindLinkByHash(req.Hash)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"visits":     link.Visits + 1,
		"last_visit": time.Now(),
	}
	if err := h.linkRepository.UpdateLink(link, updates); err != nil {
		return err
	}
	lastVisit := ""
	if link.LastVisit != nil {
		lastVisit = link.LastVisit.Format(time.RFC3339)
	}
	res.Id = int32(link.ID)
	res.UserId = int32(link.UserID)
	res.Url = link.URL
	res.Hash = link.Hash
	res.Visits = int32(link.Visits)
	res.LastVisit = lastVisit
	res.CreatedAt = link.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = link.UpdatedAt.Format(time.RFC3339)
	return nil
}

func validateCreateLink(req *proto.CreateLinkRequest) (err error) {
	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Url, validation.Required),
		validation.Field(&req.UserId, validation.Required),
	)
	return
}
