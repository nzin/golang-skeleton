package db

import (
	"time"

	retry "github.com/avast/retry-go"
	"github.com/nzin/golang-skeleton/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

var AutoMigrateTables = []interface{}{
	Todo{},
}

func connectDB(dbtype string, dsn string, retryAttempt uint, retryDelay time.Duration) (db *gorm.DB, err error) {
	logger := &Logger{
		LogLevel:                  gorm_logger.Warn,
		SlowThreshold:             1000 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
	}

	err = retry.Do(
		func() error {
			switch config.Config.DBDriver {
			case "postgres":
				db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
					Logger: logger,
				})
			case "sqlite3":
				db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
					Logger: logger,
				})
			case "mysql":
				db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
					Logger: logger,
				})
			}
			return err
		},
		retry.Attempts(retryAttempt),
		retry.Delay(retryDelay),
	)
	return db, err
}

func NewDB() (*gorm.DB, error) {
	db, err := connectDB(config.Config.DBDriver, config.Config.DBConnectionStr, config.Config.DBConnectionRetryAttempts, config.Config.DBConnectionRetryDelay)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(AutoMigrateTables...)

	return db, err
}

func NewTestDB() (*gorm.DB, error) {
	//db, err := connectDB("sqlite3", "test.sqlite", 1, time.Second)
	db, err := connectDB("sqlite3", ":memory:", 1, time.Second)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(AutoMigrateTables...)

	return db, err
}
