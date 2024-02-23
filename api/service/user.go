package service

import (
	"api/db"
	"api/model"
	"api/mq"
	"errors"
)

func CreateUser(newUser *model.User) error {
	c := db.GetUserColl()

	if newUser.Email == nil {
		return errors.New("user email cannot be nil")
	}
	if newUser.Password == nil {
		return errors.New("user password cannot be nil")
	}

	generatedId, err := insertObj[model.User](newUser, c)
	if err != nil {
		return err
	}

	newUser.Id = generatedId

	return nil
}

func GetUserById(id string) (*model.User, error) {
	c := db.GetUserColl()

	foundUser, err := getObjById[model.User](id, c)
	if err != nil {
		return nil, err
	}

	foundUser.Password = nil

	return foundUser, nil
}

func UpdateUser(updatedUser *model.User, id string) error {
	c := db.GetUserColl()

	updatedUser.Id = nil
	updatedUser.Email = nil
	updatedUser.Password = nil

	err := updateObj[model.User](updatedUser, id, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	c := db.GetUserColl()

	err := deleteObj[model.User](id, c)
	if err != nil {
		return err
	}

	return nil
}

func ChangeUserPassword(id string) error {
	c := db.GetUserColl()

	foundUser, err := getObjById[model.User](id, c)
	if err != nil {
		return err
	}

	mq.ProducePasswordChange(foundUser)

	return nil
}
