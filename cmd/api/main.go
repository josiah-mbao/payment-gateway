// cmd/api/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "pesaflow/internal/api"
    "pesaflow/internal/database"
    "pesaflow/internal/interledger"
    "pesaflow/internal/virtual"
)

func main() {
    // Database
    db := database.NewPostgresDB()
    
    // Virtual economy
    juiceCoin := virtual.NewJuiceCoinEconomy()
    
    // Real Interledger client (but for virtual settlements)
    ilp := interledger.NewRafikiClient()
    
    // API handlers
    handler := api.NewHandler(db, juiceCoin, ilp)
    
    r := gin.Default()
    
    // User routes
    r.POST("/api/v1/users", handler.CreateUser)
    r.GET("/api/v1/users/:id", handler.GetUser)
    
    // Merchant routes  
    r.POST("/api/v1/merchants", handler.CreateMerchant)
    r.GET("/api/v1/merchants/:id", handler.GetMerchant)
    
    // Payment link routes
    r.POST("/api/v1/payment-links", handler.CreatePaymentLink)
    r.GET("/api/v1/payment-links/:id", handler.GetPaymentLink)
    
    // Payment routes
    r.POST("/api/v1/payments", handler.CreatePayment)
    r.GET("/api/v1/payments/:id", handler.GetPayment)
    
    // Juice Coin wallet routes
    r.GET("/api/v1/users/:id/wallet", handler.GetWalletBalance)
    
    // Webhook for Interledger
    r.POST("/webhooks/interledger", handler.HandleILPWebhook)
    
    // Health and metrics
    r.GET("/health", handler.HealthCheck)
    r.GET("/metrics", handler.Metrics)
    
    r.Run(":8080")
}
