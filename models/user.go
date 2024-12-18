package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int
	Login    string
	Password string
}

func (u *User) TableName() string {
	return "Users"
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)

	return nil
}
