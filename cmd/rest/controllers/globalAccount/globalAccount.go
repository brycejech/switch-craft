package globalaccount

import (
	"switchcraft/core"
	"switchcraft/types"
)

type globalAccountController struct {
	logger *types.Logger
	core   *core.Core
}

func NewGlobalAccountController(logger *types.Logger, core *core.Core) *globalAccountController {
	return &globalAccountController{
		logger: logger,
		core:   core,
	}
}
