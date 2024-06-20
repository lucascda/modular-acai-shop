package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func NewPostgresDB(url string) (*postgres, error) {
	var err error
	var pg = &postgres{}

	log.Printf("Connecting to %s", url)
	pg.db, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Error while trying to connect to %s", url)
	}
	log.Print("Connected to Postgresql")
	return pg, nil
}

func (p *postgres) GetDB() (*pgxpool.Conn, error) {
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		log.Fatal("Cant get connection from connection pool")
	}
	return conn, nil
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
