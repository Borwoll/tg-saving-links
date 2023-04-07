package main

import (
	"flag"
	"log"

	"tg-saving-links/clients/telegram/telegram"
)

const (
	host = "api.telegram.org"
)

func main() {
	c := telegram.NewClient(mustToken())
}

func mustToken() string {
	token := flag.String("t", "", "token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}
