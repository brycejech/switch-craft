package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

type appController struct {
	logger *types.Logger
	core   *core.Core
}

func newAppController(logger *types.Logger, core *core.Core) *appController {
	return &appController{
		logger: logger,
		core:   core,
	}
}

type appCreateArgs struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *appController) Create(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		badRequest(w, r)
		return
	}

	body := &appCreateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	app, err := c.core.AppCreate(r.Context(),
		c.core.NewAppCreateArgs(orgSlug, body.Name, body.Slug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
		} else if errors.Is(err, types.ErrItemExists) {
			badRequest(w, r)
		} else {
			internalServerError(w, r)
		}
		return
	}

	render(w, r, http.StatusOK, app)
}

func (c *appController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		notFound(w, r)
		return
	}

	apps, err := c.core.AppGetMany(r.Context(), orgSlug)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
			return
		}
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, apps)
}

func (c *appController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		notFound(w, r)
		return
	}

	app, err := c.core.AppGetOne(r.Context(),
		c.core.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
		} else {
			internalServerError(w, r)
		}
		return
	}

	render(w, r, http.StatusOK, app)
}

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
		notFound(w, r)
		return
	}

	body := &appUpdateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	existingApp, err := c.core.AppGetOne(r.Context(),
		c.core.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
		} else {
			internalServerError(w, r)
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
		badRequest(w, r)
		return
	}

	app, err := c.core.AppUpdate(r.Context(),
		c.core.NewAppUpdateArgs(orgSlug, existingApp.ID, body.Name, body.Slug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
		} else if errors.Is(err, types.ErrItemExists) {
			badRequest(w, r)
		} else {
			internalServerError(w, r)
		}
		return
	}

	render(w, r, http.StatusOK, app)
}

func (c *appController) Delete(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	if orgSlug == "" || appSlug == "" {
		notFound(w, r)
		return
	}

	if err := c.core.AppDelete(r.Context(), orgSlug, appSlug); err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
		} else {
			internalServerError(w, r)
		}
		return
	}

	render(w, r, http.StatusOK, nil)
}
