package account

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *accountController) DeleteOrgAccount(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	accountIDStr := r.PathValue("accountID")
	if orgSlug == "" || accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var (
		accountID int64
		err       error
	)
	if accountID, err = strconv.ParseInt(accountIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	if err := c.core.AccountDelete(r.Context(), &orgSlug, accountID); err != nil {
		restutils.HandleCoreErr(w, r, err)
	}

	restutils.OK(w, r)
}
