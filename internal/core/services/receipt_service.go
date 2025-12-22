package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/billowdev/golang-api-clinic-management/pkg/pdf"
	"github.com/google/uuid"
	"time"
)

type receiptService struct {
	repo ports.ReceiptRepository
	pdfGenerator *pdf.ReceiptGenerator
}

func NewReceiptService(repo ports.ReceiptRepository, pdfGenerator *pdf.ReceiptGenerator) ports.ReceiptService {
	return &receiptService{
		repo: repo,
		pdfGenerator: pdfGenerator,
	}
}

func (s *receiptService) CreateReceipt(receipt *domain.Receipt) error {
	return s.repo.Create(receipt)
}

func (s *receiptService) GetReceiptByID(id uuid.UUID) (*domain.Receipt, error) {
	return s.repo.GetByID(id)
}

func (s *receiptService) GetAllReceipts() ([]*domain.Receipt, error) {
	return s.repo.GetAll()
}

func (s *receiptService) GetReceiptsByPatientID(patientID uuid.UUID) ([]*domain.Receipt, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *receiptService) GetReceiptsByDateRange(startDate, endDate time.Time) ([]*domain.Receipt, error) {
	return s.repo.GetByDateRange(startDate, endDate)
}

func (s *receiptService) SearchReceipts(query string) ([]*domain.Receipt, error) {
	return s.repo.Search(query)
}

func (s *receiptService) UpdateReceipt(receipt *domain.Receipt) error {
	return s.repo.Update(receipt)
}

func (s *receiptService) DeleteReceipt(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *receiptService) GenerateReceiptPDF(receiptID uuid.UUID) ([]byte, error) {
	receipt, err := s.repo.GetByID(receiptID)
	if err != nil {
		return nil, err
	}
	
	return s.pdfGenerator.Generate(receipt)
}

