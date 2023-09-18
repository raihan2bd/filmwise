package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raihan2bd/filmwise/models"
)

const version = "1.0.0"

// Application config
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config

	// get Environment variables
	_ = godotenv.Load()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	dsn := os.Getenv("DATABASE_URI")
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "jwt-secret"
	}

	// initialize config
	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Port should be a number")
	}
	cfg.port = portNum
	cfg.env = env
	cfg.db.dsn = dsn
	cfg.jwt.secret = jwtSecret

	// setup logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// connect with database
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLD_URI"))

	if err != nil {
		logger.Fatalf("failed to intialize Cloudinary, %v", err)
	}

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db, cld),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
