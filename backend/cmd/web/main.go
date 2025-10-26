package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"playwhot.io/pkg/models/postgres"
)

type application struct {
	addr        string
	infoLog     *log.Logger
	errorLog    *log.Logger
	environment string
	version     string
	users       interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, string, error)
	}
	rooms interface {
		Create(int) (int, error)
	}
}

var sessionManager *scs.SessionManager
var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	applicationPort := flag.String("port", ":4000", "API server port")
	migrationsDir := os.Getenv("GOOSE_MIGRATION_DIR")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	flag.Parse()

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false // Set to true in production with HTTP

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		errorLog.Println("DB_DSN not set; skipping migrations")
	} else {
		db, err = sql.Open("postgres", dsn)

		if err != nil {
			errorLog.Fatalf("failed to open db for migrations: %v", err)
		}

		db.SetMaxOpenConns(50)
		db.SetMaxIdleConns(50)
		db.SetConnMaxLifetime(10 * time.Minute)
		db.SetConnMaxIdleTime(5 * time.Minute)
		defer db.Close()

		migPath, err := filepath.Abs(migrationsDir)
		if err != nil {
			errorLog.Fatalf("failed to resolve migrations dir: %v", err)
		}

		dialect := os.Getenv("GOOSE_DRIVER")
		goose.SetDialect(dialect)
		if err := goose.Up(db, migPath); err != nil {
			errorLog.Fatalf("goose up failed: %v", err)
		}
		infoLog.Printf("migrations applied from %s", migPath)
	}

	app := &application{
		addr:        *applicationPort,
		infoLog:     infoLog,
		errorLog:    errorLog,
		environment: os.Getenv("ENVIRONMENT"),
		version:     os.Getenv("API_VERSION"),
		users:       &postgres.UserModel{DB: db},
		rooms:       &postgres.RoomModel{DB: db},
	}

	infoLog.Printf("Listening to port %s", app.addr)
	http.ListenAndServe(app.addr, sessionManager.LoadAndSave(app.routes()))
}
