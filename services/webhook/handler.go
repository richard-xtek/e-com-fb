package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	. "github.com/richard-xtek/e-com-fb/services/webhook/models"
	// . "github.com/mlabouardy/dialogflow-watchnow-messenger/models"
)

var verifyToken = "EAAFhMnbuC00BAAYJQTHNzqDFAQWaeZAfkwg9OJCxmnzTV8vfIALoVLS2bC1z9EidmFiWCgQu8QjbLVx6nXZAzZBUjX6ZAJS3xL4VRmrKKN08l1CKXuQmENnYZCcp5JoPOrmp3xDnuO7kMkSqGwlfm90DPiiILr20j7NgOZAZATbOQuG1zWaPq2qVVdfWev15i0ZD"

// VerificationEndPoint endpoint use abc xyz
func VerificationEndPoint(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	if mode != "" && token == verifyToken {
		w.WriteHeader(200)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error, wrong validation token"))
	}
}

func MessagesEndPoint(w http.ResponseWriter, r *http.Request) {
	var callback Callback
	fmt.Printf("called \n")
	json.NewDecoder(r.Body).Decode(&callback)
	fmt.Println(callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				if !reflect.DeepEqual(event.Message, Message{}) && event.Message.Text != "" {
					ProcessMessage(event)
				}
			}
		}
		w.WriteHeader(200)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Message not supported"))
	}
}
