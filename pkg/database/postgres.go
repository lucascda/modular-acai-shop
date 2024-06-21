package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgresDB(url string) (*Postgres, error) {
	var err error
	var pg = &Postgres{}

	log.Printf("Connecting to %s", url)
	pg.db, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Error while trying to connect to %s", url)
	}
	log.Print("Connected to Postgresql")
	return pg, nil
}

func (p *Postgres) GetDB() (*pgxpool.Conn, error) {
	conn, err := p.db.Acquire(context.Background())
	if err != nil {
		log.Fatal("Cant get connection from connection pool")
	}
	return conn, nil
}

func (p *Postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
