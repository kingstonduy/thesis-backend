package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func main() {
	// PostgreSQL connection details
	user := "postgres"      // Replace with your PostgreSQL username
	password := "postgres"  // Replace with your PostgreSQL password
	host := "10.148.0.7"    // Replace with your PostgreSQL host
	port := 5432            // Replace with your PostgreSQL port
	defaultDB := "postgres" // Default database to connect to

	// List of databases to create
	dbNames := []string{"cart", "product", "userr", "orderr", "payment"}

	// Prompt the user for an action
	fmt.Println("Choose an option:")
	fmt.Println("1. Only create databases")
	fmt.Println("2. Only execute queries in SQL files")
	fmt.Print("Enter your choice (1 or 2): ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		createDatabases(user, password, host, port, defaultDB, dbNames)
	} else if choice == 2 {
		executeSQLFiles(user, password, host, port, dbNames)
	} else {
		fmt.Println("Invalid choice. Exiting.")
	}
}

func createDatabases(user, password, host string, port int, defaultDB string, dbNames []string) {
	// Construct the connection string for the default database
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, defaultDB)

	// Connect to PostgreSQL server
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// pre set up
	sqlQuery := `
        ALTER SYSTEM SET max_replication_slots = 10;
        ALTER SYSTEM SET max_wal_senders = 10;
        SELECT pg_reload_conf();
        SET TIME ZONE 'GMT+7';
    `
	_, err = conn.Exec(context.Background(), sqlQuery)
	if err != nil {
		log.Printf("Failed to execute pre set up query: %v\n", err)
	}

	// Loop through the list of databases to check and create
	for _, dbName := range dbNames {
		var exists bool
		queryCheck := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
		err = conn.QueryRow(context.Background(), queryCheck).Scan(&exists)
		if err != nil {
			log.Printf("Failed to check database existence for '%s': %v\n", dbName, err)
			continue
		}

		if exists {
			fmt.Printf("Database '%s' already exists.\n", dbName)
		} else {
			queryCreate := fmt.Sprintf("CREATE DATABASE %s", dbName)
			_, err = conn.Exec(context.Background(), queryCreate)
			if err != nil {
				log.Printf("Failed to create database '%s': %v\n", dbName, err)
				continue
			}
			fmt.Printf("Database '%s' created successfully.\n", dbName)
		}
	}
}

func executeSQLFiles(user, password, host string, port int, dbNames []string) {
	for _, dbName := range dbNames {
		sqlFileName := fmt.Sprintf("%s.sql", dbName)
		if _, err := os.Stat(sqlFileName); os.IsNotExist(err) {
			fmt.Printf("SQL file '%s' does not exist. Skipping execution for database '%s'.\n", sqlFileName, dbName)
			continue
		}

		if err := executeSQLFile(user, password, host, port, dbName, sqlFileName); err != nil {
			log.Printf("Failed to execute SQL file '%s': %v\n", sqlFileName, err)
			continue
		}
		fmt.Printf("SQL file '%s' executed successfully for database '%s'.\n", sqlFileName, dbName)
	}
}

func executeSQLFile(user, password, host string, port int, dbName, sqlFileName string) error {
	// Read the SQL file
	sqlContent, err := os.ReadFile(sqlFileName)
	if err != nil {
		return fmt.Errorf("unable to read SQL file: %w", err)
	}

	// Construct the connection string for the specific database
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbName)

	// Connect to the specific database
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("unable to connect to database '%s': %w", dbName, err)
	}
	defer conn.Close(context.Background())

	// Execute the SQL commands
	sqlCommands := strings.Split(string(sqlContent), ";") // Split commands by semicolon
	for _, cmd := range sqlCommands {
		cmd = strings.TrimSpace(cmd) // Remove any extra whitespace
		if cmd == "" {
			continue // Skip empty commands
		}
		_, err := conn.Exec(context.Background(), cmd)
		if err != nil {
			return fmt.Errorf("failed to execute SQL command '%s': %w", cmd, err)
		}
	}

	return nil
}
