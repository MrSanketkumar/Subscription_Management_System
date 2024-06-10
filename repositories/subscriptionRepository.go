package repositories

import (
	"Subscription_Management_System/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

// subscription repository interface
type SubscriptionRepository interface {
	CreateSubscription(subscription *models.Subscription) error
	GetAllSubscriptions() ([]models.Subscription, error)
	GetTotalSubscribers() (int64, error)
	GetLongestSubscriptionDuration() (time.Duration, error)
	GetCountryWithMostSubscribers() (string, error)
}

// subscription repository struct
type subscriptionRepository struct {
	DB *gorm.DB
}

// NewSubscriptionRepository creates a new subscription repository
func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{DB: db}
}

// Create subscription in database
func (r *subscriptionRepository) CreateSubscription(subscription *models.Subscription) error {
	return r.DB.Create(subscription).Error
}

// Get all subscriptions from database
func (r *subscriptionRepository) GetAllSubscriptions() ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := r.DB.Find(&subscriptions).Error
	return subscriptions, err
}

// Get total subscribers from database
func (r *subscriptionRepository) GetTotalSubscribers() (int64, error) {
	var count int64
	err := r.DB.Model(&models.Subscription{}).Count(&count).Error
	return count, err
}

// Get longest subscription duration from database
func (r *subscriptionRepository) GetLongestSubscriptionDuration() (time.Duration, error) {
	var subscription models.Subscription
	err := r.DB.Order("subscriptiondate").First(&subscription).Error
	if err != nil {
		return 0, err
	}
	duration := time.Since(subscription.SubscriptionDate)
	return duration, nil
}

// Get country with most subscribers in database
func (r *subscriptionRepository) GetCountryWithMostSubscribers() (string, error) {
	var SubscriberCountry string
	var countries []string

	err := r.DB.Table("subscriptions").
		Select("subscribercountry").
		Group("subscribercountry").
		Order("count(*) desc").
		Limit(1).
		Pluck("subscribercountry", &countries).Error

	if err != nil {
		return "", err
	}

	if len(countries) == 0 {
		return "", errors.New("no country found with subscribers")
	}

	SubscriberCountry = countries[0]

	return SubscriberCountry, nil
}
