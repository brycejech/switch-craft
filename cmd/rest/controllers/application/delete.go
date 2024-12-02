package application

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *appController) Delete(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	if err := c.core.AppDelete(r.Context(), orgSlug, appSlug); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, nil)
}
