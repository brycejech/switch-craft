package account

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *accountController) GetOrgAccounts(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	accounts, err := c.core.AccountGetMany(r.Context(), &orgSlug)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}
