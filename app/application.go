package app

import (
	"go.uber.org/fx"
	"log"
	"robbo-assets/app/modules"
	"robbo-assets/package/config"
	"robbo-assets/package/logger"
	"robbo-assets/server"
)

func InvokeWith(options ...fx.Option) *fx.App {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	var di = []fx.Option{
		fx.Provide(logger.NewLogger),
		fx.Provide(modules.SetupHandler),
	}
	for _, option := range options {
		di = append(di, option)
	}
	return fx.New(di...)
}

func RunApp() {
	InvokeWith(
		fx.Invoke(server.NewServer),
	).Run()
}
