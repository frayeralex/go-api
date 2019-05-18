package models

import (
	"errors"
	"fmt"
	"github.com/go-bongo/bongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Username           string `bson:"username" json:"username"`
	Password           string `bson:"password" json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hash)
	return err
}

func (u *User) Validate(*bongo.Collection) []error {
	err := make([]error, 0)
	if len(u.Username) < 2 {
		err = append(err, errors.New("username should contain more then 2 characters"))
	}
	if len(u.Password) < 6 {
		err = append(err, errors.New("password should contain more then 6 characters"))
	}
	fmt.Println("Validate", len(u.Username), err)
	return err
}

func (u *User) BeforeSave(*bongo.Collection) error {
	if u.IsNew() {
		err := u.hashPassword()
		if err != nil {
			return err
		}
	}
	return nil
}
