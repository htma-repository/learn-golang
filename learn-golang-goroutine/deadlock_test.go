package learn_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBankAccount struct {
	Mutex       sync.Mutex
	AccountName string
	Balance     int
}

func (u *UserBankAccount) MutexLock() {
	u.Mutex.Lock()
}

func (u *UserBankAccount) MutexUnlock() {
	u.Mutex.Unlock()
}

func (u *UserBankAccount) ChangeAmount(amount int) {
	u.Balance += amount
}

func TransferBalance(user1 *UserBankAccount, user2 *UserBankAccount, amount int) {
	user1.Mutex.Lock()
	fmt.Println("Lock User-1, ", user1.AccountName)
	user1.ChangeAmount(-amount)

	time.Sleep(1 * time.Second)

	user2.Mutex.Lock()
	fmt.Println("Lock User-2, ", user2.AccountName)
	user2.ChangeAmount(amount)

	time.Sleep(1 * time.Second)

	user1.Mutex.Unlock()
	user2.Mutex.Unlock()

}

func TestDeadlock(t *testing.T) {
	User1 := UserBankAccount{AccountName: "Hutama", Balance: 500000}
	User2 := UserBankAccount{AccountName: "Trirahmanto", Balance: 500000}

	go TransferBalance(&User1, &User2, 100000)
	go TransferBalance(&User2, &User1, 100000)

	time.Sleep(5 * time.Second)
	fmt.Println("Final Bank balance user-1: ", User1.Balance)
	fmt.Println("Final Bank balance user-2: ", User2.Balance)
}
