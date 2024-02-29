package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	BankRWMutex sync.RWMutex
	Balance     int
}

func (b *BankAccount) AddBalance(amount int) {
	b.BankRWMutex.Lock()
	b.Balance += amount
	b.BankRWMutex.Unlock()
}

func (b *BankAccount) GetBalance() int {
	b.BankRWMutex.RLock()
	balance := b.Balance
	b.BankRWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	HutamaBankAccount := BankAccount{
		Balance: 0,
	}

	for i := 1; i <= 1000; i++ {

		go func() {
			for j := 1; j <= 100; j++ {
				HutamaBankAccount.AddBalance(1)
				fmt.Println(HutamaBankAccount.GetBalance())
			}
		}()

	}

	time.Sleep(3 * time.Second)
	fmt.Println("Final Balance, ", HutamaBankAccount.GetBalance())
}
