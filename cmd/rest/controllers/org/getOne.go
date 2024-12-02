package org

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	org, err := c.core.OrgGetOne(r.Context(),
		c.core.NewOrgGetOneArgs(nil, nil, &orgSlug),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, org)
}
