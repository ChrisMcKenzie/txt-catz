package main

import (
	"net/http"
	"os"

	"github.com/carlosdp/twiliogo"
)

const (
	accountSid = "ACe3f0bf659491650bbff689790fd574e1"
	authToken  = "c57d2b5b368d0dd291c3c5c908ccc785"
	from       = "+16193761185"
)

func sendSms(to string, message twiliogo.Body) {
	client := twiliogo.NewClient(accountSid, authToken)
	twiliogo.NewMessage(client, from, to, message)
}

func sendMms(to string, message twiliogo.MediaUrl) {
	client := twiliogo.NewClient(accountSid, authToken)
	twiliogo.NewMessage(client, from, to, message)
}

func main() {

	to := "+17605968806"
	message := "send me a command... (hint: go get catz)"
	sendSms(to, twiliogo.Body(message))

	http.HandleFunc("/hello", hello_handler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	response := r.FormValue("Body")
	to := r.FormValue("From")

	if response == "go get catz" || response == "catz" {
		sendMms(to, twiliogo.MediaUrl("http://thecatapi.com/api/images/get?format=src&type=gif"))
	} else {
		message := "http://thecatapi.com/?id=5jb&type=jpg"
		sendSms(to, twiliogo.Body(message))
	}
}
