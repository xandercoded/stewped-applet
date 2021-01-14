package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Message struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Message   string        `json:"message,omitempty"`
		Digest    string        `json:"digest,omitempty"`
		CreatedAt time.Time     `json:"createdat,omitempty"`
	}
)
