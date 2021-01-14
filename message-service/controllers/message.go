package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"stewped-applet/common/data"
	"stewped-applet/common/failures"
	md "stewped-applet/message-service/data"
	"stewped-applet/message-service/models"
)

// Saves a message. Returns 200 if successful and 500 if an unexpected error
// occurrs.
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		failures.WriteError(w, err, "Invalid message data", http.StatusInternalServerError)
		return
	}

	message.Digest = calculateSum(message.Message)

	session := data.GetMongoSession()
	defer session.Close()
	c := data.GetCollection(session, "messages")
	repo := &md.MessageRepository{c}
	repo.Create(&message)
	j, err := json.Marshal(&MessagePostResponse{Digest: message.Digest})
	if err != nil {
		failures.WriteError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Retrieves a message using its digest. Responds with 200 if a message was
// found, otherwise a 404 with a custom response body is returned.
func GetMessageByDigest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	digest := vars["hash"]

	session := data.GetMongoSession()
	defer session.Close()
	c := data.GetCollection(session, "messages")
	repo := &md.MessageRepository{c}

	message, err := repo.GetByDigest(digest)
	if err != nil {
		if err == mgo.ErrNotFound {
			// write custom response per requirements
			failures.WriteCustomError(
				w, err, http.StatusNotFound,
				map[string]string{
					"error":          "unable to find message",
					"message_sha256": digest,
				})
			return
		}

		failures.WriteError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(&MessageGetResponse{Message: message.Message})
	if err != nil {
		failures.WriteError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Deletes a message. Responds with 200 if; a) message was
// successfully deleted, or b) the message wasn't found, otherwise 500.
func DeleteMessageByDigest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	digest := vars["hash"]

	session := data.GetMongoSession()
	defer session.Close()
	c := data.GetCollection(session, "messages")
	repo := &md.MessageRepository{c}

	err := repo.Delete(digest)
	if err != nil {
		if err != mgo.ErrNotFound {
			failures.WriteError(w, err, "An unexpected error has occurred", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func calculateSum(message string) string {
	sum := sha256.Sum256([]byte(message))
	return hex.EncodeToString(sum[:])
}
