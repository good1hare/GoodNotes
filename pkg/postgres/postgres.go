// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"os"
)

// New -.
func New(url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	return conn, err
}
