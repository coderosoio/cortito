package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	shortenerProto "shortener/proto/shortener"

	"api/option"
)

type linkHandler struct {
	linkService shortenerProto.LinkService
}

func NewLinkHandler(options *option.Options) *linkHandler {
	return &linkHandler{
		linkService: options.LinkService,
	}
}

func (h *linkHandler) Create(ctx *gin.Context) {
	user := CurrentUser(ctx)
	req := &shortenerProto.CreateLinkRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	req.UserId = user.Id
	if err := validateCreateLink(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	res, err := h.linkService.CreateLink(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *linkHandler) Index(ctx *gin.Context) {
	user := CurrentUser(ctx)
	req := &shortenerProto.ListLinksRequest{
		UserId: user.Id,
	}
	res, err := h.linkService.ListLinks(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Links)
}

func validateCreateLink(req *shortenerProto.CreateLinkRequest) (err error) {
	err = validation.ValidateStruct(
		req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Url, validation.Required, is.URL),
	)
	return
}
