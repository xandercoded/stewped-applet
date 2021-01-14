package common

import (
	"os"
)

type (
	configuration struct {
		MongoHost, MongoUser, MongoPassword, DatabaseName, Server string
	}
)

const (
	ContentTypeBinary = "application/octet-stream; charset=utf-8"
	ContentTypeForm   = "application/x-www-form-urlencoded; charset=utf-8"
	ContentTypeJSON   = "application/json; charset=utf-8"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
