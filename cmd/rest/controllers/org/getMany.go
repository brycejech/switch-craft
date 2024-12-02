package org

import (
	"fmt"
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgs, err := c.core.OrgGetMany(r.Context())
	if err != nil {
		fmt.Println(err)
		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, orgs)
}
