package main

import (
	"log"
	"os"

	"github.com/teitei-tk/go-tweetdel/cli"
)

func main() {
	flags, err := cli.ParseFlags(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	app := InitializeApp(flags)
	if err := app.Run(); err != nil {
		log.Fatalf("app fail. %v", err)
	}
}
