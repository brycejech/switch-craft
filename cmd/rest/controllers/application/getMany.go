package application

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *appController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	apps, err := c.core.AppGetMany(r.Context(), orgSlug)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, apps)
}
