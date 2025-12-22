package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type medicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) ports.MedicineRepository {
	return &medicineRepository{db: db}
}

func (r *medicineRepository) Create(medicine *domain.Medicine) error {
	return r.db.Create(medicine).Error
}

func (r *medicineRepository) GetByID(id uuid.UUID) (*domain.Medicine, error) {
	var medicine domain.Medicine
	err := r.db.Preload("MedicineType").Where("id = ?", id).First(&medicine).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *medicineRepository) GetAll() ([]*domain.Medicine, error) {
	var medicines []*domain.Medicine
	err := r.db.Preload("MedicineType").Find(&medicines).Error
	return medicines, err
}

func (r *medicineRepository) Search(query string) ([]*domain.Medicine, error) {
	var medicines []*domain.Medicine
	searchPattern := "%" + strings.ToLower(query) + "%"
	err := r.db.Preload("MedicineType").
		Where("LOWER(name) LIKE ? OR LOWER(detail) LIKE ? OR LOWER(dosage) LIKE ?",
			searchPattern, searchPattern, searchPattern).Find(&medicines).Error
	return medicines, err
}

func (r *medicineRepository) Update(medicine *domain.Medicine) error {
	return r.db.Save(medicine).Error
}

func (r *medicineRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Medicine{}, id).Error
}

func (r *medicineRepository) SoftDelete(id uuid.UUID) error {
	// If using soft delete pattern
	return r.db.Delete(&domain.Medicine{}, id).Error
}

func (r *medicineRepository) Restore(id uuid.UUID) error {
	return nil
}

