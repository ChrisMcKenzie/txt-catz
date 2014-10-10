package main

import (
	"log"
	"net/http"
	"os"

	"github.com/carlosdp/twiliogo"
)

// +1 442-242-6855

func main() {
	accountSid := "ACe3f0bf659491650bbff689790fd574e1"
	authToken := "c57d2b5b368d0dd291c3c5c908ccc785"
	client := twiliogo.NewClient(accountSid, authToken)

	from := "+16193761185"
	to := "+16195591274"
	message := "send me a command... (hint: go get catz)"
	twiliogo.NewMessage(client, from, to, twiliogo.Body(message))

	http.HandleFunc("/hello", hello_handler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v", r)

	response := r.FormValue("Body")
	accountSid := "ACe3f0bf659491650bbff689790fd574e1"
	authToken := "c57d2b5b368d0dd291c3c5c908ccc785"
	client := twiliogo.NewClient(accountSid, authToken)

	if response == "catz" {
		// http://thecatapi.com/?id=5jb&type=jpg
		from := "+16193761185"
		to := r.FormValue("From")
		message := "http://thecatapi.com/?id=5jb&type=jpg"
		twiliogo.NewMessage(client, from, to, twiliogo.MediaUrl(message))
	} else {
		from := "+16193761185"
		to := r.FormValue("From")
		message := "http://thecatapi.com/?id=5jb&type=jpg"
		twiliogo.NewMessage(client, from, to, twiliogo.Body(message))
	}

}
