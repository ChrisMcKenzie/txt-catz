package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sfreiberg/gotwilio"
)

// +1 442-242-6855

func main() {
	accountSid := "ACe3f0bf659491650bbff689790fd574e1"
	authToken := "c57d2b5b368d0dd291c3c5c908ccc785"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+16193761185"
	to := "+17605968806"
	message := "send me a command... (hint: go get catz)"
	twilio.SendSMS(from, to, message, "", "")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%#v", r)
	})

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
