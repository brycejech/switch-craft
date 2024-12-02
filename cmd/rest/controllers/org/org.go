package org

import (
	"switchcraft/core"
	"switchcraft/types"
)

type orgController struct {
	logger *types.Logger
	core   *core.Core
}

func NewOrgController(logger *types.Logger, core *core.Core) *orgController {
	return &orgController{
		logger: logger,
		core:   core,
	}
}
