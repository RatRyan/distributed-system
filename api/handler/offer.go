package handler

import (
	"api/model"
	"api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *GameApi) PostOffers(c *gin.Context) {
	var offer model.Offer
	if err := c.BindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	if err := service.CreateOffer(&offer); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &offer)
}

func (a *GameApi) GetOffersId(c *gin.Context, id string) {
	foundOffer, err := service.GetOfferById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &foundOffer)
}

func (a *GameApi) PatchOffersId(c *gin.Context, id string) {
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

func (a *GameApi) DeleteOffersId(c *gin.Context, id string) {
	err := service.DeleteOffer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "Offer successfully deleted"})
}

func (*GameApi) PostOffersIdAccept(c *gin.Context, id string) {
	err := service.AcceptOffer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "Successfully accepted offer"})
}

func (*GameApi) PostOffersIdReject(c *gin.Context, id string) {
	err := service.RejectOffer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Message{Message: "Successfully rejected offer"})
}

func (*GameApi) GetOffersIdIncoming(c *gin.Context, id string) {
	foundOffers, err := service.GetOffersByStatus(id, c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, foundOffers)
}

func (*GameApi) GetOffersIdOutgoing(c *gin.Context, id string) {
	foundOffers, err := service.GetOffersByStatus(id, c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, foundOffers)
}
