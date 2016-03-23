package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	session *mgo.Session
}

func NewMongoDbConnection() *MongoDB {
	mongodb := &MongoDB{}

	dialInfo, err := mgo.ParseURL("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Error creating dial for mongodb")
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("err connection to mongodb")
	}
	mongodb.session = session

	return mongodb
}

func (m *MongoDB) Ping() string {
	err := m.session.Ping()
	if err == nil {
		return "WORKING"
	} else {
		return "NOT WORKING"
	}
}
