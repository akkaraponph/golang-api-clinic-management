package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineOrderDetailCreateRequest struct {
	MedicineID uuid.UUID `json:"medicine_id" validate:"required"`
	Price      float64   `json:"price" validate:"min=0"`
	Qty        int       `json:"qty" validate:"required,min=1"`
	Remain     int       `json:"remain" validate:"min=0"`
}

type MedicineOrderCreateRequest struct {
	Date        time.Time                        `json:"date" validate:"required"`
	TotalPrice  float64                          `json:"total_price" validate:"min=0"`
	ReceiveDate *time.Time                       `json:"receive_date,omitempty"`
	Status      string                           `json:"status" validate:"required"`
	EmployeeID  uuid.UUID                        `json:"employee_id" validate:"required"`
	AgentID     uuid.UUID                        `json:"agent_id" validate:"required"`
	Details     []MedicineOrderDetailCreateRequest `json:"details"`
}

type MedicineOrderDetailResponse struct {
	MedicineOrderID uuid.UUID `json:"medicine_order_id"`
	MedicineID      uuid.UUID `json:"medicine_id"`
	Price           float64   `json:"price"`
	Qty             int       `json:"qty"`
	Remain          int       `json:"remain"`
}

type MedicineOrderResponse struct {
	ID           uuid.UUID                      `json:"id"`
	Date         time.Time                      `json:"date"`
	TotalPrice   float64                        `json:"total_price"`
	ReceiveDate  *time.Time                     `json:"receive_date,omitempty"`
	Status       string                         `json:"status"`
	EmployeeID   uuid.UUID                      `json:"employee_id"`
	AgentID      uuid.UUID                      `json:"agent_id"`
	Details      []MedicineOrderDetailResponse  `json:"details,omitempty"`
	CreatedAt    time.Time                      `json:"created_at"`
	UpdatedAt    time.Time                      `json:"updated_at"`
}

func ToMedicineOrderDetailResponse(mod *domain.MedicineOrderDetail) *MedicineOrderDetailResponse {
	return &MedicineOrderDetailResponse{
		MedicineOrderID: mod.MedicineOrderID,
		MedicineID:      mod.MedicineID,
		Price:           mod.Price,
		Qty:             mod.Qty,
		Remain:          mod.Remain,
	}
}

func ToMedicineOrderResponse(mo *domain.MedicineOrder) *MedicineOrderResponse {
	details := make([]MedicineOrderDetailResponse, len(mo.MedicineOrderDetails))
	for i, d := range mo.MedicineOrderDetails {
		details[i] = *ToMedicineOrderDetailResponse(&d)
	}

	return &MedicineOrderResponse{
		ID:          mo.ID,
		Date:        mo.Date,
		TotalPrice:  mo.TotalPrice,
		ReceiveDate: mo.ReceiveDate,
		Status:      string(mo.Status),
		EmployeeID:  mo.EmployeeID,
		AgentID:     mo.AgentID,
		Details:     details,
		CreatedAt:   mo.CreatedAt,
		UpdatedAt:   mo.UpdatedAt,
	}
}

func ToMedicineOrderDomain(req *MedicineOrderCreateRequest, orderID uuid.UUID) *domain.MedicineOrder {
	details := make([]domain.MedicineOrderDetail, len(req.Details))
	for i, d := range req.Details {
		details[i] = domain.MedicineOrderDetail{
			MedicineOrderID: orderID,
			MedicineID:      d.MedicineID,
			Price:           d.Price,
			Qty:             d.Qty,
			Remain:          d.Remain,
		}
	}

	return &domain.MedicineOrder{
		Date:               req.Date,
		TotalPrice:         req.TotalPrice,
		ReceiveDate:        req.ReceiveDate,
		Status:             domain.MedicineOrderStatus(req.Status),
		EmployeeID:         req.EmployeeID,
		AgentID:            req.AgentID,
		MedicineOrderDetails: details,
	}
}
