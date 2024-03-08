package repository

import (
	"github.com/akhil-is-watching/sainterview-backend/models"
	"gorm.io/gorm"
)

type InterviewQuestionRepo struct {
	db *gorm.DB
}

func NewInterviewQuestionRepo(db *gorm.DB) InterviewQuestionRepo {
	return InterviewQuestionRepo{
		db: db,
	}
}

func (repo InterviewQuestionRepo) CreateQuestion(input []models.InterviewQuestion) error {
	if err := repo.db.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

func (repo InterviewQuestionRepo) MarkStatus(ID string, Status string) error {
	if err := repo.db.Where("id = ?", ID).Update("status", Status).Error; err != nil {
		return err
	}

	return nil
}
