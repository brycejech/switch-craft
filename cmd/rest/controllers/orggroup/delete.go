package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) Delete(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	groupIDStr := r.PathValue("groupID")
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

	if err := c.core.OrgGroupDelete(r.Context(), orgSlug, groupID); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.OK(w, r)
}
