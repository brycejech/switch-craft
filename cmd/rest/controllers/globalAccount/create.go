package globalaccount

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type createGlobalAccountArgs struct {
	IsInstanceAdmin bool   `json:"isInstanceAdmin"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

func (c *globalAccountController) Create(w http.ResponseWriter, r *http.Request) {
	body := &createGlobalAccountArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	account, err := c.core.GlobalAccountCreate(r.Context(),
		c.core.NewGlobalAccountCreateArgs(
			body.IsInstanceAdmin,
			body.FirstName,
			body.LastName,
			body.Email,
			body.Username,
			body.Password,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, account)
}
