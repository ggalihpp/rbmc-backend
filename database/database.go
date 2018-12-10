package database

import (
	"os"

	"github.com/globalsign/mgo"
)

// Connection - Return a client connection of MonggoDB
func Connection(className string) (session *mgo.Session, collection *mgo.Collection, err error) {
	session, err = mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		return
	}

	db := session.DB("rbmc")

	collection = db.C(className)

	return
}

// Disconnect - Will disconnect a client connection
func Disconnect(db *mgo.Session) {
	db.Close()
	return
}
