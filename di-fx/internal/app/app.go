package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewServer() *echo.Echo {
	return echo.New()
}

func StartServer(lc fx.Lifecycle, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := e.Start(":1323"); err != nil {
					e.Logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

func StartApp() {
	fx.New(
		fx.Provide(
			NewServer,
		),
		Modules,
		fx.Invoke(StartServer),
	).Run()
}
