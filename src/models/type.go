package models

type findAllUsers struct {
	Users []Users `json:"users"`
	Count int64   `json:"count"`
}

type findAllWebhook struct {
	Webhooks []Webhooks `json:"users"`
	Count    int64      `json:"count"`
}
