package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/entity"

	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// this file should be a connect Database only

// DB is struct of this file.
type DB struct {
	Connection *gorm.DB
	sql        *sql.DB
	env        config.Configuration
}

func (db *DB) IsErrorRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// Close Connection DB
func (db *DB) Close() {
	if err := db.sql.Close(); err != nil {
		fmt.Printf("Error closing db connection %s", err)
	} else {
		fmt.Println("DB connection closed")
	}
}

func (db *DB) MigrateDB() {
	fmt.Println("Start migrate db READ")

	if !db.Connection.Migrator().HasTable(entity.Product{}.TableName()) {
		err := db.Connection.AutoMigrate(&entity.Product{})

		log.Println("Error :", err)
	}
	if !db.Connection.Migrator().HasTable(entity.Category{}.TableName()) {
		err := db.Connection.AutoMigrate(&entity.Category{})

		log.Println("Error :", err)
	}
}

// NewServerBase is start connection database.
func NewServerBase(env config.Configuration) *DB {
	log.Println("start NewserverBase")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.DbHost,
		env.DbPort,
		env.DbUser,
		env.DbName,
		env.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if env.Env != "production" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(7)
	sqlDB.SetMaxIdleConns(5)

	return &DB{
		Connection: db,
		sql:        sqlDB,
		env:        env,
	}
}
