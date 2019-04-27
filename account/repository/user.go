package repository

import (
	"github.com/jinzhu/gorm"

	"account/model"
)

// UserRepository performs operations on users.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns an instance of `UserRepository`.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) UserExists(email string) (bool, error) {
	user := &model.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateUser creates a new user.
func (r *UserRepository) CreateUser(user *model.User) (err error) {
	err = r.db.Create(&user).Error
	return
}

// UpdateUser updates an existing user.
func (r *UserRepository) UpdateUser(user *model.User, updates map[string]interface{}) (err error) {
	err = r.db.Model(&user).Update(updates).Error
	return
}

// FindUser finds the user for the given `id`.
func (r *UserRepository) FindUser(id interface{}) (user *model.User, err error) {
	user = &model.User{}
	err = r.db.Find(&user, id).Error
	return
}

// FindUserByEmail finds the user with the given `email`.
func (r *UserRepository) FindUserByEmail(email string) (user *model.User, err error) {
	user = &model.User{}
	err = r.db.Where("email = ?", email).First(&user).Error
	return
}
