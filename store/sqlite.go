package store

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteStore struct {
}

func NewSQLiteStore() *SQLiteStore {
	return &SQLiteStore{}
}

// Connect connects to the SQLite database.
func connect(ctx context.Context) (*gorm.DB, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("sql store startup deadline exceeded")
		default:
			dialector := sqlite.Open("test.db")

			// Connect to the SQLite database.
			db, err := gorm.Open(dialector, &gorm.Config{})
			if err != nil {
				// Handle the error.
				continue
			}

			sqlDB, err := db.DB()
			if err != nil {
				return nil, err
			}

			// Ping the database to check if the connection is alive.
			err = sqlDB.Ping()

			if err != nil {
				// Handle the error.
				continue
			}

			return db, nil
		}
	}
}
