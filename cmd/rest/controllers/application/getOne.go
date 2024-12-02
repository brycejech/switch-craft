package application

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
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
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
		} else {
			restutils.InternalServerError(w, r)
		}
		return
	}

	restutils.Render(w, r, http.StatusOK, app)
}
