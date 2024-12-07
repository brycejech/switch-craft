package orgaccount

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgAccountController) GetOneOrgAccount(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	accountIDStr := r.PathValue("accountID")
	if orgSlug == "" || accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var accountID int64
	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		restutils.BadRequest(w, r)
		return
	}

	account, err := c.core.OrgAccountGetOne(r.Context(),
		c.core.NewOrgAccountGetOneArgs(orgSlug, &accountID, nil, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, account)
}
