package account

import (
	"switchcraft/core"
	"switchcraft/types"
)

type accountController struct {
	logger *types.Logger
	core   *core.Core
}

func NewAccountController(logger *types.Logger, core *core.Core) *accountController {
	return &accountController{
		logger: logger,
		core:   core,
	}
}
