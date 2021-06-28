package main

import (
	"github.com/google/wire"
	"github.com/teitei-tk/goodbyte-twitter-history/app"
	"github.com/teitei-tk/goodbyte-twitter-history/cli"
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
