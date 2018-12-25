package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/richard-xtek/e-com-fb/services/webhook/models"
)

const (
	FACEBOOK_API      = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
	IMAGE             = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
	VISIT_SHOW_URL    = "http://labouardy.com"
	PAGE_ACCESS_TOKEN = "EAAUpbN6bxIEBADGGHV4ZBfsZA0zougTpwaVsv4XVuJ4XvqboBZCjPiLMHFX1RCPJPU91r27fBfoZBVawdYn8oKzKEJZAWp2dK5CG6qgIb14lL8BJPFeI7vDx9GiA5fdHnESwsx7P5ut1ey61wWzloPNIqxibRvthwCug0v00DovpwXn1VdBB6"
)

func ProcessMessage(event Messaging) {
	var userQuery = event.Message.Text
	fmt.Println(event)
	fmt.Printf(userQuery)
	client := &http.Client{}

	var response Response = Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: Message{
			Text: event.Message.Text,
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	url := fmt.Sprintf(FACEBOOK_API, PAGE_ACCESS_TOKEN)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
		fmt.Printf("format string")
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error", err)
	}
	defer resp.Body.Close()

}
