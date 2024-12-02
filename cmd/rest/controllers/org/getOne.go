package org

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

func (c *orgController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	org, err := c.core.OrgGetOne(r.Context(),
		c.core.NewOrgGetOneArgs(nil, nil, &orgSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
			return
		}

		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, org)
}
