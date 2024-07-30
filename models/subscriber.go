package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscriber struct {
	gorm.Model
	// ID                  uuid.UUID `gorm:"type:uuid;"`
	Name                string    `json:"name"`
	SubscribedToChannel string    `json:"subscribedToChannel"`
	SubscriberDate      time.Time `json:"subscriberDate"`
}

var SubsCriber Subscriber
