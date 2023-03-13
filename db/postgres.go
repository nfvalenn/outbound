package db

import (
	//"fmt"
	"os"
	"warehouse/model"

	//"gorm.io/driver/postgres"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connectdb() error {
	//dsn := fmt.Sprintln("host=localhost user=postgres password=123 dbname=postgres port=5432 ssmlmode=disable TimeZone=Asia/Jakarta")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN: os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(model.Outbound{})
	SetupDBConnection(db)

	return nil
}

func SetupDBConnection(DB *gorm.DB)  {
	db = DB
}

func GetDBConnection() *gorm.DB  {
	return db
}