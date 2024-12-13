package orggroup

import (
	"switchcraft/core"
	"switchcraft/types"
)

type orgGroupController struct {
	logger *types.Logger
	core   *core.Core
}

func NewOrgGroupController(logger *types.Logger, core *core.Core) *orgGroupController {
	return &orgGroupController{
		logger: logger,
		core:   core,
	}
}
