package auth

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type authnArgs struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authnResponse struct {
	Token string `json:"token"`
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	args := new(authnArgs)
	if err := restutils.DecodeBody(r, args); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	account, ok := c.core.Authn(r.Context(), args.Username, args.Password)
	if !ok {
		restutils.Unauthorized(w, r)
		return
	}

	token, err := c.core.AuthCreateJWT(account)
	if err != nil {
		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, authnResponse{Token: token})
}
