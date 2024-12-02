package application

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

func (c *appController) Delete(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	if err := c.core.AppDelete(r.Context(), orgSlug, appSlug); err != nil {
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
		} else {
			restutils.InternalServerError(w, r)
		}
		return
	}

	restutils.Render(w, r, http.StatusOK, nil)
}
