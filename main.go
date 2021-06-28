package main

import (
	"log"

	"github.com/google/wire"
	"github.com/teitei-tk/goodbyte-twitter-history/cli"
)

var appSet = wire.NewSet(
	wire.Struct(new(AppConf), "*"),
	NewApp,
)

type AppConf struct {
	Input *cli.CliFlags
}

type App struct {
	params *AppConf
}

func NewApp(p *AppConf) *App {
	return &App{p}
}

func (a *App) Run() error {
	return nil
}

func main() {
	flags := cli.ParseFlags()

	app := InitializeApp(flags)
	if err := app.Run(); err != nil {
		log.Fatalf("app fail. %v", err)
	}
}
