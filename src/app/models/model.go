package models

type VisitAPI struct {
	Visits int `json:"visits"`
}

type MessageAPI struct {
	Message string `json:"message"`
}

type ErrorAPI struct {
	Error string `json:"error"`
}
