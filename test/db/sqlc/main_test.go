package sqlc

import (
	"context"
	db "github.com/daniel-vuky/gogento-auth/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testStore db.Store

// TestMain
// Init store and database connection
func TestMain(m *testing.M) {
	connectionString := "postgresql://root:secret@localhost:5432/gogento?sslmode=disable"
	connPool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	testStore = db.NewStore(connPool)

	os.Exit(m.Run())
}
