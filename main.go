package main

import (
	"fmt"
	"os"
)

const envKeyVar = "bot_password_key"

func main() {
	token := os.Getenv(envKeyVar)

	if len(token) == 0 {
		fmt.Println(fmt.Printf("The environment variable %s is not defined", envKeyVar))
		os.Exit(1)
	}

	tbot, err := NewBot(token)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//tbot.Debug(true)

	tbot.Start()
}
