package repository

import (
	"gorm.io/gorm"
	ruser "shiva/shiva-auth/internal/accounts/repository"
	"shiva/shiva-auth/internal/orders"
	rprod "shiva/shiva-auth/internal/products/repository"
	"time"
)

type Transactions struct {
	gorm.Model
	UserId              uint
	Users               ruser.Users `gorm:"foreignKey:user_id"`
	ProductId           uint
	Products            rprod.Products `gorm:"foreignKey:product_id"`
	DetailTransactionId uint
	DetailTransactions  DetailTransactions `gorm:"foreignKey:transaction_id"`
	Status              string
	SuccessDateTime     time.Time
	PendingDateTime     time.Time
	FailDateTime        time.Time
	ExpirationPayment   time.Time
	TotalPrice          int
	AccountNumber       string
	BankCode            string
	CreatedAt           time.Time `gorm:"autoCreateTime"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`
}

type DetailTransactions struct {
	gorm.Model
	Sku                       string
	Name                      string
	AdminFee                  int
	Price                     int
	DetailUniqueUser          string
	DetailUniqueValue         string
	DetailProductClassName    string
	DetailProductClassImage   string
	DetailProductClassTax     int
	DetailProductCategoryName string
}

func FromDomain(orders.Domain) Transactions {
	panic("")
}
