package application

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type appUpdateArgs struct {
	OrgID int64  `json:"orgId"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
}

func (c *appController) Update(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &appUpdateArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	existingApp, err := c.core.AppGetOne(r.Context(),
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

	if existingApp.ID != body.ID {
		tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
		c.logger.Warn(tracer, "Potential tenant breakout detected, application ID mismatch", map[string]any{
			"user":          tracer.AuthAccount.Username,
			"existingSlug":  existingApp.Slug,
			"existingID":    existingApp.ID,
			"requestBodyID": body.ID,
		})
		restutils.BadRequest(w, r)
		return
	}

	app, err := c.core.AppUpdate(r.Context(),
		c.core.NewAppUpdateArgs(orgSlug, existingApp.ID, body.Name, body.Slug),
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
