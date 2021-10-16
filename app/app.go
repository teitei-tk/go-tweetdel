package app

import (
	"github.com/google/wire"
	"github.com/teitei-tk/go-tweetdel/cli"
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
	switch a.params.Input.RunMode {
	case cli.RunMode_Detect:
		if err := NewDetectMode(a); err != nil {
			return err
		}
	}

	return nil
}
