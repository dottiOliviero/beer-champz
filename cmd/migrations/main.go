package main

import (
	"beerchampz/pkg/config"
	"database/sql"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var direction string

func init() {
	flag.StringVar(&direction, "dir", "up", "migration direction")
	flag.Parse()
	if direction != "up" && direction != "down" {
		log.Fatalf("unknown direction received: %q. Possible values are: up, down", direction)
	}
}

func main() {
	conf := config.Get()
	connStr := conf.GetDBConnStr()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening connection : %v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("error with driver %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("error NewWithDatabaseInstance %v", err)
	}

	err = exec(m, direction)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("after exec %v", err)
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}

func exec(m *migrate.Migrate, dir string) error {
	if dir == "up" {
		return m.Up()
	}

	return m.Down()
}
