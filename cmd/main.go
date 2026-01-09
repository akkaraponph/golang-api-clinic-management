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
	appointmentRepo := repositories.NewAppointmentRepository(db)
	materialRepo := repositories.NewMaterialRepository(db)
	materialOrderRepo := repositories.NewMaterialOrderRepository(db)
	prescriptionRepo := repositories.NewPrescriptionRepository(db)
	doctorRepo := repositories.NewDoctorRepository(db)
	agentRepo := repositories.NewAgentRepository(db)
	medicineTypeRepo := repositories.NewMedicineTypeRepository(db)
	medicineOrderRepo := repositories.NewMedicineOrderRepository(db)
	materialDisburseRepo := repositories.NewMaterialDisburseRepository(db)
	medicineDisburseRepo := repositories.NewMedicineDisburseRepository(db)
	referralRepo := repositories.NewReferralRepository(db)
	feedbackRepo := repositories.NewFeedbackRepository(db)
	insuranceRepo := repositories.NewInsuranceRepository(db)

	// Initialize PDF generator
	pdfGenerator := pdf.NewReceiptGenerator()

	// Initialize services
	employeeService := services.NewEmployeeService(employeeRepo)
	patientService := services.NewPatientService(patientRepo)
	receiptService := services.NewReceiptService(receiptRepo, pdfGenerator)
	courseService := services.NewCourseService(courseRepo)
	medicineService := services.NewMedicineService(medicineRepo)
	medicineStockService := services.NewMedicineStockService(medicineStockRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo)
	materialService := services.NewMaterialService(materialRepo)
	materialOrderService := services.NewMaterialOrderService(materialOrderRepo)
	prescriptionService := services.NewPrescriptionService(prescriptionRepo)
	doctorService := services.NewDoctorService(doctorRepo)
	agentService := services.NewAgentService(agentRepo)
	medicineTypeService := services.NewMedicineTypeService(medicineTypeRepo)
	medicineOrderService := services.NewMedicineOrderService(medicineOrderRepo)
	materialDisburseService := services.NewMaterialDisburseService(materialDisburseRepo)
	medicineDisburseService := services.NewMedicineDisburseService(medicineDisburseRepo)
	referralService := services.NewReferralService(referralRepo)
	feedbackService := services.NewFeedbackService(feedbackRepo)
	insuranceService := services.NewInsuranceService(insuranceRepo)

	// Initialize handlers
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	patientHandler := handlers.NewPatientHandler(patientService)
	receiptHandler := handlers.NewReceiptHandler(receiptService)
	courseHandler := handlers.NewCourseHandler(courseService)
	medicineHandler := handlers.NewMedicineHandler(medicineService)
	medicineStockHandler := handlers.NewMedicineStockHandler(medicineStockService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	materialHandler := handlers.NewMaterialHandler(materialService)
	materialOrderHandler := handlers.NewMaterialOrderHandler(materialOrderService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService)
	doctorHandler := handlers.NewDoctorHandler(doctorService)
	agentHandler := handlers.NewAgentHandler(agentService)
	medicineTypeHandler := handlers.NewMedicineTypeHandler(medicineTypeService)
	medicineOrderHandler := handlers.NewMedicineOrderHandler(medicineOrderService)
	materialDisburseHandler := handlers.NewMaterialDisburseHandler(materialDisburseService)
	medicineDisburseHandler := handlers.NewMedicineDisburseHandler(medicineDisburseService)
	referralHandler := handlers.NewReferralHandler(referralService)
	feedbackHandler := handlers.NewFeedbackHandler(feedbackService)
	insuranceHandler := handlers.NewInsuranceHandler(insuranceService)

	// Initialize Fiber app
	params := configs.NewFiberHttpServiceParams()
	app := configs.NewFiberHTTPService(params)

	// Setup routes
	api := app.Group("/api/" + configs.APP_API_VERSION)
	routes.SetupEmployeeRoutes(api, employeeHandler)
	routes.SetupPatientRoutes(api, patientHandler)
	routes.SetupReceiptRoutes(api, receiptHandler)
	routes.SetupCourseRoutes(api, courseHandler)
	routes.SetupMedicineRoutes(api, medicineHandler)
	routes.SetupMedicineStockRoutes(api, medicineStockHandler)
	routes.SetupAppointmentRoutes(api, appointmentHandler)
	routes.SetupMaterialRoutes(api, materialHandler)
	routes.SetupMaterialOrderRoutes(api, materialOrderHandler)
	routes.SetupPrescriptionRoutes(api, prescriptionHandler)
	routes.SetupDoctorRoutes(api, doctorHandler)
	routes.SetupAgentRoutes(api, agentHandler)
	routes.SetupMedicineTypeRoutes(api, medicineTypeHandler)
	routes.SetupMedicineOrderRoutes(api, medicineOrderHandler)
	routes.SetupMaterialDisburseRoutes(api, materialDisburseHandler)
	routes.SetupMedicineDisburseRoutes(api, medicineDisburseHandler)
	routes.SetupReferralRoutes(api, referralHandler)
	routes.SetupFeedbackRoutes(api, feedbackHandler)
	routes.SetupInsuranceRoutes(api, insuranceHandler)

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
