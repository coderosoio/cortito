package repository

import (
	"github.com/jinzhu/gorm"

	"account/model"
)

type UserTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository(db *gorm.DB) *UserTokenRepository {
	return &UserTokenRepository{
		db: db,
	}
}

func (r *UserTokenRepository) CreateUserToken(userToken *model.UserToken) (err error) {
	err = r.db.Create(&userToken).Error
	return
}

func (r *UserTokenRepository) DeleteUserToken(userToken *model.UserToken) (err error) {
	err = r.db.Delete(&userToken).Error
	return
}

func (r *UserTokenRepository) FindUserTokenByToken(token string) (userToken *model.UserToken, err error) {
	userToken = &model.UserToken{}
	err = r.db.Where("token = ?", token).First(&userToken).Error
	return
}
