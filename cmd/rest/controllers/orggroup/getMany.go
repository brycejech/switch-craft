package orggroup

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgGroupController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	groups, err := c.core.OrgGroupGetMany(r.Context(), orgSlug)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, groups)
}
