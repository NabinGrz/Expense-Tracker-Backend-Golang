package expenseModels

import "github.com/google/uuid"

type Category struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name  string    `gorm:"not null"  json:"name"`
	Image string    `gorm:"not null"  json:"image"`
}
