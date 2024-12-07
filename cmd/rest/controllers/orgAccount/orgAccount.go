package orgaccount

import (
	"switchcraft/core"
	"switchcraft/types"
)

type orgAccountController struct {
	logger *types.Logger
	core   *core.Core
}

func NewOrgAccountController(logger *types.Logger, core *core.Core) *orgAccountController {
	return &orgAccountController{
		logger: logger,
		core:   core,
	}
}
