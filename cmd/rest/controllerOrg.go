package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

type orgController struct {
	logger *types.Logger
	core   *core.Core
}

func newOrgController(logger *types.Logger, core *core.Core) *orgController {
	return &orgController{
		logger: logger,
		core:   core,
	}
}

type orgCreateArgs struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *orgController) Create(w http.ResponseWriter, r *http.Request) {
	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
	body := &orgCreateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	args := c.core.NewOrgCreateArgs(body.Name, body.Slug, tracer.AuthAccount.ID)

	org, err := c.core.OrgCreate(r.Context(), args)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, org)
}

func (c *orgController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgs, err := c.core.OrgGetMany(r.Context())
	if err != nil {
		fmt.Println(err)
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, orgs)
}

func (c *orgController) GetOne(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	org, err := c.core.OrgGetOne(r.Context(),
		c.core.NewOrgGetOneArgs(nil, nil, &orgSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
			return
		}

		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, org)
}

type orgUpdateArgs struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *orgController) Update(w http.ResponseWriter, r *http.Request) {
	body := &orgUpdateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	var orgSlug string
	if orgSlug := r.PathValue("orgSlug"); orgSlug == "" {
		notFound(w, r)
		return
	}

	org, err := c.core.OrgGetOne(r.Context(),
		c.core.NewOrgGetOneArgs(nil, nil, &orgSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
			return
		}
		internalServerError(w, r)
		return
	}

	updatedOrg, err := c.core.OrgUpdate(r.Context(),
		c.core.NewOrgUpdateArgs(org.ID, body.Name, body.Slug, org.Owner),
	)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, updatedOrg)
}
