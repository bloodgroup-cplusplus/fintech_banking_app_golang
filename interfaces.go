package interfaces


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
	UserId uint
}



type ResponseAccount struct {
	ID uint
	Name string
	Balance int
}



type ResponseUser struct {
	ID uint
	Username string
	Email string 
}
