package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	// Loading postgres driver
	_ "github.com/lib/pq"
)

func Open() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=test sslmode=disable")

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)

	return db, nil
}
