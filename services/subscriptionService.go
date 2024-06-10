package services

import (
	"Subscription_Management_System/models"
	"Subscription_Management_System/repositories"
	"time"
)

// subscription service interface
type SubscriptionService interface {
	CreateSubscription(subscription *models.Subscription) error
	GetAllSubscriptions() ([]models.Subscription, error)
	GetTotalSubscribers() (int64, error)
	GetLongestSubscriptionDuration() (time.Duration, error)
	GetCountryWithMostSubscribers() (string, error)
}

// subscription service struct
type subscriptionService struct {
	repository repositories.SubscriptionRepository
}

// NewSubscriptionService creates a new subscription service
func NewSubscriptionService(repo repositories.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repository: repo}
}

// subscription service methods
func (s *subscriptionService) CreateSubscription(subscription *models.Subscription) error {
	return s.repository.CreateSubscription(subscription)
}

// Get all subscriptions from database
func (s *subscriptionService) GetAllSubscriptions() ([]models.Subscription, error) {
	return s.repository.GetAllSubscriptions()
}

// Get total subscribers in database
func (s *subscriptionService) GetTotalSubscribers() (int64, error) {
	return s.repository.GetTotalSubscribers()
}

// Get longest subscription duration in database
func (s *subscriptionService) GetLongestSubscriptionDuration() (time.Duration, error) {
	return s.repository.GetLongestSubscriptionDuration()
}

// / Get country with most subscribers in database
func (s *subscriptionService) GetCountryWithMostSubscribers() (string, error) {
	return s.repository.GetCountryWithMostSubscribers()
}
