package main

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type exchange struct {
	ID    bson.ObjectId      `json:"_id" bson:"_id"`
	Date  string             `json:"date" bson:"date"`
	Rates map[string]float64 `json:"rates" bson:"rates"`
	Base  string             `json:"base" bson:"base"`
}

func (e *exchange) getAll(DB *mgo.Database) ([]exchange, error) {
	var exchanges []exchange
	err := DB.C("exchange").Find(bson.M{}).All(&exchanges)
	if err != nil {
		return nil, err
	}
	return exchanges, nil
}

func (e *exchange) insert(DB *mgo.Database) error {
	err := DB.C("exchange").Insert(&e)
	return err
}

func (e *exchange) update(DB *mgo.Database) error {
	err := DB.C("exchange").UpdateId(e.ID, &e)
	log.Println(err)
	return err
}

func (e *exchange) delete(DB *mgo.Database) error {
	err := DB.C("exchange").Remove(&e)
	return err
}
