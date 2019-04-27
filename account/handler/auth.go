package handler

import (
	"context"
	"fmt"
	"time"

	"common/auth"
	"common/hashing"

	"account/model"
	"account/option"
	proto "account/proto/account"
	"account/repository"
)

type authHandler struct {
	auth                auth.Auth
	userRepository      *repository.UserRepository
	userTokenRepository *repository.UserTokenRepository
	hashing             hashing.Hashing
}

// NewAuthHandler returns an instance to `authHandler`.
func NewAuthHandler(options *option.Options) proto.AuthHandler {
	return &authHandler{
		auth:                options.AuthStrategy,
		userRepository:      options.UserRepository,
		userTokenRepository: options.UserTokenRepository,
		hashing:             options.HashingStrategy,
	}
}

// CreateToken creates a new token.
func (h *authHandler) CreateToken(ctx context.Context, req *proto.AuthRequest, res *proto.AuthResponse) error {
	user, err := h.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return err
	}
	passwordMatch, err := h.hashing.ComparePasswordAndHash(req.Password, user.PasswordHash)
	if err != nil {
		return err
	}
	if !passwordMatch {
		return fmt.Errorf("wrong password")
	}
	data := map[string]interface{}{
		"sub":   user.ID,
		"email": user.Email,
	}
	token, err := h.auth.GenerateToken(data)
	if err != nil {
		return fmt.Errorf("error generating token: %v", err)
	}
	userToken := &model.UserToken{
		Token:  token,
		UserID: user.ID,
	}
	if err := h.userTokenRepository.CreateUserToken(userToken); err != nil {
		return err
	}
	res.Token = token
	res.User = &proto.UserResponse{
		Id:        int32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
	return nil
}

// VerifyToken verifies that a token is valid.
func (h *authHandler) VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest, res *proto.VerifyTokenResponse) error {
	userToken, err := h.userTokenRepository.FindUserTokenByToken(req.Token)
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}
	data, err := h.auth.VerifyToken(req.Token)
	if err != nil {
		if err := h.userTokenRepository.DeleteUserToken(userToken); err != nil {
			return err
		}
		return err
	}
	id := uint(data["sub"].(float64))
	user, err := h.userRepository.FindUser(id)
	if err != nil {
		return err
	}
	res.Token = req.Token
	res.User = &proto.UserResponse{
		Id:        int32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
	return nil
}

func (h *authHandler) RevokeToken(ctx context.Context, req *proto.RevokeTokenRequest, res *proto.RevokeTokenResponse) error {
	userToken, err := h.userTokenRepository.FindUserTokenByToken(req.Token)
	if err != nil {
		return err
	}
	if err := h.userTokenRepository.DeleteUserToken(userToken); err != nil {
		return err
	}
	return nil
}
