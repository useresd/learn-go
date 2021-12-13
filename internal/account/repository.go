package account

import "context"

type Repository interface {
	GetByID(int) (*Account, error)
	Get(context.Context, int, int, string) ([]*Account, error)
	Delete(context.Context, *Account) error
}
