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
	"fmt"
	"os"
	"reflect"
	"strconv"

	toml "github.com/pelletier/go-toml"
)

var (
	CurrentConfig Config
	Hostname      string
)

type Config struct {
	LabName  string
	Telegram struct {
		Token         string
		ChatId        int
		MaxCharacters int
	}
}

func GetConfig(cfgFile string) (Config, error) {
	// Open the config file
	file, err := os.Open(cfgFile)
	if err != nil {
		return Config{}, fmt.Errorf("config file %s can't be opened -> %s", cfgFile, err)
	}

	// Close the config file at the end
	defer file.Close()

	// Decode the configuration file
	cfg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(cfg); err != nil {
		return Config{}, fmt.Errorf("can't decode the configuration file -> %s", err)
	}

	// Return the config
	return *cfg, nil
}

func (c *Config) CheckConfig() error {
	// Check LabName
	if StringIsEmpty(c.LabName) {
		return fmt.Errorf("string LabName is empty")
	}

	// Check Token
	if StringIsEmpty(c.Telegram.Token) {
		return fmt.Errorf("string Telegram.Token is empty")
	}

	// Check ChatId and MaxCharacters
	for _, key := range []string{"ChatId", "MaxCharacters"} {
		if IntegerIsEmpty(reflect.ValueOf(c.Telegram).FieldByName(key).Int()) {
			return fmt.Errorf("integer Telegram.%s is empty", key)
		}
	}

	// Default
	return nil
}

func init() {
	var configPath string
	var err error

	// Config
	envToken, _ := GetEnvVariable("TELEGRAM_TOKEN")
	envChatId, _ := GetEnvVariable("TELEGRAM_CHAT_ID")
	if StringIsEmpty(envToken) && StringIsEmpty(envChatId) {
		// Use config from file
		configPath, err = GetEnvVariable("CYBERHOMELAB_CONFIG")
		if err != nil {
			configPath = ConfigFilePath
		}
		CurrentConfig, err = GetConfig(configPath)
		if err != nil {
			fmt.Printf("ERROR: Couldn't get the config -> %s", err)
			os.Exit(2)
		}
		err = CurrentConfig.CheckConfig()
		if err != nil {
			fmt.Printf("ERROR: There is an issue in the config -> %s", err)
			os.Exit(2)
		}
	} else {
		// Use a config from the environment variables
		envLabName, err := GetEnvVariable("LAB_NAME")
		if err == nil {
			CurrentConfig.LabName = envLabName
		} else {
			CurrentConfig.LabName = "My Lab"
		}

		CurrentConfig.Telegram.Token = envToken

		CurrentConfig.Telegram.ChatId, err = strconv.Atoi(envChatId)
		if err != nil {
			fmt.Println("ERROR: Value from the environment variable TELEGRAM_CHAT_ID coudln't be converted to int")
			os.Exit(2)
		}

		envMaxChar, err := GetEnvVariable("TELEGRAM_MAX_CHARACTERS")
		if err == nil {
			CurrentConfig.Telegram.MaxCharacters, err = strconv.Atoi(envMaxChar)
			if err != nil {
				CurrentConfig.Telegram.MaxCharacters = 4000
			}
		} else {
			CurrentConfig.Telegram.MaxCharacters = 4000
		}
	}

	// Hostname
	Hostname, err = os.Hostname()
	if err != nil {
		Hostname = "Unknown"
		fmt.Printf("WARNING: Couldn't get the hostname -> %s", err)
	}
}
