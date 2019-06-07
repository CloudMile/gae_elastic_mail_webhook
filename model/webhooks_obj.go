package model

// WebhooksObj is creatd by elasticmail webhooks
type WebhooksObj struct {
	Transaction string `schema:"transaction"`
	To          string `schema:"to"`
	Date        string `schema:"date"`
	Status      string `schema:"status"`
	Channel     string `schema:"channel"`
	Category    string `schema:"category"`
	Subject     string `schema:"subject"`
}

// EmailBody return email body
func (webhooksObj *WebhooksObj) EmailBody() (body string) {
	body = "emails: `" + webhooksObj.To + "` have been " + webhooksObj.Status + " at " + webhooksObj.Date + "\n"
	body = body + "subject: " + webhooksObj.Subject + "\n"
	body = body + "category: " + webhooksObj.Category + "\n"
	body = body + "click here https://elasticemail.com/account/#/activity/emails to get more info"
	return
}
