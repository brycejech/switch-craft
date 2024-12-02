package application

import (
	"switchcraft/core"
	"switchcraft/types"
)

type appController struct {
	logger *types.Logger
	core   *core.Core
}

func NewAppController(logger *types.Logger, core *core.Core) *appController {
	return &appController{
		logger: logger,
		core:   core,
	}
}
