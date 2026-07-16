package main

import (
	"event-management-backend/config"
	"event-management-backend/delivery"
	"event-management-backend/infrastructure/chapa"
	"event-management-backend/infrastructure/cloudinary"
	"event-management-backend/infrastructure/hasura"
	"event-management-backend/repository"
	"event-management-backend/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize infrastructure
	hasuraClient := hasura.NewClient(cfg.HasuraEndpoint, cfg.HasuraAdminSecret)
	
	cloudinaryClient, err := cloudinary.NewClient(cfg.CloudinaryName, cfg.CloudinaryAPIKey, cfg.CloudinarySecret)
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary: %v", err)
	}
	
	chapaClient := chapa.NewClient(cfg.ChapaSecretKey)

	// Initialize repositories
	userRepo := repository.NewUserRepository(hasuraClient)
	ticketRepo := repository.NewTicketRepository(hasuraClient)
	orderRepo := repository.NewOrderRepository(hasuraClient)

	// Initialize usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)
	uploadUsecase := usecase.NewUploadUsecase(cloudinaryClient)
	paymentUsecase := usecase.NewPaymentUsecase(ticketRepo, orderRepo, chapaClient, cfg.ChapaCallbackURL, cfg.ChapaReturnURL)

	// Initialize handlers
	authHandler := delivery.NewAuthHandler(authUsecase)
	uploadHandler := delivery.NewUploadHandler(uploadUsecase)
	paymentHandler := delivery.NewPaymentHandler(paymentUsecase)

	// Setup routes
	r := mux.NewRouter()
	
	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	
	// Auth actions
	r.HandleFunc("/actions/signup", authHandler.Signup).Methods("POST")
	r.HandleFunc("/actions/login", authHandler.Login).Methods("POST")
	
	// Upload actions
	r.HandleFunc("/actions/upload", uploadHandler.UploadFiles).Methods("POST")
	r.HandleFunc("/actions/delete-files", uploadHandler.DeleteFiles).Methods("POST")
	
	// Payment actions
	r.HandleFunc("/actions/initiate-payment", paymentHandler.InitiatePayment).Methods("POST")
	r.HandleFunc("/actions/verify-payment", paymentHandler.VerifyPayment).Methods("POST")
	
	// Chapa callback
	r.HandleFunc("/webhook/chapa", paymentHandler.PaymentCallback).Methods("POST")

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, r))
}
