package entity

import "time"

// TransactionStatus: pending, approved, denied
type TransactionStatus string

const (
	TransactionStatusPending  TransactionStatus = "pending"
	TransactionStatusApproved TransactionStatus = "approved"
	TransactionStatusDenied   TransactionStatus = "denied"
)

func (t TransactionStatus) String() string {
	return string(t)
}

type Transaction struct {
	ID                     uint32
	CreatedAt              time.Time
	UpdatedAt              time.Time
	UUID                   string
	AccountID              uint32
	Status                 TransactionStatus
	Name                   string
	CreditCardNumber       string
	CreditCardSecurityCode string
	CreditCardExpires      time.Time
	Amount                 float64
}
