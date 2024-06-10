package models

import (
	"time"
)

// subscription model
type Subscription struct {
	SubscriberID      int       `gorm:"column:subscriberid" json:"subscriberId"`
	SubscriberName    string    `gorm:"column:subscribername" json:"subscriberName"`
	SubscriberCountry string    `gorm:"column:subscribercountry" json:"subscriberCountry"`
	SubscriptionDate  time.Time `gorm:"column:subscriptiondate" json:"subscriptionDate"`
	SubscriptionID    int       `gorm:"column:subscriptionid" json:"subscriptionId"`
}
