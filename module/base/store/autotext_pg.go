package store

import (
	"context"
	"database/sql"
)

// SQLStore provides all functions to execute db queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Repository {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (s *SQLStore) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *SQLStore) RollbackTx(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) CommitTx(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// execTX executes a function within a database transaction
// func (s *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
