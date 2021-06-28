package main

import (
	"log"

	"github.com/teitei-tk/goodbyte-twitter-history/cli"
)

func main() {
	flags := cli.ParseFlags()

	app := InitializeApp(flags)
	if err := app.Run(); err != nil {
		log.Fatalf("app fail. %v", err)
	}
}
