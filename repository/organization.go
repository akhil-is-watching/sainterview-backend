package repository

import (
	"fmt"

	"github.com/akhil-is-watching/sainterview-backend/helpers"
	"github.com/akhil-is-watching/sainterview-backend/models"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return OrganizationRepository{
		db: db,
	}
}

func (repo OrganizationRepository) CreateOrganization(input models.Organization) error {
	input.ID = helpers.UIDGen().GenerateID("OR")
	if err := repo.db.Create(&input).Error; err != nil {
		return err
	}

	return nil
}

func (repo OrganizationRepository) FindOrganizationByID(ID string) (models.Organization, error) {
	var organization models.Organization

	err := repo.db.Where("id = ?", ID).First(&organization).Error
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (repo OrganizationRepository) FindByCredentials(email, password string) (models.Organization, error) {
	var organization models.Organization

	err := repo.db.Where("email = ?", email).Where("password = ?", password).First(&organization).Error
	if err != nil {
		return organization, fmt.Errorf("invalid credentials")
	}

	return organization, nil
}
