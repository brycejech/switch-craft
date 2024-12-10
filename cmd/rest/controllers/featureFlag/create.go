package featureflag

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type featFlagCreateArgs struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	IsEnabled   bool   `json:"isEnabled"`
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
		c.core.NewFeatFlagCreateArgs(orgSlug,
			appSlug,
			body.Name,
			body.Label,
			body.Description,
			body.IsEnabled,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, flag)
}
