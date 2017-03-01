package account

import "sync"

const testVersion = 1

// Account stores the bank account balance and whether or not the account is
// still open
type Account struct {
	balance int64
	open    bool
	mutex   sync.Mutex
}

// Open creates a new bank account with the initial deposit given.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	var a Account
	a.open = true
	a.balance = initialDeposit

	return &a
}

// Deposit changes the balance for the account if it is still open.  If amount
// is negative then it will count as a withdrawal.  If the account is closed
// or would go negative, ok will be false.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// Check if the account is closed, or if the "deposit" is negative and
	// would overdraw the account.
	if !a.open || (a.balance+amount) < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

// Balance returns the account balance.  ok will be false if the account is closed.
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.open {
		return 0, false
	}

	return a.balance, true
}

// Close will mark the account as closed and return payout as the final
// balance to be paid out.  It will return ok as false if the account is
// already closed.
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.open {
		return 0, false
	}

	balance := a.balance
	a.balance = 0
	a.open = false

	return balance, true
}
