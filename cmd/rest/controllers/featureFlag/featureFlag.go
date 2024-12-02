package featureflag

import (
	"switchcraft/core"
	"switchcraft/types"
)

type featureFlagController struct {
	logger *types.Logger
	core   *core.Core
}

func NewFeatureFlagController(logger *types.Logger, core *core.Core) *featureFlagController {
	return &featureFlagController{
		logger: logger,
		core:   core,
	}
}
