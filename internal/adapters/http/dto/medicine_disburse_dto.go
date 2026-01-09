package dto

import (
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/google/uuid"
)

type MedicineDisburseDetailCreateRequest struct {
	MedicineStockID uuid.UUID  `json:"medicine_stock_id" validate:"required"`
	Price           float64    `json:"price" validate:"min=0"`
	Unit            string     `json:"unit" validate:"max=50"`
	Qty             int        `json:"qty" validate:"required,min=1"`
	Dosage          string     `json:"dosage" validate:"max=100"`
	AdminMethod     string     `json:"admin_method" validate:"max=100"`
	ApplyMethod     string     `json:"apply_method" validate:"max=100"`
	TimeOfAdmin     *time.Time `json:"time_of_admin,omitempty"`
}

type MedicineDisburseCreateRequest struct {
	PrescriptionID uuid.UUID                          `json:"prescription_id" validate:"required"`
	Date           time.Time                          `json:"date" validate:"required"`
	Details        []MedicineDisburseDetailCreateRequest `json:"details"`
}

type MedicineDisburseDetailResponse struct {
	MedicineDisburseID uuid.UUID  `json:"medicine_disburse_id"`
	MedicineStockID    uuid.UUID  `json:"medicine_stock_id"`
	Price              float64    `json:"price"`
	Unit               string     `json:"unit"`
	Qty                int        `json:"qty"`
	Dosage             string     `json:"dosage"`
	AdminMethod        string     `json:"admin_method"`
	ApplyMethod        string     `json:"apply_method"`
	TimeOfAdmin        *time.Time `json:"time_of_admin,omitempty"`
}

type MedicineDisburseResponse struct {
	ID            uuid.UUID                        `json:"id"`
	PrescriptionID uuid.UUID                       `json:"prescription_id"`
	Date          time.Time                        `json:"date"`
	Details       []MedicineDisburseDetailResponse `json:"details,omitempty"`
	CreatedAt     time.Time                        `json:"created_at"`
	UpdatedAt     time.Time                        `json:"updated_at"`
}

func ToMedicineDisburseDetailResponse(mdd *domain.MedicineDisburseDetail) *MedicineDisburseDetailResponse {
	return &MedicineDisburseDetailResponse{
		MedicineDisburseID: mdd.MedicineDisburseID,
		MedicineStockID:    mdd.MedicineStockID,
		Price:              mdd.Price,
		Unit:               mdd.Unit,
		Qty:                mdd.Qty,
		Dosage:             mdd.Dosage,
		AdminMethod:        mdd.AdminMethod,
		ApplyMethod:        mdd.ApplyMethod,
		TimeOfAdmin:        mdd.TimeOfAdmin,
	}
}

func ToMedicineDisburseResponse(md *domain.MedicineDisburse) *MedicineDisburseResponse {
	details := make([]MedicineDisburseDetailResponse, len(md.MedicineDisburseDetails))
	for i, d := range md.MedicineDisburseDetails {
		details[i] = *ToMedicineDisburseDetailResponse(&d)
	}

	return &MedicineDisburseResponse{
		ID:             md.ID,
		PrescriptionID: md.PrescriptionID,
		Date:           md.Date,
		Details:        details,
		CreatedAt:      md.CreatedAt,
		UpdatedAt:      md.UpdatedAt,
	}
}

func ToMedicineDisburseDomain(req *MedicineDisburseCreateRequest, disburseID uuid.UUID) *domain.MedicineDisburse {
	details := make([]domain.MedicineDisburseDetail, len(req.Details))
	for i, d := range req.Details {
		details[i] = domain.MedicineDisburseDetail{
			MedicineDisburseID: disburseID,
			MedicineStockID:    d.MedicineStockID,
			Price:              d.Price,
			Unit:               d.Unit,
			Qty:                d.Qty,
			Dosage:             d.Dosage,
			AdminMethod:        d.AdminMethod,
			ApplyMethod:        d.ApplyMethod,
			TimeOfAdmin:        d.TimeOfAdmin,
		}
	}

	return &domain.MedicineDisburse{
		PrescriptionID:            req.PrescriptionID,
		Date:                      req.Date,
		MedicineDisburseDetails:   details,
	}
}
