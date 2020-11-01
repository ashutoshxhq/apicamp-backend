package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// DatabasePool is global database pool
	DatabasePool *pgxpool.Pool
)

//InitializeDatabase ...
func InitializeDatabase() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	DatabasePool = dbpool
}
