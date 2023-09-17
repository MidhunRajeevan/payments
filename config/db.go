package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB config
var DB *pgxpool.Pool

const (
	dbHost = "DB_HOST"
	dbPort = "DB_PORT"
	dbUser = "DB_USER"
	dbPass = "DB_PASS"
	dbName = "DB_NAME"
	dbMxcx = "DB_MAX_CONNECTIONS"
)

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbHost)
	if !ok {
		panic("DB_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbPort)
	if !ok {
		panic("DB_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbUser)
	if !ok {
		panic("DB_USER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbPass)
	if !ok {
		panic("DB_PASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbName)
	if !ok {
		panic("DB_NAME environment variable required but not set")
	}
	mxcx, ok := os.LookupEnv(dbMxcx)
	if !ok {
		panic("DB_MAX_CONNECTIONS environment variable required but not set")
	}

	conf[dbHost] = host
	conf[dbPort] = port
	conf[dbUser] = user
	conf[dbPass] = password
	conf[dbName] = name
	conf[dbMxcx] = mxcx
	return conf
}

type queryLogger struct {
}

func (l *queryLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	// remove false condition to debug
	if false && level == pgx.LogLevelInfo && msg == "Query" {
		fmt.Printf("SQL:\n%s\nARGS:%v\n", data["sql"], data["args"])
	}
}

// InitializeDB PostgreSQL
func InitializeDB() {
	var err error

	dbEnv := dbConfig()
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbEnv[dbUser], dbEnv[dbPass], dbEnv[dbHost], dbEnv[dbPort], dbEnv[dbName])

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Database Configuration Error: %v\n", err)
	}
	mxcx, err := strconv.Atoi(dbEnv[dbMxcx])
	if err != nil {
		log.Fatalf("Invalid Max Connections: %v\n", err)
	}
	config.MaxConns = int32(mxcx)

	config.ConnConfig.Logger = &queryLogger{}

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Database config established.")

	// Setup database object if not exists
	Setup()
}
