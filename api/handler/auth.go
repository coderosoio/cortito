package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	accountProto "account/proto/account"

	"api/helper"
	"api/option"
)

type AuthHandler struct {
	authService accountProto.AuthService
}

func NewAuthHandler(options *option.Options) *AuthHandler {
	return &AuthHandler{
		authService: options.AuthService,
	}
}

func (h *AuthHandler) CreateToken(ctx *gin.Context) {
	req := &accountProto.AuthRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	if err := validateCreateToken(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	res, err := h.authService.CreateToken(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) RevokeToken(ctx *gin.Context) {
	token, found := helper.CurrentToken(ctx)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "token not present",
		})
		return
	}
	req := &accountProto.RevokeTokenRequest{
		Token: token,
	}
	_, err := h.authService.RevokeToken(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func validateCreateToken(req *accountProto.AuthRequest) (err error) {
	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
	)
	return
}
