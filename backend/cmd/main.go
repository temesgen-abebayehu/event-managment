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
	paymentUsecase := usecase.NewPaymentUsecase(ticketRepo, orderRepo, userRepo, chapaClient, cfg.ChapaCallbackURL, cfg.ChapaReturnURL)

	// Initialize handlers
	authHandler := delivery.NewAuthHandler(authUsecase)
	uploadHandler := delivery.NewUploadHandler(uploadUsecase, cfg.JWTSecret)
	paymentHandler := delivery.NewPaymentHandler(paymentUsecase)
	tagHandler := delivery.NewTagHandler(hasuraClient, cfg.JWTSecret)

	// Setup routes
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(corsMiddleware)
	
	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	
	// Auth actions
	r.HandleFunc("/actions/signup", authHandler.Signup).Methods("POST", "OPTIONS")
	r.HandleFunc("/actions/login", authHandler.Login).Methods("POST", "OPTIONS")
	
	// Upload actions
	r.HandleFunc("/actions/upload", uploadHandler.UploadFiles).Methods("POST", "OPTIONS")
	r.HandleFunc("/actions/delete-files", uploadHandler.DeleteFiles).Methods("POST", "OPTIONS")
	
	// Tag actions
	r.HandleFunc("/actions/link-tags", tagHandler.LinkTags).Methods("POST", "OPTIONS")
	
	// Payment actions
	r.HandleFunc("/actions/initiate-payment", paymentHandler.InitiatePayment).Methods("POST", "OPTIONS")
	r.HandleFunc("/actions/verify-payment", paymentHandler.VerifyPayment).Methods("POST", "OPTIONS")
	
	// Chapa callback
	r.HandleFunc("/webhook/chapa", paymentHandler.PaymentCallback).Methods("POST")

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, r))
}

// corsMiddleware adds CORS headers so the frontend can reach the backend.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from common frontend origins
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
