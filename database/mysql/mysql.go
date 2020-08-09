package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Database struct {
	*gorm.DB
}

func ConnectDB() *Database {

	database, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/integration?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
	
	err = Migrate("file://database/mysql/migrations", database)
	if err != nil {
		log.Panic(err)
	}

	return &Database{database}
}

func Migrate(filPath string, database *gorm.DB) error {

	driver, err := mysql.WithInstance(database.DB(), &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(filPath, "mysql", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}