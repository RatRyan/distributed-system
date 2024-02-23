package service

import (
	"api/db"
	"api/model"
	"api/mq"
	"context"
	"time"

	"github.com/oapi-codegen/runtime/types"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOffer(newOffer *model.Offer) error {
	c := db.GetOfferColl()

	newOffer.Timestamp = types.Date{Time: time.Now()}
	newOffer.Status = model.Pending

	generatedId, err := insertObj[model.Offer](newOffer, c)
	if err != nil {
		return err
	}

	newOffer.Id = generatedId

	mq.ProduceOfferCreated(newOffer)

	return nil
}

func GetOfferById(id string) (*model.Offer, error) {
	c := db.GetOfferColl()
	foundOffer, err := getObjById[model.Offer](id, c)
	if err != nil {
		return nil, err
	}

	return foundOffer, nil
}

func GetOffersByStatus(id string, status string) ([]*model.Offer, error) {
	c := db.GetOfferColl()

	filter := bson.M{"status": status}
	cursor, err := c.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var foundOffers []*model.Offer
	if err = cursor.All(context.Background(), &foundOffers); err != nil {
		return nil, err
	}

	return foundOffers, nil
}

func UpdateOffer(updatedOffer *model.Offer, id string) error {
	c := db.GetOfferColl()

	updatedOffer.Id = nil

	err := updateObj[model.Offer](updatedOffer, id, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOffer(id string) error {
	c := db.GetOfferColl()

	err := deleteObj[model.Offer](id, c)
	if err != nil {
		return err
	}

	return nil
}

func AcceptOffer(id string) error {
	c := db.GetOfferColl()

	offerToAccept, err := getObjById[model.Offer](id, c)
	if err != nil {
		return err
	}

	offerToAccept.Status = model.Accepted
	if err := updateObj[model.Offer](offerToAccept, id, c); err != nil {
		return err
	}

	mq.ProduceOfferAccepted(offerToAccept)

	return nil
}

func RejectOffer(id string) error {
	c := db.GetOfferColl()

	offerToReject, err := getObjById[model.Offer](id, c)
	if err != nil {
		return err
	}

	offerToReject.Status = model.Accepted
	if err := updateObj[model.Offer](offerToReject, id, c); err != nil {
		return err
	}

	mq.ProduceOfferRejected(offerToReject)

	return nil
}
