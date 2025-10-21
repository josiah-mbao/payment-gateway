// internal/api/payments.go
package api

func (h *Handler) CreatePayment(c *gin.Context) {
    var req CreatePaymentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Get payment link
    var paymentLink PaymentLink
    if err := h.db.Preload("Merchant").First(&paymentLink, "id = ?", req.PaymentLinkID).Error; err != nil {
        c.JSON(404, gin.H{"error": "Payment link not found"})
        return
    }
    
    // Check if payer has enough Juice Coins
    payerBalance, err := h.juiceCoin.GetBalance(req.PayerID)
    if err != nil {
        c.JSON(400, gin.H{"error": "Payer wallet not found"})
        return
    }
    
    if payerBalance < paymentLink.Amount {
        c.JSON(400, gin.H{"error": "Insufficient Juice Coins"})
        return
    }
    
    // Create payment record
    payment := &Payment{
        ID:            generateID(),
        PaymentLinkID: req.PaymentLinkID,
        PayerID:       req.PayerID,
        Amount:        paymentLink.Amount,
        Currency:      "JUC",
        Status:        "pending",
        PayerCountry:  req.PayerCountry,
        PayeeCountry:  paymentLink.Merchant.Country,
        CreatedAt:     time.Now(),
    }
    
    if err := h.db.Create(payment).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create payment"})
        return
    }
    
    // Process virtual Interledger payment
    go h.processJuiceCoinPayment(payment, paymentLink.Merchant.UserID)
    
    c.JSON(201, gin.H{
        "payment_id": payment.ID,
        "status":     "processing",
        "message":    "Payment is being processed with Juice Coins",
    })
}

func (h *Handler) processJuiceCoinPayment(payment *Payment, merchantUserID string) {
    // Real Interledger integration with virtual accounts
    ilpResponse, err := h.ilp.SendVirtualPayment(InterledgerRequest{
        Source:      fmt.Sprintf("ilp.rafiki.money/juicecoin/%s", payment.PayerCountry),
        Destination: fmt.Sprintf("ilp.rafiki.money/juicecoin/%s", payment.PayeeCountry),
        Amount:      payment.Amount,
        AssetCode:   "JUC",
        Memo:        fmt.Sprintf("Juice Coin Payment: %s", payment.ID),
    })
    
    if err != nil {
        payment.Status = "failed"
    } else {
        // Transfer Juice Coins between users
        if err := h.juiceCoin.Transfer(payment.PayerID, merchantUserID, payment.Amount); err != nil {
            payment.Status = "failed"
        } else {
            payment.Status = "completed"
            payment.ILPTransaction = ilpResponse.TransactionID
            payment.CompletedAt = &time.Now()
            
            // Mark payment link as paid
            h.db.Model(&PaymentLink{}).Where("id = ?", payment.PaymentLinkID).Update("status", "paid")
        }
    }
    
    // Update payment status
    h.db.Save(payment)
}
