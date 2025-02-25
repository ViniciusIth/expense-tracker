package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(connString string) (*pgxpool.Pool, error) {
	var err error

	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}

func Close(db *pgxpool.Pool) {
	if db != nil {
		db.Close()
		log.Println("Database connection pool closed.")
	}
}
