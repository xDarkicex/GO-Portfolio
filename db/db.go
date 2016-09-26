package db

import mgo "gopkg.in/mgo.v2"

//Session is our database session
var _session *mgo.Session

//Dial dials for dialing mongo server
func Dial() {
	var err error
	_session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	_session.SetMode(mgo.Monotonic, true)
}

func Session() *mgo.Session {
	return _session.Clone()
}
