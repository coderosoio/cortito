package handler

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"

	"web/option"
)

type AuthHandler struct {
	authService accountProto.AuthService
}

type AuthRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func init() {
	gob.Register(&accountProto.UserResponse{})
}

func NewAuthHandler(options *option.Options) *AuthHandler {
	return &AuthHandler{
		authService: options.AuthService,
	}
}

func (h *AuthHandler) New(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "auth/new", gin.H{
		"user": gin.H{},
	})
}

func (h *AuthHandler) Create(ctx *gin.Context) {
	authRequest := &AuthRequest{}
	if err := ctx.ShouldBind(authRequest); err != nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	req := &accountProto.AuthRequest{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}
	res, err := h.authService.CreateToken(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "auth/new", gin.H{
			"user":   authRequest,
			"errors": err,
		})
		return
	}
	session := sessions.Default(ctx)
	session.Set("token", res.Token)
	session.Set("user", res.User)
	if err := session.Save(); err != nil {
		ctx.HTML(http.StatusInternalServerError, "auth/new", gin.H{
			"user":   authRequest,
			"errors": err,
		})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func (h *AuthHandler) Delete(ctx *gin.Context) {
	session := sessions.Default(ctx)
	token := session.Get("token")
	req := &accountProto.RevokeTokenRequest{
		Token: token.(string),
	}
	res, err := h.authService.RevokeToken(ctx, req)
	if err != nil {
		log.Printf("error revoking token: %v", err)
	}
	log.Printf("response: %+v", res)
	session.Clear()
	ctx.Redirect(http.StatusMovedPermanently, "/")
}
