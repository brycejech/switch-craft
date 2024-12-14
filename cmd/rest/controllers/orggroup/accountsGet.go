package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) AccountsGet(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug    = r.PathValue("orgSlug")
		groupIDStr = r.PathValue("groupID")
	)
	if orgSlug == "" || groupIDStr == "" {
		restutils.NotFound(w, r)
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

	accounts, err := c.core.OrgGroupAccountGetAll(r.Context(), orgSlug, groupID)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}
