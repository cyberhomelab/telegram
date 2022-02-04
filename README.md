# Telegram client written in Go

## Config example

Create a file in `/opt/cyberhomelab/config.toml` or anywhere else (check "how to run" in this case)

```toml
LabName = "MyLab"

[Telegram]
Token = "<Token>"
ChatId = 12345678
MaxCharacters = 4000
```

## How to get the token or the chat id?

https://medium.com/geekculture/how-to-use-go-to-send-telegram-messages-to-your-phone-a819bdf7f35c

## How to run

```bash
export CYBERHOMELAB_CONFIG=$(pwd)/config.toml
go run main.go -message "Hello World!" 
```
