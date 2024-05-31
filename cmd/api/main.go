package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/crossevol/goadmin/internal/database"
	"github.com/crossevol/goadmin/internal/env"
	"github.com/crossevol/goadmin/internal/smtp"
	"github.com/crossevol/goadmin/internal/version"

	"github.com/lmittmann/tint"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	godotenv.Load()
	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL   string
	httpPort  int
	basicAuth struct {
		username       string
		hashedPassword string
	}
	db struct {
		dsn         string
		automigrate bool
	}
	jwt struct {
		secretKey string
	}
	notifications struct {
		email string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		from     string
	}
}

type application struct {
	config config
	ctx    *context.Context
	db     *database.DB
	logger *slog.Logger
	mailer *smtp.Mailer
	wg     sync.WaitGroup
	q      *mysqlDao.Queries
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = env.GetInt("HTTP_PORT", 4444)
	cfg.basicAuth.username = env.GetString("BASIC_AUTH_USERNAME", "admin")
	cfg.basicAuth.hashedPassword = env.GetString("BASIC_AUTH_HASHED_PASSWORD", "$2a$10$jRb2qniNcoCyQM23T59RfeEQUbgdAXfR6S0scynmKfJa5Gj3arGJa")
	cfg.db.dsn = env.GetString("DB_DSN", "user:pass@tcp(localhost:3306)/example?parseTime=true")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)
	cfg.jwt.secretKey = env.GetString("JWT_SECRET_KEY", "m56c55pbkd4guua2t7yfmwyzqu4gco5q")
	cfg.notifications.email = env.GetString("NOTIFICATIONS_EMAIL", "")
	cfg.smtp.host = env.GetString("SMTP_HOST", "example.smtp.host")
	cfg.smtp.port = env.GetInt("SMTP_PORT", 25)
	cfg.smtp.username = env.GetString("SMTP_USERNAME", "example_username")
	cfg.smtp.password = env.GetString("SMTP_PASSWORD", "pa55word")
	cfg.smtp.from = env.GetString("SMTP_FROM", "Example Name <no_reply@example.org>")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	mailer, err := smtp.NewMailer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.from)
	if err != nil {
		return err
	}

	ctx := context.Background()
	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
		mailer: mailer,
		ctx:    &ctx,
		q:      mysqlDao.New(db),
	}

	return app.serveHTTP()
}
