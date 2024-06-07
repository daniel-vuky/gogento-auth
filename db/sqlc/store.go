package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store interface provide method of all support queries
type Store interface {
	Querier
	// If we had tx functions, we would add them here
}

// SQLStore struct provide the connection to the database and have all the implemented queries
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore create a new store instance that implement Store interface
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(connPool),
		connPool: connPool,
	}
}
