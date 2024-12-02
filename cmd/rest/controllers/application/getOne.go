package application

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *appController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	app, err := c.core.AppGetOne(r.Context(),
		c.core.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, app)
}
