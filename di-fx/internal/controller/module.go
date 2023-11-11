package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAnyController),
	fx.Invoke(RegisterAnyController),
)