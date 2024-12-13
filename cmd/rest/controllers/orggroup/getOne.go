package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	groupIDStr := r.PathValue("groupID")
	if orgSlug == "" || groupIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var groupID int64
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		restutils.BadRequest(w, r)
		return
	}

	group, err := c.core.OrgGroupGetOne(r.Context(),
		c.core.NewOrgGroupGetOneArgs(orgSlug, &groupID, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, group)
}
