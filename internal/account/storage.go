package account

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/petradevsteam/sms/internal/storage"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) GetByID(id int) (*Account, error) {

	account := new(Account)

	err := s.DB.QueryRow("SELECT id, name FROM accounts WHERE id = ?", id).Scan(&account.ID, &account.Name)

	if err != nil {
		return nil, &storage.ErrorRecordNotFound{Message: fmt.Sprintf("couldn't find account by id %v", id)}
	}

	return account, nil
}

func (s *Storage) Get(ctx context.Context, limit int, offset int, q string) ([]*Account, error) {

	accounts := make([]*Account, 0)

	rows, err := s.DB.QueryContext(ctx, "SELECT id, name FROM accounts LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, &account.Name)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *Storage) Delete(ctx context.Context, account *Account) error {
	sql := "DELETE FROM accounts WHERE id = ?"
	result, err := s.DB.ExecContext(ctx, sql, account.ID)
	log.Println(result)
	if err != nil {
		return err
	}
	return nil
}
