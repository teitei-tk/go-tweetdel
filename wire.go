package main

import (
	"github.com/google/wire"
	"github.com/teitei-tk/goodbyte-twitter-history/cli"
)

func appConfig(flags *cli.CliFlags) *AppConf {
	return &AppConf{
		Input: flags,
	}
}

func InitializeApp(flags *cli.CliFlags) *App {
	wire.Build(appSet)

	return nil
}
