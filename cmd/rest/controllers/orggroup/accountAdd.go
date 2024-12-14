package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) AccountAdd(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug      = r.PathValue("orgSlug")
		groupIDStr   = r.PathValue("groupID")
		accountIDStr = r.PathValue("accountID")
	)
	if orgSlug == "" || groupIDStr == "" || accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var (
		groupID   int64
		accountID int64
		err       error
	)
	if groupID, err = strconv.ParseInt(groupIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}
	if accountID, err = strconv.ParseInt(accountIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	groupAccount, err := c.core.OrgGroupAccountAdd(r.Context(), c.core.NewOrgGroupAccountAddArgs(orgSlug, groupID, accountID))
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, groupAccount)
}
