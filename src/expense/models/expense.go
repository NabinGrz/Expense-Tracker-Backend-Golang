package expenseModels

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name       string    `gorm:"not null"  json:"name"`
	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time `gorm:"not null"  json:"created_at"`
	Amount     float64   `gorm:"not null"  json:"amount"`
	IsCash     bool      `gorm:"not null"  json:"is_cash"`
}
