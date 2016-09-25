package db

import mgo "gopkg.in/mgo.v2"

//Session is our database session
var Session *mgo.Session

//Dial dials for dialing mongo server
func Dial() {
	var err error
	Session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	Session.SetMode(mgo.Monotonic, true)
}
