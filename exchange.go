package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type exchange struct {
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
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
	err := DB.C("").Insert(&e)
	return err
}

func (e *exchange) update(DB *mgo.Database) error {
	err := DB.C("").Update(e.Date, &e)
	return err
}

func (e *exchange) delete(DB *mgo.Database) error {
	err := DB.C("").Remove(&e)
	return err
}
