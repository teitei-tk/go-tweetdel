package app

import (
	"github.com/google/wire"
	"github.com/teitei-tk/goodbyte-twitter-history/cli"
)

var AppSet = wire.NewSet(
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