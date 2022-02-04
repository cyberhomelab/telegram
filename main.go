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

package main

import (
	"flag"
	"fmt"
	"os"

	"cyberhomelab.com/telegram/telegram"
)

type Flags struct {
	Severity string
	Message  string
}

var (
	currentFlags Flags
)

func init() {
	flag.StringVar(&currentFlags.Severity, "severity", "Info", "The severity level of the message")
	flag.StringVar(&currentFlags.Message, "message", "", "The actual message that will be sent")
	flag.Parse()
}

func main() {
	if telegram.StringIsEmpty(currentFlags.Message) {
		fmt.Printf("Message is empty")
		os.Exit(2)
	}

	err := telegram.SendMessage(currentFlags.Message)
	if err != nil {
		fmt.Printf("Message couldn't be sent -> %s", err)
		os.Exit(2)
	}
}
