package data

import (
	"log"
	"time"

	"stewped-applet/common"

	"gopkg.in/mgo.v2"
)

func GetMongoSession() *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{common.GetEnv("MONGO_HOST", "mongo")},
		Username: common.GetEnv("MONGO_USERNAME", ""),
		Password: common.GetEnv("MONGO_PASSWORD", ""),
		Database: common.GetEnv("MONGO_DATABASE", "messages"),
		Timeout:  30 * time.Second,
	})
	if err != nil {
		log.Fatalf("failure in GetMongoSession: %s\n", err)
	}
	return session
}

func GetCollection(session *mgo.Session, name string) *mgo.Collection {
	return session.DB(common.GetEnv("MONGO_DATABASE", "")).C(name)
}
