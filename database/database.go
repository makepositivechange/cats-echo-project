package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setting connection poll const
const POOL = 25

func MySQLConn(ctx context.Context, username string, password string, host string, port string, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error while connecting to database:%v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error while getting database object:%v", err)
		return nil, err
	}
	// Setting connection pool setting
	sqlDB.SetMaxOpenConns(POOL)
	sqlDB.SetMaxIdleConns(POOL)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Pinging to verify connection timeout
	pingCh := make(chan error, 1)
	go func() {
		pingCh <- sqlDB.PingContext(ctx)
	}()
	select {
	case <-ctx.Done():
		log.Printf("Database ping timed out:%v", ctx.Err())
		return nil, ctx.Err()
	case err := <-pingCh:
		if err != nil {
			log.Printf("Error pinging database:%v", err)
			return nil, err
		}

	}
	return db, nil
}
