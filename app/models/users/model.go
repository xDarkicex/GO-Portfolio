package users

//Address Struct
import (
	"github.com/xDarkicex/PortfolioGo/config"
	"golang.org/x/crypto/bcrypt"
)

//Address of user
type Address struct {
	Zip    string
	State  string
	City   string
	Street string
}

//User Struct
type User struct {
	Name     string
	Email    string
	Password string
	Address  Address
}

// Create create a new user in the database
func Create(email string, name string, password string) bool {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// http.Error(res, err.Error(), 500)
		return false
	}
	user := User{
		Email:    email,
		Name:     name,
		Password: string(hashedPass),
	}
	c := config.Session.DB(config.ENV).C("user")
	// Insert Datas
	err = c.Insert(user)

	if err != nil {
		panic(err)
	}

	return true
}
