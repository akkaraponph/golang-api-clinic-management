package services

import (
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/ports"
	"github.com/google/uuid"
)

type medicineStockService struct {
	repo ports.MedicineStockRepository
}

func NewMedicineStockService(repo ports.MedicineStockRepository) ports.MedicineStockService {
	return &medicineStockService{repo: repo}
}

func (s *medicineStockService) CreateStock(stock *domain.MedicineStock) error {
	return s.repo.Create(stock)
}

func (s *medicineStockService) GetStockByID(id uuid.UUID) (*domain.MedicineStock, error) {
	return s.repo.GetByID(id)
}

func (s *medicineStockService) GetStockByMedicineID(medicineID uuid.UUID) ([]*domain.MedicineStock, error) {
	return s.repo.GetByMedicineID(medicineID)
}

func (s *medicineStockService) GetAllStocks() ([]*domain.MedicineStock, error) {
	return s.repo.GetAll()
}

func (s *medicineStockService) UpdateStock(stock *domain.MedicineStock) error {
	return s.repo.Update(stock)
}

func (s *medicineStockService) DeleteStock(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *medicineStockService) GetRemainingStocks() ([]*domain.MedicineStock, error) {
	return s.repo.GetRemainingStock()
}

