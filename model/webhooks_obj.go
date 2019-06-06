package model

// WebhooksObj is creatd by elasticmail webhooks
type WebhooksObj struct {
	Transaction string `schema:"transaction"`
	To          string `schema:"to"`
	Date        string `schema:"date"`
	Status      string `schema:"status"`
	Channel     string `schema:"channel"`
}
