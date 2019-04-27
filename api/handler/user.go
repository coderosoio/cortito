package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"

	"api/option"
)

type UserHandler struct {
	userService accountProto.UserService
}

func NewUserHandler(options *option.Options) *UserHandler {
	return &UserHandler{
		userService: options.UserService,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	req := &accountProto.UserRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	if err := validateCreateUser(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	res, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid user id",
		})
		return
	}
	req := &accountProto.UserRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}
	req.Id = int32(id)
	if err := validateUpdateUser(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
		return
	}
	res, err := h.userService.UpdateUser(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *UserHandler) Find(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid user id",
		})
		return
	}
	req := &accountProto.UserRequest{
		Id: int32(id),
	}
	res, err := h.userService.FindUser(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func validateCreateUser(req *accountProto.UserRequest) (err error) {
	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.PasswordConfirmation, validation.By(validatePasswordConfirmation(req))),
	)
	return
}

func validateUpdateUser(req *accountProto.UserRequest) (err error) {
	rules := make([]*validation.FieldRules, 0)
	if len(strings.TrimSpace(req.Name)) > 0 {
		rule := validation.Field(&req.Name, validation.Required)
		rules = append(rules, rule)
	}
	if len(strings.TrimSpace(req.Password)) > 0 {
		rule := validation.Field(&req.Password, validation.Required)
		rules = append(rules, rule)

		rule = validation.Field(&req.PasswordConfirmation, validation.Required, validation.By(validatePasswordConfirmation(req)))
		rules = append(rules, rule)
	}
	err = validation.ValidateStruct(req, rules...)
	return
}

func validatePasswordConfirmation(req *accountProto.UserRequest) validation.RuleFunc {
	return func(value interface{}) (err error) {
		confirmation, _ := value.(string)
		if confirmation != req.Password {
			err = fmt.Errorf("does not match password")
		}
		return
	}
}
