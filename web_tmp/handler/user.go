package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"

	"web/option"
)

type UserHandler struct {
	userService accountProto.UserService
}

type UserRequest struct {
	Name                 string `form:"name" json:"name"`
	Email                string `form:"email" json:"email"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
}

func NewUserHandler(options *option.Options) *UserHandler {
	return &UserHandler{
		userService: options.UserService,
	}
}

func (h *UserHandler) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/new", gin.H{
		"user": gin.H{},
	})
}

func (h *UserHandler) Create(ctx *gin.Context) {
	userRequest := &UserRequest{}
	if err := ctx.ShouldBind(userRequest); err != nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	// TODO: Validate request before send it to user service.
	req := &accountProto.UserRequest{
		Name:                 userRequest.Name,
		Email:                userRequest.Email,
		Password:             userRequest.Password,
		PasswordConfirmation: userRequest.PasswordConfirmation,
	}
	_, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusUnprocessableEntity, "user/new", gin.H{
			"user":   userRequest,
			"errors": err,
		})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/sessions/new")
	ctx.Status(http.StatusCreated)
}
