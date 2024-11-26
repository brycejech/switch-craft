package rest

import (
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

type authController struct {
	logger *types.Logger
	core   *core.Core
}

func newAuthController(logger *types.Logger, core *core.Core) *authController {
	return &authController{
		logger: logger,
		core:   core,
	}
}

type authnArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authnResponse struct {
	Token string `json:"token"`
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	args := new(authnArgs)
	if err := decodeBody(r, args); err != nil {
		badRequest(w, r)
		return
	}

	account, ok := c.core.Authn(r.Context(), args.Username, args.Password)
	if !ok {
		unauthorized(w, r)
		return
	}

	token, err := c.core.AuthCreateJWT(account)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, authnResponse{Token: token})
}
