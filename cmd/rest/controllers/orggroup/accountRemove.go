package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) AccountRemove(w http.ResponseWriter, r *http.Request) {
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

	if err = c.core.OrgGroupAccountRemove(r.Context(), orgSlug, groupID, accountID); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.OK(w, r)
}
