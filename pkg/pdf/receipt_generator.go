package pdf

import (
	"bytes"
	"fmt"
	"time"

	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/phpdave11/gofpdf"
)

type ReceiptGenerator struct{}

func NewReceiptGenerator() *ReceiptGenerator {
	return &ReceiptGenerator{}
}

func (g *ReceiptGenerator) Generate(receipt *domain.Receipt) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// Set font
	pdf.SetFont("Arial", "B", 16)
	
	// Header - Clinic Name
	pdf.Cell(190, 10, "Clinic Management System")
	pdf.Ln(10)
	
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 10, "Receipt")
	pdf.Ln(15)
	
	// Receipt Info
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(50, 8, "Receipt ID: ")
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(140, 8, receipt.ID.String())
	pdf.Ln(6)
	
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(50, 8, "Date: ")
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(140, 8, receipt.Date.Format("2006-01-02 15:04:05"))
	pdf.Ln(10)
	
	// Patient Information
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 10, "Patient Information")
	pdf.Ln(8)
	
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(95, 8, fmt.Sprintf("Name: %s %s", receipt.Patient.Name, receipt.Patient.Surname))
	pdf.Cell(95, 8, fmt.Sprintf("Phone: %s", receipt.Patient.Phone))
	pdf.Ln(6)
	
	if receipt.Patient.IDCard != "" {
		pdf.Cell(95, 8, fmt.Sprintf("ID Card: %s", receipt.Patient.IDCard))
		pdf.Ln(6)
	}
	
	pdf.Ln(5)
	
	// Course/Medical Information
	if receipt.Course != nil {
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(190, 10, "Treatment Details")
		pdf.Ln(8)
		
		pdf.SetFont("Arial", "", 10)
		if receipt.Course.Weight != nil {
			pdf.Cell(95, 8, fmt.Sprintf("Weight: %.2f kg", *receipt.Course.Weight))
		}
		if receipt.Course.Height != nil {
			pdf.Cell(95, 8, fmt.Sprintf("Height: %.2f cm", *receipt.Course.Height))
		}
		pdf.Ln(6)
		
		if receipt.Course.Systolic != nil && receipt.Course.Diastolic != nil {
			pdf.Cell(95, 8, fmt.Sprintf("Blood Pressure: %d/%d mmHg", *receipt.Course.Systolic, *receipt.Course.Diastolic))
		}
		if receipt.Course.HeartRate != nil {
			pdf.Cell(95, 8, fmt.Sprintf("Heart Rate: %d bpm", *receipt.Course.HeartRate))
		}
		pdf.Ln(6)
		
		if receipt.Course.Diagnose != "" {
			pdf.Cell(190, 8, fmt.Sprintf("Diagnosis: %s", receipt.Course.Diagnose))
			pdf.Ln(6)
		}
		
		if receipt.Course.TotalPrice > 0 {
			pdf.SetFont("Arial", "B", 10)
			pdf.Cell(95, 8, fmt.Sprintf("Treatment Fee: %.2f", receipt.Course.TotalPrice))
			pdf.Ln(6)
		}
		
		pdf.Ln(5)
	}
	
	// Medicine Information
	if receipt.MedicineDisburse != nil && len(receipt.MedicineDisburse.MedicineDisburseDetails) > 0 {
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(190, 10, "Medicines")
		pdf.Ln(8)
		
		// Table header
		pdf.SetFont("Arial", "B", 9)
		pdf.Cell(80, 8, "Medicine")
		pdf.Cell(30, 8, "Quantity")
		pdf.Cell(40, 8, "Dosage")
		pdf.Cell(40, 8, "Price")
		pdf.Ln(8)
		
		pdf.SetFont("Arial", "", 9)
		for _, detail := range receipt.MedicineDisburse.MedicineDisburseDetails {
			medicineName := "N/A"
			if detail.MedicineStock.Medicine.Name != "" {
				medicineName = detail.MedicineStock.Medicine.Name
			}
			
			pdf.Cell(80, 8, medicineName)
			pdf.Cell(30, 8, fmt.Sprintf("%d %s", detail.Qty, detail.Unit))
			pdf.Cell(40, 8, detail.Dosage)
			pdf.Cell(40, 8, fmt.Sprintf("%.2f", detail.Price))
			pdf.Ln(8)
		}
		
		pdf.Ln(5)
	}
	
	// Total
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(150, 10, "Total: ")
	pdf.Cell(40, 10, fmt.Sprintf("%.2f", receipt.TotalPrice))
	pdf.Ln(15)
	
	// Footer
	pdf.SetFont("Arial", "I", 8)
	pdf.Cell(190, 8, fmt.Sprintf("Generated on: %s", time.Now().Format("2006-01-02 15:04:05")))
	pdf.Ln(5)
	pdf.Cell(190, 8, "Thank you for your visit!")
	
	// Generate PDF bytes
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	
	return buf.Bytes(), nil
}
