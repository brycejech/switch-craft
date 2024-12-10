package featureflag

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type featFlagUpdateArgs struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"Description"`
	IsEnabled   bool   `json:"isEnabled"`
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

	if flagID != body.ID {
		tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
		c.logger.Warn(tracer, "Potential breakout detected, feature flag ID mismatch", map[string]any{
			"user":          tracer.AuthAccount.Username,
			"urlParamId":    flagID,
			"requestBodyID": body.ID,
		})
		restutils.BadRequest(w, r)
		return
	}

	flag, err := c.core.FeatFlagUpdate(r.Context(),
		c.core.NewFeatFlagUpdateArgs(
			orgSlug,
			appSlug,
			flagID,
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
