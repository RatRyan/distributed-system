package handler

import (
	"api/model"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *GameApi) PostGames(c *gin.Context) {
	var game model.Game
	if err := c.BindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	if err := service.CreateGame(&game); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &game)
}

func (a *GameApi) GetGamesId(c *gin.Context, id string) {
	foundGame, err := service.GetGameById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &foundGame)
}

func (a *GameApi) PatchGamesId(c *gin.Context, id string) {
	var game model.Game
	if err := c.BindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	err := service.UpdateGame(&game, id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &game)
}

func (a *GameApi) DeleteGamesId(c *gin.Context, id string) {
	err := service.DeleteGame(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "Game successfully deleted"})
}
