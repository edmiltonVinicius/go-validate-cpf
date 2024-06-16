package persistence

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	pgMigrations "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/edmiltonVinicius/go-validate-cpf/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	path_migrations = "file://internal/adapters/persistence/sql/migrations"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Global.DB_HOST, config.Global.DB_PORT, config.Global.DB_USER, config.Global.DB_PASSWORD, config.Global.DB_NAME, config.Global.SSL_MODE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if config.Global.RUN_MIGRATIONS {
		runMigrations(db)
	}

	config.Global.DB = db

	return db, nil
}

func runMigrations(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pgMigrations.WithInstance(conn, &pgMigrations.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		path_migrations,
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
