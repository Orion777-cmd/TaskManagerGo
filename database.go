// database.go
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func InitializeDB() (*bun.DB, error) {
	fmt.Println("Connecting to the database...")

	// Assign to the global variable DB instead of the local variable db
	DB, err := connectToDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil, err
	}

	// Add a query hook for logging
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	// Ping the database to test the connection
	err = DB.Ping()
	if err != nil {
		fmt.Println("Failed to ping the database:", err)
		return nil, err
	}

	// Connection successful
	fmt.Println("Connected to the database")

	// Create the "tasks" table in the database if it doesn't exist
	err = createTasksTable()
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return nil, err
	}

	return DB, nil
}

func createTasksTable() error {
	ctx := context.Background()
	_, err := DB.NewCreateTable().Model(&Task{}).IfNotExists().Exec(ctx)
	return err
}

func connectToDatabase() (*bun.DB, error) {
	dsn := "postgres://mhhkwskw:NzQKOrx8ELb_7SiI3wxdJP6vpGezA5_t@horton.db.elephantsql.com/mhhkwskw"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(sqldb, pgdialect.New())
	return DB, nil
}
