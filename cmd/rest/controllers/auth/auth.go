package auth

import (
	"switchcraft/core"
	"switchcraft/types"
)

type authController struct {
	logger *types.Logger
	core   *core.Core
}

func NewAuthController(logger *types.Logger, core *core.Core) *authController {
	return &authController{
		logger: logger,
		core:   core,
	}
}
