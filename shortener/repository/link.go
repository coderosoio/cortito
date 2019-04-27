package repository

import (
	"github.com/jinzhu/gorm"

	"common/random"

	"shortener/model"
)

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) CreateLink(link *model.Link) (err error) {
	err = r.db.Create(&link).Error
	return
}

func (r *LinkRepository) ListLinksByUserID(userID interface{}) (links []*model.Link, err error) {
	links = make([]*model.Link, 0)
	err = r.db.Where("user_id = ?", userID).Find(&links).Error
	return
}

func (r *LinkRepository) FindLinkByHash(hash string) (link *model.Link, err error) {
	link = &model.Link{}
	err = r.db.Where("hash = ?", hash).First(&link).Error
	return
}

func (r *LinkRepository) UpdateLink(link *model.Link, updates map[string]interface{}) (err error) {
	err = r.db.Model(link).Update(updates).Error
	return
}

func (r *LinkRepository) GenerateHash(url string) (string, error) {
	for {
		hash := random.String(5)
		_, err := r.FindLinkByHash(hash)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return hash, nil
			}
			return "", err
		}
	}
}
