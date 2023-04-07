package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
}

func mustToken() string {
	token := flag.String("t", "", "token for access to telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}
