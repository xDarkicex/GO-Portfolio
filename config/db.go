package config

import mgo "gopkg.in/mgo.v2"

//Session is our database session
var Session *mgo.Session

//Dial dials shit
func Dial() *mgo.Session {

	Session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	return Session
}
