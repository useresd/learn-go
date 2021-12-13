package account

import (
	"context"
	"errors"
)

type Mock struct {
	Accounts []*Account
	Account  *Account
}

func (m *Mock) GetByID(id int) (*Account, error) {
	if id != m.Account.ID {
		return nil, errors.New("account not found")
	}
	return m.Account, nil
}

func (m *Mock) Get(ctx context.Context, limit int, offset int, q string) ([]*Account, error) {
	return m.Accounts, nil
}

func (m *Mock) Delete(ctx context.Context, account *Account) error {
	m.Account = &Account{}
	return nil
}
