/*
   Copyright (c) 2022 Cyber Home Lab authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	Token  string
	ChatId string
)

type Body struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}
type Result struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Message struct {
	MessageId int    `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}
type From struct {
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", CurrentConfig.Telegram.Token)
}

func convertToBody(input io.ReadCloser) (Body, error) {
	var body Body

	// Convert to bytes
	bodyBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return Body{}, err
	}

	// Convert to string
	bodyString := string(bodyBytes)

	// Convert to Body{}
	err = json.Unmarshal([]byte(bodyString), &body)
	if err != nil {
		return Body{}, fmt.Errorf("couldn't convert to the Body struct -> %s", err)
	}

	return body, nil
}

func GetMessages() (Body, error) {
	// Get the messages
	url := fmt.Sprintf("%s/getUpdates", getUrl())
	response, err := http.Post(url, "application/json", nil)
	if err != nil {
		return Body{}, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Get the body
	body, err := convertToBody(response.Body)
	if err != nil {
		return Body{}, fmt.Errorf("couldn't convert to the Body struct -> %s", err)
	}
	if !body.Ok {
		return Body{}, fmt.Errorf("couldn't get the messages, the response received is not ok")
	}

	// Return
	return body, nil
}

func SendMessage(text string) error {
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	bodyBytesSend, _ := json.Marshal(map[string]string{
		"chat_id": fmt.Sprint(CurrentConfig.Telegram.ChatId),
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(bodyBytesSend),
	)
	if err != nil {
		return err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	bodyString := string(body)
	if !strings.Contains(bodyString, "ok") {
		return fmt.Errorf("couldn't send the message")
	} else {
		fmt.Println("Message was sent successfully")
	}

	return nil
}
