package models

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/LordRahl90/little_quiz_backend/src/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//Token struct to keep the struct details.
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//Account struct to manage User Account.
type Account struct {
	gorm.Model
	Fullname string   `json:"fullname"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Password string   `json:"password"`
	Token    string   `json:"token" sql:"-"`
	UserRole UserRole `gorm:"foreignkey:UserID" json:"user_role"`
}

//Validate Account before moving to the next operation
func (account Account) Validate() (map[string]interface{}, bool) {
	if account.Email == "" {
		return utils.Message(false, "Email is required"), false
	}

	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Invalid Email Format Detected"), false
	}

	if account.Fullname == "" {
		return utils.Message(false, "Fullname needs to be provided please"), false
	}

	if account.Phone == "" {
		return utils.Message(false, "Phone Number needs to be provided please"), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password must be greater than 5 characters"), false
	}

	db := GetDB()
	temp := &Account{}
	err := db.Table("accounts").Where("email=?", account.Email).First(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection Error, Please try again."), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email has been taken already."), false
	}

	return utils.Message(false, "Validation Successful"), true
}

//Create - function to create an account. This creates an account and generates a token
func (account *Account) Create() (responseData map[string]interface{}) {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Message(false, "Cannot Gneerate password hash at this time, please try again.")
	}
	account.Password = string(hashedPassword)
	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, Connection error.")
	}

	//proceed to generate token.
	tk := Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	response := utils.Message(true, "Account Created Successfully.")
	response["account"] = account
	return response
}

//Login - Function to authenticate a user account.
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email=?", email).First(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email Does not exist")
		}
		return utils.Message(false, "Connection error, Please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return utils.Message(false, "Invalid Login Credentials, Please try again")
	}

	account.Password = ""
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	response := utils.Message(true, "Login Successful")
	response["account"] = account
	return response
}

//GetUser - Retrieve the user details
func GetUser(u uint) *Account {
	account := &Account{}
	GetDB().Table("accounts").Where("id=?", u).First(&account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}
