package main

import (
	"github.com/google/wire"
	"github.com/teitei-tk/go-tweetdel/app"
	"github.com/teitei-tk/go-tweetdel/cli"
)

func appConfig(flags *cli.CliFlags) *app.AppConf {
	return &app.AppConf{
		Input: flags,
	}
}

func InitializeApp(flags *cli.CliFlags) *app.App {
	wire.Build(app.AppSet)

	return nil
}
