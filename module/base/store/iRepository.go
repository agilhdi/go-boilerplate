package store

import (
	"context"
	"database/sql"
)

// Repository for
type Repository interface {

	// Tx Function
	WithTx(tx *sql.Tx) *Queries
	BeginTx(ctx context.Context) (*sql.Tx, error)
	RollbackTx(tx *sql.Tx) error
	CommitTx(tx *sql.Tx) error
}
