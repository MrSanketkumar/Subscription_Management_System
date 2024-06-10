package handlers

import (
	"Subscription_Management_System/constants"
	"Subscription_Management_System/models"
	"Subscription_Management_System/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service services.SubscriptionService
}

// Initialize subscription handler
func NewSubscriptionHandler(service services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// Create subscription in database
func (h *SubscriptionHandler) CreateSubscription(c *gin.Context) {
	var subscription models.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{constants.ErrorMessage: constants.ErrInvalidInput})
		return
	}

	if err := h.service.CreateSubscription(&subscription); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{constants.ErrorMessage: constants.ErrDatabase})
		return
	}

	c.JSON(http.StatusOK, gin.H{constants.Message: constants.Success})
}

// Get all subscriptions from database
func (h *SubscriptionHandler) GetAllSubscriptions(c *gin.Context) {
	subscriptions, err := h.service.GetAllSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{constants.ErrorMessage: constants.ErrDatabase})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

// Get total number of subscribers in database
func (h *SubscriptionHandler) GetTotalSubscribers(c *gin.Context) {
	count, err := h.service.GetTotalSubscribers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{constants.ErrorMessage: constants.ErrDatabase})
		return
	}
	c.JSON(http.StatusOK, gin.H{constants.TotalSubscribers: count})
}

// Get country with most subscribers in database
func (h *SubscriptionHandler) GetLongestSubscriptionDuration(c *gin.Context) {
	duration, err := h.service.GetLongestSubscriptionDuration()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{constants.ErrorMessage: constants.ErrDatabase})
		return
	}
	c.JSON(http.StatusOK, gin.H{constants.LongestDuration: duration.String()})
}

// Get country with most subscribers
func (h *SubscriptionHandler) GetCountryWithMostSubscribers(c *gin.Context) {
	country, err := h.service.GetCountryWithMostSubscribers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{constants.ErrorMessage: constants.ErrDatabase})
		return
	}
	c.JSON(http.StatusOK, gin.H{constants.MostSubscribers: country})
}
