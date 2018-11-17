package main

import (
	"log"

	"github.com/globalsign/mgo"
)

type connector struct {
	Server   string
	Database string
}

func (c *connector) connect() *mgo.Database {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(c.Database)
}
