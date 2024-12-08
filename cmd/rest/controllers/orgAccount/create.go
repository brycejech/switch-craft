package orgaccount

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type createOrgAccountArgs struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  *string `json:"password"`
}

func (c *orgAccountController) Create(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &createOrgAccountArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	account, err := c.core.OrgAccountCreate(r.Context(),
		c.core.NewOrgAccountCreateArgs(
			orgSlug,
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
