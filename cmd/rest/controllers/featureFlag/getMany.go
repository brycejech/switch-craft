package featureflag

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *featureFlagController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	flags, err := c.core.FeatFlagGetMany(r.Context(), orgSlug, appSlug)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, flags)
}
