package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"common/hashing"

	"account/model"
	"account/option"
	proto "account/proto/account"
	"account/repository"
)

type userHandler struct {
	userRepository  *repository.UserRepository
	hashingStrategy hashing.Hashing
}

// NewUserHandler returns an instance of `UserHandler`.
func NewUserHandler(options *option.Options) proto.UserHandler {
	return &userHandler{
		userRepository:  options.UserRepository,
		hashingStrategy: options.HashingStrategy,
	}
}

func (h *userHandler) UserExists(ctx context.Context, req *proto.UserRequest, res *proto.UserExistResponse) error {
	if len(strings.TrimSpace(req.Email)) == 0 {
		return fmt.Errorf("invalid email address")
	}
	exists, err := h.userRepository.UserExists(req.Email)
	if err != nil {
		return err
	}
	res.Exists = exists
	return nil
}

// CreateUser creates a new user.
func (h *userHandler) CreateUser(ctx context.Context, req *proto.UserRequest, res *proto.UserResponse) error {
	if err := validateCreateUser(req); err != nil {
		return err
	}
	passwordHash, err := h.hashingStrategy.GenerateFromPassword(req.Password)
	if err != nil {
		return err
	}
	exists, err := h.userRepository.UserExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	user := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}
	if err := h.userRepository.CreateUser(user); err != nil {
		return err
	}
	res.Id = int32(user.ID)
	res.Name = user.Name
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
	return nil
}

// UpdateUser updates an existing user.
func (h *userHandler) UpdateUser(ctx context.Context, req *proto.UserRequest, res *proto.UserResponse) error {
	user, err := h.userRepository.FindUser(req.Id)
	if err != nil {
		return err
	}
	if err := validateUpdateUser(req); err != nil {
		return err
	}
	updates := make(map[string]interface{})
	if len(strings.TrimSpace(req.Name)) > 0 {
		updates["name"] = req.Name
	}
	if len(strings.TrimSpace(req.Password)) > 0 {
		passwordHash, err := h.hashingStrategy.GenerateFromPassword(req.Password)
		if err != nil {
			return err
		}
		updates["password_hash"] = passwordHash
	}
	if err := h.userRepository.UpdateUser(user, updates); err != nil {
		return err
	}
	res.Id = int32(user.ID)
	res.Name = user.Name
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
	return nil
}

func (h *userHandler) FindUser(ctx context.Context, req *proto.UserRequest, res *proto.UserResponse) error {
	user, err := h.userRepository.FindUser(req.Id)
	if err != nil {
		return err
	}
	res.Id = int32(user.ID)
	res.Name = user.Name
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = user.UpdatedAt.Format(time.RFC3339)
	return nil
}

func validateCreateUser(req *proto.UserRequest) (err error) {
	err = validation.ValidateStruct(
		req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.PasswordConfirmation, validation.By(validatePasswordConfirmation(req))),
	)
	return
}

func validateUpdateUser(req *proto.UserRequest) (err error) {
	rules := make([]*validation.FieldRules, 0)
	if len(strings.TrimSpace(req.Name)) > 0 {
		rule := validation.Field(&req.Name, validation.Required)
		rules = append(rules, rule)
	}
	if len(strings.TrimSpace(req.Password)) > 0 {
		rule := validation.Field(
			&req.PasswordConfirmation,
			validation.Required,
			validation.By(validatePasswordConfirmation(req)),
		)
		rules = append(rules, rule)
	}
	err = validation.ValidateStruct(req, rules...)
	return
}

func validatePasswordConfirmation(req *proto.UserRequest) validation.RuleFunc {
	return func(value interface{}) (err error) {
		confirmation, _ := value.(string)
		if confirmation != req.Password {
			err = fmt.Errorf("does not match password")
		}
		return
	}
}
