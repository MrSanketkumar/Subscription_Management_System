package main

import (
	"Subscription_Management_System/config"
	"Subscription_Management_System/handlers"
	"Subscription_Management_System/models"
	"Subscription_Management_System/repositories"
	"Subscription_Management_System/routes"
	"Subscription_Management_System/services"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()
	cfg.DB.AutoMigrate(&models.Subscription{})

	// Set up repository, service, and handler
	subscriptionRepo := repositories.NewSubscriptionRepository(cfg.DB)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)

	// Set up router and start the server
	router := routes.SetupRouter(subscriptionHandler)
	router.Run(":8080")
}
