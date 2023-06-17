package Database

import (
	"HospitalFinpro/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/hospitalapi"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect!")
	}

	db.AutoMigrate(
		&Models.Doctor{},
		&Models.User{},
		&Models.Patient{},
		&Models.Room{},
		&Models.Diagnose{},
		&Models.Payment{},
	)

	DB = db
}
