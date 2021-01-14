package routers

import (
	c "stewped-applet/message-service/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/message/{hash}", c.GetMessageByDigest).Methods("GET")
	r.HandleFunc("/messages", c.CreateMessage).Methods("POST")
	r.HandleFunc("/messages/{hash}", c.DeleteMessageByDigest).Methods("DELETE")
	return r
}
