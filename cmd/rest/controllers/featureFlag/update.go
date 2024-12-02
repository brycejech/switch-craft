package featureflag

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

type featFlagUpdateArgs struct {
	Name      string `json:"name"`
	IsEnabled bool   `json:"isEnabled"`
}

func (c *featureFlagController) Update(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	flagIDStr := r.PathValue("flagID")
	if orgSlug == "" || appSlug == "" || flagIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var (
		flagID int64
		err    error
	)
	if flagID, err = strconv.ParseInt(flagIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	body := &featFlagUpdateArgs{}
	if err = restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	flag, err := c.core.FeatFlagUpdate(r.Context(),
		c.core.NewFeatFlagUpdateArgs(orgSlug, appSlug, flagID, body.Name, body.IsEnabled),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, flag)
}
