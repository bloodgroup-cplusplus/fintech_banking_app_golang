package helpers


import (
	"golang.org/x/crypto/bcrypt"
)


func HandleErr(err rror) {

	if err !=nil {
		panic(err.Error())
	}
}


func HashAndSalt (pass [] byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass,bcrypt.MinCost)

	HnaldeErr(err)
	return string(hashed)
}

