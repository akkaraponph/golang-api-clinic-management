package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) ports.MaterialRepository {
	return &materialRepository{db: db}
}

func (r *materialRepository) Create(material *domain.Material) error {
	return r.db.Create(material).Error
}

func (r *materialRepository) GetByID(id uuid.UUID) (*domain.Material, error) {
	var material domain.Material
	err := r.db.Preload("Agent").Where("id = ?", id).First(&material).Error
	if err != nil {
		return nil, err
	}
	return &material, nil
}

func (r *materialRepository) GetAll() ([]*domain.Material, error) {
	var materials []*domain.Material
	err := r.db.Preload("Agent").Find(&materials).Error
	return materials, err
}

func (r *materialRepository) Update(material *domain.Material) error {
	return r.db.Save(material).Error
}

func (r *materialRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Material{}, id).Error
}

