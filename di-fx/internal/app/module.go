package app

import (
	"github.com/labtrilhadev/go-tutorials/di-fx/internal/controller"
	"github.com/labtrilhadev/go-tutorials/di-fx/internal/usecase"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	controller.Module,
	usecase.Module,
)
