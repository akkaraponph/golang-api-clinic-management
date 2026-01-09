package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MaterialOrderDetailCreateRequest struct {
	MaterialID uuid.UUID `json:"material_id" validate:"required"`
	Qty        int       `json:"qty" validate:"required,min=1"`
	Price      float64   `json:"price" validate:"min=0"`
}

type MaterialOrderCreateRequest struct {
	TotalPrice   float64                          `json:"total_price" validate:"min=0"`
	PurchaseDate *time.Time                       `json:"purchase_date,omitempty"`
	ReceiveDate  *time.Time                       `json:"receive_date,omitempty"`
	Status       string                           `json:"status" validate:"required"`
	EmployeeID   uuid.UUID                        `json:"employee_id" validate:"required"`
	AgentID      uuid.UUID                        `json:"agent_id" validate:"required"`
	Details      []MaterialOrderDetailCreateRequest `json:"details"`
}

type MaterialOrderDetailUpdateRequest struct {
	MaterialID *uuid.UUID `json:"material_id,omitempty"`
	Qty        *int       `json:"qty,omitempty" validate:"omitempty,min=1"`
	Price      *float64   `json:"price,omitempty" validate:"omitempty,min=0"`
}

type MaterialOrderUpdateRequest struct {
	TotalPrice   *float64  `json:"total_price,omitempty" validate:"omitempty,min=0"`
	PurchaseDate *time.Time `json:"purchase_date,omitempty"`
	ReceiveDate  *time.Time `json:"receive_date,omitempty"`
	Status       string    `json:"status"`
	EmployeeID   *uuid.UUID `json:"employee_id,omitempty"`
	AgentID      *uuid.UUID `json:"agent_id,omitempty"`
}

type MaterialOrderDetailResponse struct {
	ID             uuid.UUID `json:"id"`
	MaterialOrderID uuid.UUID `json:"material_order_id"`
	MaterialID     uuid.UUID `json:"material_id"`
	Qty            int       `json:"qty"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MaterialOrderResponse struct {
	ID            uuid.UUID                      `json:"id"`
	TotalPrice    float64                        `json:"total_price"`
	PurchaseDate  *time.Time                     `json:"purchase_date,omitempty"`
	ReceiveDate   *time.Time                     `json:"receive_date,omitempty"`
	Status        string                         `json:"status"`
	EmployeeID    uuid.UUID                      `json:"employee_id"`
	AgentID       uuid.UUID                      `json:"agent_id"`
	Details       []MaterialOrderDetailResponse  `json:"details,omitempty"`
	CreatedAt     time.Time                      `json:"created_at"`
	UpdatedAt     time.Time                      `json:"updated_at"`
}

func ToMaterialOrderDetailResponse(mod *domain.MaterialOrderDetail) *MaterialOrderDetailResponse {
	return &MaterialOrderDetailResponse{
		ID:             mod.ID,
		MaterialOrderID: mod.MaterialOrderID,
		MaterialID:     mod.MaterialID,
		Qty:            mod.Qty,
		Price:          mod.Price,
		CreatedAt:      mod.CreatedAt,
		UpdatedAt:      mod.UpdatedAt,
	}
}

func ToMaterialOrderResponse(mo *domain.MaterialOrder) *MaterialOrderResponse {
	details := make([]MaterialOrderDetailResponse, len(mo.MaterialOrderDetails))
	for i, d := range mo.MaterialOrderDetails {
		details[i] = *ToMaterialOrderDetailResponse(&d)
	}

	return &MaterialOrderResponse{
		ID:           mo.ID,
		TotalPrice:   mo.TotalPrice,
		PurchaseDate: mo.PurchaseDate,
		ReceiveDate:  mo.ReceiveDate,
		Status:       string(mo.Status),
		EmployeeID:   mo.EmployeeID,
		AgentID:      mo.AgentID,
		Details:      details,
		CreatedAt:    mo.CreatedAt,
		UpdatedAt:    mo.UpdatedAt,
	}
}

func ToMaterialOrderDomain(req *MaterialOrderCreateRequest) *domain.MaterialOrder {
	details := make([]domain.MaterialOrderDetail, len(req.Details))
	for i, d := range req.Details {
		details[i] = domain.MaterialOrderDetail{
			MaterialID: d.MaterialID,
			Qty:        d.Qty,
			Price:      d.Price,
		}
	}

	return &domain.MaterialOrder{
		TotalPrice:            req.TotalPrice,
		PurchaseDate:          req.PurchaseDate,
		ReceiveDate:           req.ReceiveDate,
		Status:                domain.MaterialOrderStatus(req.Status),
		EmployeeID:            req.EmployeeID,
		AgentID:               req.AgentID,
		MaterialOrderDetails:  details,
	}
}
