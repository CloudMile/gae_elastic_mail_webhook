package main

import (
	"net/http"

	"github.com/cloudmile/gae_elastic_mail_webhook/controller"
	"google.golang.org/appengine"
)

func main() {
	router := controller.Router()
	http.Handle("/", router)
	appengine.Main()
}
