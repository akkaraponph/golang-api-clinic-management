package ports

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
	"time"
)

type ReceiptRepository interface {
	Create(receipt *domain.Receipt) error
	GetByID(id uuid.UUID) (*domain.Receipt, error)
	GetAll() ([]*domain.Receipt, error)
	GetByPatientID(patientID uuid.UUID) ([]*domain.Receipt, error)
	GetByDateRange(startDate, endDate time.Time) ([]*domain.Receipt, error)
	Search(query string) ([]*domain.Receipt, error)
	Update(receipt *domain.Receipt) error
	Delete(id uuid.UUID) error
}

type ReceiptService interface {
	CreateReceipt(receipt *domain.Receipt) error
	GetReceiptByID(id uuid.UUID) (*domain.Receipt, error)
	GetAllReceipts() ([]*domain.Receipt, error)
	GetReceiptsByPatientID(patientID uuid.UUID) ([]*domain.Receipt, error)
	GetReceiptsByDateRange(startDate, endDate time.Time) ([]*domain.Receipt, error)
	SearchReceipts(query string) ([]*domain.Receipt, error)
	UpdateReceipt(receipt *domain.Receipt) error
	DeleteReceipt(id uuid.UUID) error
	GenerateReceiptPDF(receiptID uuid.UUID) ([]byte, error)
}

