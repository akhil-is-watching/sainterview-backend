package repository

import (
	"github.com/akhil-is-watching/sainterview-backend/models"
	"gorm.io/gorm"
)

type InterviewRepository struct {
	db *gorm.DB
}

func NewInterviewRepository(db *gorm.DB) InterviewRepository {
	return InterviewRepository{
		db: db,
	}
}

func (repo InterviewRepository) Create(input models.Interview) error {
	if err := repo.db.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

func (repo InterviewRepository) GetInterviewByID(ID string) (models.Interview, error) {
	var interview models.Interview

	if err := repo.db.Where("id = ?", ID).First(&interview).Error; err != nil {
		return interview, err
	}

	return interview, nil
}

func (repo InterviewRepository) GetInterviewsByOrganizationID(OrganizationID string) ([]models.Interview, error) {
	var interviews []models.Interview

	if err := repo.db.Where("organization_id = ?", OrganizationID).Find(&interviews).Error; err != nil {
		return interviews, err
	}

	return interviews, nil
}
