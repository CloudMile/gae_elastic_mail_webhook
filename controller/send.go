package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/cloudmile/gae_elastic_mail_webhook/model"
	"google.golang.org/appengine/log"
)

// Send is the an endpoint "POST /send"
func Send(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "parse err")
		return
	}
	ctx := r.Context()
	log.Infof(ctx, "POST /send")
	ct := r.Header.Get("Content-Type")
	log.Infof(ctx, "Content-Type is: %s", ct)

	contentType := checkContentType(ctx, ct)
	form, err := makeMailParams(r, contentType)

	if err != nil {
		log.Errorf(ctx, "err %v", err)
	}
	log.Infof(ctx, "url values is %+v", form)

	gaeMail := makeGaeMail(ctx, &form)
	gaeMail.Send()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", `{"result": "sent success"}`)
}

func makeMailParams(r *http.Request, contentType string) (form model.Form, err error) {
	switch contentType {
	case "multipart/form-data":
		form = model.Form{
			FromEmail:  r.FormValue("from_email"),
			FromName:   r.FormValue("from_name"),
			EnvFrom:    r.FormValue("env_from"),
			EnvToList:  r.FormValue("env_to_list"),
			ToList:     r.FormValue("to_list"),
			HeaderList: r.FormValue("header_list"),
			Subject:    r.FormValue("subject"),
			BodyText:   r.FormValue("body_text"),
			BodyHTML:   r.FormValue("body_html"),
		}
	case "application/x-www-form-urlencoded":
		form = model.Form{
			FromEmail:  strings.Join(r.Form["from_email"], ","),
			FromName:   strings.Join(r.Form["from_name"], ","),
			EnvFrom:    strings.Join(r.Form["env_from"], ","),
			EnvToList:  strings.Join(r.Form["env_to_list"], ","),
			ToList:     strings.Join(r.Form["to_list"], ","),
			HeaderList: strings.Join(r.Form["header_list"], ","),
			Subject:    strings.Join(r.Form["subject"], ","),
			BodyText:   strings.Join(r.Form["body_text"], ","),
			BodyHTML:   strings.Join(r.Form["body_html"], ","),
		}
	case "application/json":
		body, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &form)
	}
	return
}

func checkContentType(ctx context.Context, contentType string) (reContentType string) {
	reContentType = strings.Split(contentType, ";")[0]
	log.Infof(ctx, "split contentType is: %s", reContentType)
	return
}

func makeGaeMail(ctx context.Context, form *model.Form) (gaeMail model.GaeMail) {
	gaeMail = model.GaeMail{
		Ctx:      ctx,
		To:       os.Getenv("SEND_TO"),
		Subject:  os.Getenv("SEND_SUBJECT") + form.Subject,
		Body:     "From: " + form.FromEmail + "\nTo: " + form.ToList + "\n" + form.BodyText,
		HTMLBody: "<p>From: " + form.FromEmail + "</p><p>To: " + form.ToList + "</p>" + form.BodyHTML,
	}

	return
}
