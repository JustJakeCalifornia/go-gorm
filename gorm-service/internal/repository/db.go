package repository

import (
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/trite8q1/go-gorm/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabase() *gorm.DB {
	dsn := os.Getenv("DSN")
	counts := 0

	for {
		connection, err := openDB(dsn)
		connection.AutoMigrate(&entity.User{})
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}

// TODO: conncet to datbase with gorm
func openDB(dsn string) (*gorm.DB, error) {
	// db, err := sql.Open("pgx", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	return db, nil
}
