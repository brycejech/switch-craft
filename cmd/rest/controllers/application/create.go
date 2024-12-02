package application

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type appCreateArgs struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *appController) Create(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.BadRequest(w, r)
		return
	}

	body := &appCreateArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	app, err := c.core.AppCreate(r.Context(),
		c.core.NewAppCreateArgs(orgSlug, body.Name, body.Slug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
		} else if errors.Is(err, types.ErrItemExists) {
			restutils.BadRequest(w, r)
		} else {
			restutils.InternalServerError(w, r)
		}
		return
	}

	restutils.Render(w, r, http.StatusOK, app)
}
