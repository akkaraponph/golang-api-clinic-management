package repositories

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type receiptRepository struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ports.ReceiptRepository {
	return &receiptRepository{db: db}
}

func (r *receiptRepository) Create(receipt *domain.Receipt) error {
	return r.db.Create(receipt).Error
}

func (r *receiptRepository) GetByID(id uuid.UUID) (*domain.Receipt, error) {
	var receipt domain.Receipt
	err := r.db.Preload("Patient").Preload("Employee").
		Preload("Course").Preload("MedicineDisburse").
		Where("id = ?", id).First(&receipt).Error
	if err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (r *receiptRepository) GetAll() ([]*domain.Receipt, error) {
	var receipts []*domain.Receipt
	err := r.db.Preload("Patient").Preload("Employee").
		Order("date DESC").Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepository) GetByPatientID(patientID uuid.UUID) ([]*domain.Receipt, error) {
	var receipts []*domain.Receipt
	err := r.db.Where("patient_id = ?", patientID).
		Order("date DESC").Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepository) GetByDateRange(startDate, endDate time.Time) ([]*domain.Receipt, error) {
	var receipts []*domain.Receipt
	err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).
		Order("date DESC").Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepository) Search(query string) ([]*domain.Receipt, error) {
	var receipts []*domain.Receipt
	searchPattern := "%" + query + "%"
	
	// Search across receipt fields and related entities
	err := r.db.Preload("Patient").Preload("Employee").
		Where("id::text LIKE ? OR date::text LIKE ? OR total_price::text LIKE ?", 
			searchPattern, searchPattern, searchPattern).
		Or("EXISTS (SELECT 1 FROM courses WHERE courses.id = receipts.course_id AND "+
			"(courses.id::text LIKE ? OR courses.weight::text LIKE ? OR courses.height::text LIKE ?))",
			searchPattern, searchPattern, searchPattern).
		Or("EXISTS (SELECT 1 FROM patients WHERE patients.id = receipts.patient_id AND "+
			"(patients.id::text LIKE ? OR patients.name LIKE ? OR patients.surname LIKE ?))",
			searchPattern, searchPattern, searchPattern).
		Or("EXISTS (SELECT 1 FROM employees WHERE employees.id = receipts.employee_id AND "+
			"(employees.id::text LIKE ? OR employees.name LIKE ? OR employees.surname LIKE ?))",
			searchPattern, searchPattern, searchPattern).
		Order("date DESC").Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepository) Update(receipt *domain.Receipt) error {
	return r.db.Save(receipt).Error
}

func (r *receiptRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Receipt{}, id).Error
}

