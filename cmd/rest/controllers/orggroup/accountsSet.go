package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) AccountsSet(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug    = r.PathValue("orgSlug")
		groupIDStr = r.PathValue("groupID")
	)
	if orgSlug == "" || groupIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var accountIDs []int64
	if err := restutils.DecodeBody(r, &accountIDs); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	var (
		groupID int64
		err     error
	)
	if groupID, err = strconv.ParseInt(groupIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	accounts, err := c.core.OrgGroupAccountsSet(r.Context(), c.core.NewOrgGroupAccountsSetArgs(orgSlug, groupID, accountIDs))
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}
