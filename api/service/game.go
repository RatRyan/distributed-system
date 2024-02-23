package service

import (
	"api/db"
	"api/model"
)

func CreateGame(newGame *model.Game) error {
	c := db.GetGameColl()
	generatedId, err := insertObj[model.Game](newGame, c)
	if err != nil {
		return err
	}

	newGame.Id = generatedId
	return nil
}

func GetGameById(id string) (*model.Game, error) {
	c := db.GetGameColl()
	foundGame, err := getObjById[model.Game](id, c)
	if err != nil {
		return nil, err
	}

	return foundGame, nil
}

func UpdateGame(updatedGame *model.Game, id string) error {
	c := db.GetGameColl()

	updatedGame.Id = nil

	err := updateObj[model.Game](updatedGame, id, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGame(id string) error {
	c := db.GetGameColl()

	err := deleteObj[model.Game](id, c)
	if err != nil {
		return err
	}

	return nil
}
