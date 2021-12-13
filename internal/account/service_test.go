package account

import (
	"context"
	"testing"
)

func assertEqual(t *testing.T, x interface{}, y interface{}) {
	if x != y {
		t.Errorf("Expected %v got %v", x, y)
	}
}

func TestGet(t *testing.T) {
	r := &Mock{Account: &Account{ID: 1, Name: "Cash"}}
	s := &Service{repo: r}
	account, err := s.GetByID(1)
	if err != nil {
		t.Error(err)
	}
	assertEqual(t, account.Name, "Cash")
}

func TestDeleteById(t *testing.T) {
	r := &Mock{Account: &Account{ID: 1, Name: "Cash"}}
	s := &Service{repo: r}
	err := s.DeleteByID(context.Background(), 1)
	if err != nil {
		t.Errorf("error deleting account")
	}
	assertEqual(t, r.Account.ID, 0)
}
