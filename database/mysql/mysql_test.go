//+build integration

package mysql

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

var db *Database
const (
	databaseName = "integration"
)

func TestMain(m *testing.M) {

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := retry(30, 1*time.Second, ping); err != nil {
		os.Exit(1)
	}

	setup()

	v := m.Run()

	clean()

	os.Exit(v)
}

func setup() {

	database, err := connectAndMigrate()

	if err != nil {
		log.Fatal(err)
		return
	}

	db = &Database{database}
}

func clean() {
	err := db.Exec("DROP DATABASE " + databaseName).Error
	if err != nil {
		log.Fatal(err)
	}
}

func ping() error {
	parameter := fmt.Sprintf("root:password@tcp(%s:13306)/?parseTime=true", os.Getenv("DATABASE_HOST"))
	_, err := gorm.Open("mysql", parameter)
	return err
}

func connectAndMigrate() (*gorm.DB, error) {

	parameter := fmt.Sprintf("root:password@tcp(%s:13306)/%s?parseTime=true&multiStatements=true", os.Getenv("DATABASE_HOST"), databaseName)
	database, err := gorm.Open("mysql", parameter)
	if err != nil {
		return nil, err
	}

	err = Migrate("file://migrations", database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

type stop struct {
	error
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, sleep, fn)
		}
		return err
	}
	return nil
}
