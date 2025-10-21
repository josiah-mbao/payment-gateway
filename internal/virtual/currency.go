// internal/virtual/currency.go
package virtual

import "sync"

type JuiceCoinEconomy struct {
    mu sync.RWMutex
    userBalances map[string]int64 // userID -> JUC balance
    exchangeRates map[string]float64
    totalSupply int64
}

func NewJuiceCoinEconomy() *JuiceCoinEconomy {
    return &JuiceCoinEconomy{
        userBalances: make(map[string]int64),
        exchangeRates: map[string]float64{
            "JUC_KES": 100.0, // 1 JUC = 100 virtual KES
            "JUC_NGN": 50.0,  // 1 JUC = 50 virtual NGN
            "JUC_GHS": 80.0,  // 1 JUC = 80 virtual GHS  
            "JUC_ZAR": 20.0,  // 1 JUC = 20 virtual ZAR
            "JUC_USD": 1.0,   // 1 JUC = 1 virtual USD
        },
        totalSupply: 1_000_000_000, // 1 billion JUC
    }
}

func (j *JuiceCoinEconomy) CreateUserWallet(userID string) error {
    j.mu.Lock()
    defer j.mu.Unlock()
    
    initialBalance := int64(10_000) // 10,000 JUC for new users
    j.userBalances[userID] = initialBalance
    return nil
}

func (j *JuiceCoinEconomy) GetBalance(userID string) (int64, error) {
    j.mu.RLock()
    defer j.mu.RUnlock()
    
    balance, exists := j.userBalances[userID]
    if !exists {
        return 0, fmt.Errorf("user wallet not found")
    }
    return balance, nil
}

func (j *JuiceCoinEconomy) Transfer(senderID, receiverID string, amount int64) error {
    j.mu.Lock()
    defer j.mu.Unlock()
    
    senderBalance, sExists := j.userBalances[senderID]
    receiverBalance, rExists := j.userBalances[receiverID]
    
    if !sExists || !rExists {
        return fmt.Errorf("one or more wallets not found")
    }
    
    if senderBalance < amount {
        return fmt.Errorf("insufficient Juice Coins")
    }
    
    j.userBalances[senderID] = senderBalance - amount
    j.userBalances[receiverID] = receiverBalance + amount
    
    return nil
}
