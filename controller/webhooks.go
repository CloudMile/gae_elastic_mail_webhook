package controller

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cloudmile/gae_elastic_mail_webhook/model"
	"google.golang.org/appengine/log"
)

// Webhooks is the an endpoint "GET /webhooks"
func Webhooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Infof(ctx, "GET /webhooks")
	var webhooksObj model.WebhooksObj
	makeWebhooksObj(&webhooksObj, r.URL.Query())
	log.Infof(ctx, "webhooksObj is: %+v", webhooksObj)

	gaeMail := makeGaeMailForWebhooks(ctx, &webhooksObj)
	gaeMail.Send()
}

func makeWebhooksObj(webhooksObj *model.WebhooksObj, queryValues url.Values) {
	webhooksObj.Transaction = strings.Join(queryValues["transaction"], ", ")
	webhooksObj.To = strings.Join(queryValues["to"], ", ")
	webhooksObj.Date = strings.Join(queryValues["date"], ", ")
	webhooksObj.Status = strings.Join(queryValues["status"], ", ")
	webhooksObj.Channel = strings.Join(queryValues["channel"], ", ")
	webhooksObj.Category = strings.Join(queryValues["category"], ", ")

	decodedSubject, _ := url.QueryUnescape(strings.Join(queryValues["subject"], ", "))
	webhooksObj.Subject = decodedSubject
}

func makeGaeMailForWebhooks(ctx context.Context, webhooksObj *model.WebhooksObj) (gaeMail model.GaeMail) {
	gaeMail = model.GaeMail{
		Ctx:     ctx,
		To:      os.Getenv("WEBHOOKS_TO"),
		Subject: os.Getenv("WEBHOOKS_SUBJECT") + webhooksObj.Status + "]",
		Body:    webhooksObj.EmailBody(),
	}

	return
}
