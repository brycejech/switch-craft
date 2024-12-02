package featureflag

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type featFlagCreateArgs struct {
	Name      string `json:"name"`
	IsEnabled bool   `json:"isEnabled"`
}

func (c *featureFlagController) Create(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &featFlagCreateArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	flag, err := c.core.FeatFlagCreate(r.Context(),
		c.core.NewFeatFlagCreateArgs(orgSlug, appSlug, body.Name, body.IsEnabled),
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

	restutils.Render(w, r, http.StatusOK, flag)
}
