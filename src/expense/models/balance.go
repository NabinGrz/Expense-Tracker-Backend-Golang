package expenseModels

import (
	"github.com/google/uuid"
)

type Balance struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CashAmount float64   `gorm:"not null"  json:"cash_amount"`
	BankAmount float64   `gorm:"not null"  json:"bank_amount"`
}
