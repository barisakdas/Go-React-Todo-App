package entity

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Dsn string = "host=db user=admin password=Acaba1234!** dbname=todo port=5432 sslmode=disable"

func DbMigration() {
	Todo{}.Migrate()
}

func CreateDb() {
	db, _ := gorm.Open("postgres", "host=db port=5432 user=admin dbname=postgres password=Acaba1234!** sslmode=disable")
	db = db.Exec("CREATE DATABASE todo;")
	if db.Error != nil {
		fmt.Println("Unable to create DB todo, attempting to connect assuming it exists...")
	}
}
