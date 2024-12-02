package org

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgs, err := c.core.OrgGetMany(r.Context())
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, orgs)
}
