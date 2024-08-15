package db

import (
	"fmt"
	"log"

	"gihtub.com/heyyakash/switchr/internal/modals"
	"gihtub.com/heyyakash/switchr/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Host     = utils.GetString("POSTGRES_HOST")
	Database = utils.GetString("POSTGRES_DB")
	User     = utils.GetString("POSTGRES_USER")
	Password = utils.GetString("POSTGRES_PASSWORD")
	Port     = utils.GetString("POSTGRES_PORT")
)

type PostgresStore struct {
	DB *gorm.DB
}

var Store PostgresStore

func (p *PostgresStore) CreateTable() {
	if err := p.DB.AutoMigrate(&modals.Users{}); err != nil {
		log.Fatal("Couldn't Migrate : ", err)
	}
}

func Init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", Host, User, Password, Port)
	// start connection to db
	Store.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to DB : ", err)
	}
	dbName := Database

	// create db
	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	err = Store.DB.Exec(createDBQuery).Error
	if err != nil && err.Error() != fmt.Sprintf("ERROR: database \"%s\" already exists (SQLSTATE 42P04)", dbName) {
		log.Fatalf("Error creating database: %v", err.Error())
	}
	//add extension for uuid
	Query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	err = Store.DB.Exec(Query).Error
	if err != nil {
		log.Fatalf("Error: ", err)
	}

	log.Print("DB Connected")

	// create table if not exists
	Store.CreateTable()
}
