package controllers

type (
	// GET /messages/<hash>
	MessageGetResponse struct {
		Message string `json:"message"`
	}
	// POST /messages/
	MessagePostResponse struct {
		Digest string `json:"digest"`
	}
)
