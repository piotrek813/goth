package daos

import (
	"piotrek813/goth/db"
	"piotrek813/goth/models"
)

func FindUserByID(id int) (*models.User, error) {
	db := db.GetDB()

	var user models.User

	err := db.QueryRow("SELECT id, login, password FROM Users WHERE id = ?", id).Scan(&user.Id, &user.Login, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByLogin(login string) (*models.User, error) {
	db := db.GetDB()

	var user models.User

	err := db.QueryRow("SELECT id, login, password FROM Users WHERE login = ?", login).Scan(&user.Id, &user.Login, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SaveUser(user *models.User) error {
	db := db.GetDB()

	stmt, err := db.Prepare("INSERT INTO Users(login, password) VALUES(?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Login, user.Password)

	if err != nil {
		return err
	}

	return nil
}
