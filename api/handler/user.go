package handler

import (
	"api/model"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *GameApi) PostUsers(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	if err := service.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &newUser)
}

func (*GameApi) GetUsersId(c *gin.Context, id string) {
	foundUser, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &foundUser)
}

func (a *GameApi) PatchUsersId(c *gin.Context, id string) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	err := service.UpdateUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func (*GameApi) DeleteUsersId(c *gin.Context, id string) {
	err := service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "User successfully deleted"})
}

func (a *GameApi) PostUsersIdChangePassword(c *gin.Context, id string) {
	service.ChangeUserPassword(id)
	c.JSON(http.StatusOK, model.Message{Message: "email sent"})
}

func (a *GameApi) PostUsersIdResetPassword(c *gin.Context, id string) {
	service.ChangeUserPassword(id)
}
