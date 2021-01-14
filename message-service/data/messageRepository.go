package data

import (
	"time"

	"stewped-applet/message-service/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IMessageRepository interface {
	Create(*models.Message) error
	GetByDigest(digest string) (models.Message, error)
	Delete(digest string) error
}

type MessageRepository struct {
	C *mgo.Collection
}

func (r *MessageRepository) Create(message *models.Message) error {
	obj_id := bson.NewObjectId()
	message.Id = obj_id
	message.CreatedAt = time.Now()
	err := r.C.Insert(&message)
	return err
}

func (r *MessageRepository) GetByDigest(digest string) (message models.Message, err error) {
	err = r.C.Find(bson.M{"digest": digest}).One(&message)
	return
}

func (r *MessageRepository) Delete(digest string) error {
	err := r.C.Remove(bson.M{"digest": digest})
	return err
}
