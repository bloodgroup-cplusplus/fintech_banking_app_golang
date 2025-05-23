package migrations

import (
	"github.com/bloodgroup-cplusplus/fintech_banking_app/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}


type Account struct {
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}


func connectDB() *grom.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")

	helpers.HanldeErr(err)
	return db
}

func createAccounts() {
	db := connectDB()
	user := [2] User {
	{Username: "Martin", Email:"martin@martin.com"}
	{Username:"Chad", Email:"chad@chad.com"},
	}

	for i :=0 ; i<len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, users[i].Email, Password:genratedPassword}
		db.Create(&user)
account := Account{Type:"Daily Account", Name:string(users[i].Username+"'s"+ "account"), Balance:uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

// functions to migrate our tables 
func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Close()
	createAccounts()
}

	




