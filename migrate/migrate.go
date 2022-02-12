package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/instill-ai/model-backend/configs"
)

func checkExist(databaseConfig configs.DatabaseConfig) error {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable TimeZone=%s",
			databaseConfig.Host,
			databaseConfig.Username,
			databaseConfig.Password,
			databaseConfig.Port,
			databaseConfig.TimeZone,
		),
	)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Open() may just validate its arguments without creating a connection to the database.
	// To verify that the data source name is valid, call Ping().
	if err = db.Ping(); err != nil {
		panic(err)
	}

	var rows *sql.Rows
	rows, err = db.Query(fmt.Sprintf("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('%s');", databaseConfig.Name))

	if err != nil {
		panic(err)
	}

	dbExist := false
	defer rows.Close()
	for rows.Next() {
		var database_name string
		if err := rows.Scan(&database_name); err != nil {
			panic(err)
		}

		if databaseConfig.Name == database_name {
			dbExist = true
			fmt.Printf("Database %s exist\n", database_name)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	if !dbExist {
		fmt.Printf("Create database %s\n", databaseConfig.Name)
		if _, err := db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\";", databaseConfig.Name)); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	mydir, _ := os.Getwd()

	_ = configs.Init()
	databaseConfig := configs.Config.Database
	if err := checkExist(databaseConfig); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name,
		"sslmode=disable",
	)

	m, err := migrate.New(fmt.Sprintf("file:///%s/internal/db/migrations", mydir), dsn)

	if err != nil {
		panic(err)
	}

	version, dirty, err := m.Version()

	if err != nil && version != 0 {
		panic(err)
	}

	ExpectedVersion := databaseConfig.Version

	fmt.Printf("Expected migration version is %d\n", ExpectedVersion)
	fmt.Printf("The current schema version is %d, and dirty flag is %t\n", version, dirty)
	if dirty {
		panic("The database has dirty flag, please fix it")
	}

	step := version
	for {
		if ExpectedVersion <= step {
			fmt.Printf("Migration to version %d complete\n", step)
			break
		} else {
			fmt.Printf("Step up to version %d\n", step+1)
			if err := m.Up(); err != nil {
				panic(err)
			}
		}

		step, _, err = m.Version()

		if err != nil {
			panic(err)
		}
	}
}
