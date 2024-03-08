package repository

import (
	"github.com/akhil-is-watching/sainterview-backend/models"
	"gorm.io/gorm"
)

type AvatarRepository struct {
	db *gorm.DB
}

func NewAvatarRepository(db *gorm.DB) AvatarRepository {
	return AvatarRepository{
		db: db,
	}
}

func (repo AvatarRepository) CreateAvatar(input []models.Avatar) error {
	if err := repo.db.Save(&input).Error; err != nil {
		return err
	}

	return nil
}

func (repo AvatarRepository) GetByID(ID int) (models.Avatar, error) {
	var avatar models.Avatar

	if err := repo.db.Where("id = ?", ID).First(&avatar).Error; err != nil {
		return avatar, err
	}

	return avatar, nil
}

func (repo AvatarRepository) All() ([]models.Avatar, error) {
	var avatars []models.Avatar

	if err := repo.db.Find(&avatars).Error; err != nil {
		return avatars, err
	}

	return avatars, nil
}
