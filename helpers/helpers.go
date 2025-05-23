package helpers


import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/junzhu/gorm/dialects/postgres"	
)


func HandleErr(err error) {

	if err !=nil {
		panic(err.Error())
	}
}


func HashAndSalt (pass [] byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass,bcrypt.MinCost)

	HandleErr(err)
	return string(hashed)
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	HanldeErr(err)
	return db
}
