// internal/domain/models.go
package domain

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        string    `gorm:"primaryKey;size:36"`
    Email     string    `gorm:"uniqueIndex;not null"`
    Name      string    
    Country   string    `gorm:"size:2"` // KE, NG, GH, ZA, etc.
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Merchant struct {
    ID           string    `gorm:"primaryKey;size:36"`
    UserID       string    `gorm:"not null"`
    BusinessName string    `gorm:"not null"`
    Description  string    
    Country      string    `gorm:"size:2;not null"`
    Category     string    // tech, fashion, services, etc.
    Website      string    
    CreatedAt    time.Time
    UpdatedAt    time.Time
    
    User         User      `gorm:"foreignKey:UserID"`
}

type PaymentLink struct {
    ID          string    `gorm:"primaryKey;size:36"`
    MerchantID  string    `gorm:"not null"`
    Amount      int64     `gorm:"not null"` // Amount in JUC
    Currency    string    `gorm:"default:JUC"`
    Title       string    `gorm:"not null"`
    Description string    
    Status      string    `gorm:"default:active"` // active, paid, expired
    CreatedAt   time.Time
    ExpiresAt   *time.Time
    
    Merchant    Merchant  `gorm:"foreignKey:MerchantID"`
}

type Payment struct {
    ID             string    `gorm:"primaryKey;size:36"`
    PaymentLinkID  string    `gorm:"not null"`
    PayerID        string    `gorm:"not null"`
    Amount         int64     `gorm:"not null"`
    Currency       string    `gorm:"default:JUC"`
    Status         string    // pending, completed, failed
    PayerCountry   string    `gorm:"size:2"`
    PayeeCountry   string    `gorm:"size:2"`
    ILPTransaction string    // Real Interledger transaction ID
    RiskScore      float64   // ML fraud detection score
    CreatedAt      time.Time
    CompletedAt    *time.Time
    
    PaymentLink    PaymentLink `gorm:"foreignKey:PaymentLinkID"`
    Payer          User        `gorm:"foreignKey:PayerID"`
}
