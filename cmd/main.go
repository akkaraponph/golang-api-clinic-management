package main

import (
	"log"

	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/handlers"
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/http/routes"
	"github.com/billowdev/golang-api-clinic-management/internal/adapters/secondary/repositories"
	"github.com/billowdev/golang-api-clinic-management/internal/core/domain"
	"github.com/billowdev/golang-api-clinic-management/internal/core/services"
	"github.com/billowdev/golang-api-clinic-management/pkg/configs"
	"github.com/billowdev/golang-api-clinic-management/pkg/database"
	"github.com/billowdev/golang-api-clinic-management/pkg/pdf"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if configs.DB_RUN_MIGRATION {
		err = runMigrations(db)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Database migrations completed successfully")
	}

	// Initialize repositories
	employeeRepo := repositories.NewEmployeeRepository(db)
	patientRepo := repositories.NewPatientRepository(db)
	receiptRepo := repositories.NewReceiptRepository(db)
	courseRepo := repositories.NewCourseRepository(db)
	medicineRepo := repositories.NewMedicineRepository(db)
	medicineStockRepo := repositories.NewMedicineStockRepository(db)
	_ = repositories.NewAppointmentRepository(db)
	_ = repositories.NewMaterialRepository(db)
	_ = repositories.NewMaterialOrderRepository(db)
	_ = repositories.NewPrescriptionRepository(db)

	// Initialize PDF generator
	pdfGenerator := pdf.NewReceiptGenerator()

	// Initialize services
	employeeService := services.NewEmployeeService(employeeRepo)
	patientService := services.NewPatientService(patientRepo)
	receiptService := services.NewReceiptService(receiptRepo, pdfGenerator)
	_ = services.NewCourseService(courseRepo)
	_ = services.NewMedicineService(medicineRepo)
	_ = services.NewMedicineStockService(medicineStockRepo)

	// Initialize handlers
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	patientHandler := handlers.NewPatientHandler(patientService)
	receiptHandler := handlers.NewReceiptHandler(receiptService)

	// Initialize Fiber app
	params := configs.NewFiberHttpServiceParams()
	app := configs.NewFiberHTTPService(params)

	// Setup routes
	api := app.Group("/api/" + configs.APP_API_VERSION)
	routes.SetupEmployeeRoutes(api, employeeHandler)
	routes.SetupPatientRoutes(api, patientHandler)
	routes.SetupReceiptRoutes(api, receiptHandler)

	// Start server
	log.Printf("Starting server on port %s", params.Port)
	if err := app.Listen(":" + params.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func runMigrations(db *gorm.DB) error {
	// Run migrations for all domain models
	return db.AutoMigrate(
		&domain.Employee{},
		&domain.Doctor{},
		&domain.Patient{},
		&domain.Appointment{},
		&domain.Agent{},
		&domain.Material{},
		&domain.MaterialOrder{},
		&domain.MaterialOrderDetail{},
		&domain.MaterialDisburse{},
		&domain.MaterialDisburseDetail{},
		&domain.Medicine{},
		&domain.MedicineType{},
		&domain.MedicineOrder{},
		&domain.MedicineOrderDetail{},
		&domain.MedicineStock{},
		&domain.Course{},
		&domain.Prescription{},
		&domain.MedicineDisburse{},
		&domain.MedicineDisburseDetail{},
		&domain.Receipt{},
		&domain.Feedback{},
		&domain.Insurance{},
		&domain.Referral{},
		&domain.EmailLog{},
		&domain.Promotion{},
		&domain.Geography{},
		&domain.Province{},
		&domain.District{},
		&domain.Subdistrict{},
		&domain.SystemSetting{},
	)
}
