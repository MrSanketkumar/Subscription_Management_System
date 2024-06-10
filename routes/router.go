package routes

import (
	"Subscription_Management_System/constants"
	"Subscription_Management_System/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router
func SetupRouter(handler *handlers.SubscriptionHandler) *gin.Engine {
	r := gin.Default()

	r.POST(constants.CreateSubscriptionPath, handler.CreateSubscription)
	r.GET(constants.GetAllSubscriptionsPath, handler.GetAllSubscriptions)
	r.GET(constants.GetTotalSubscribersPath, handler.GetTotalSubscribers)
	r.GET(constants.GetLongestSubscriptionDurationPath, handler.GetLongestSubscriptionDuration)
	r.GET(constants.GetMostSubscribersPath, handler.GetCountryWithMostSubscribers)
	// r.StaticFS("/ui", http.Dir("./ui/index.html"))
	r.Static("/ui", "./ui")

	r.GET("/", func(c *gin.Context) {
		c.File("./ui/index.html")
	})
	return r
}
