package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		// sqlx.DB: database/sql.DBをラップし、より簡単に使えるようにしたもの
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// NewStatus : Create status repository
// var _ repository.Status = (*status)(nil) で、status が repository.Status インターフェースを実装しているかをチェックしている
// なので、返り値の型を *status ではなく repository.Status になる
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// AddStatus : データベースにステータスを追加
func (s *status) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status)error{
	_, err := s.db.ExecContext(ctx, "insert into status (account_id, content, url, created_at) values (?, ?, ?, ?)",
		status.AccountID, status.Content, status.URL, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}

func (s *status) FindByID(ctx context.Context, id int) (*object.Status, error) {
	status := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(status)
	if err != nil {
		return nil, fmt.Errorf("failed to select status: %w", err)
	}

	return status, nil
}

func (s *status) FindAccountByID(ctx context.Context, id int) (*object.Account, error) {
	account := new(object.Account)
	err := s.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(account)
	if err != nil {
		return nil, fmt.Errorf("failed to select account: %w", err)
	}

	return account, nil
}