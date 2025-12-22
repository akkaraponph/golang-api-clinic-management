package repositories

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) ports.PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(patient *domain.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) GetByID(id uuid.UUID) (*domain.Patient, error) {
	var patient domain.Patient
	err := r.db.Where("id = ? AND is_remove = ?", id, false).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetAll() ([]*domain.Patient, error) {
	var patients []*domain.Patient
	err := r.db.Where("is_remove = ?", false).Order("updated_at ASC").Find(&patients).Error
	return patients, err
}

func (r *patientRepository) Search(query string) ([]*domain.Patient, error) {
	var patients []*domain.Patient
	queryLower := strings.ToLower(query)
	
	// Handle Thai gender translations
	if queryLower == "หญิง" || queryLower == "เพศหญิง" {
		query = "female"
	} else if queryLower == "ชาย" || queryLower == "เพศชาย" {
		query = "male"
	}
	
	searchPattern := "%" + query + "%"
	err := r.db.Where("is_remove = ? AND ("+
		"name LIKE ? OR surname LIKE ? OR phone LIKE ? OR id_card LIKE ? OR "+
		"sex LIKE ? OR birth_date::text LIKE ? OR blood_type LIKE ? OR "+
		"disease LIKE ? OR allergic LIKE ? OR address LIKE ? OR "+
		"subdistrict LIKE ? OR district LIKE ? OR province LIKE ? OR zipcode LIKE ?)",
		false, searchPattern, searchPattern, searchPattern, searchPattern,
		searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		searchPattern, searchPattern, searchPattern, searchPattern, searchPattern).Find(&patients).Error
	return patients, err
}

func (r *patientRepository) Update(patient *domain.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Patient{}, id).Error
}

func (r *patientRepository) SoftDelete(id uuid.UUID) error {
	return r.db.Model(&domain.Patient{}).Where("id = ?", id).Update("is_remove", true).Error
}

func (r *patientRepository) Restore(id uuid.UUID) error {
	return r.db.Model(&domain.Patient{}).Where("id = ?", id).Update("is_remove", false).Error
}

func (r *patientRepository) GetTrash() ([]*domain.Patient, error) {
	var patients []*domain.Patient
	err := r.db.Where("is_remove = ?", true).Find(&patients).Error
	return patients, err
}

