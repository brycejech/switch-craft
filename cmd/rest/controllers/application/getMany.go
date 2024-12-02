package application

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

func (c *appController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	apps, err := c.core.AppGetMany(r.Context(), orgSlug)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
			return
		}
		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, apps)
}
