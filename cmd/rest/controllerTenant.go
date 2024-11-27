package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

type tenantController struct {
	logger *types.Logger
	core   *core.Core
}

func newTenantController(logger *types.Logger, core *core.Core) *tenantController {
	return &tenantController{
		logger: logger,
		core:   core,
	}
}

func (c *tenantController) GetMany(w http.ResponseWriter, r *http.Request) {
	tenants, err := c.core.TenantGetMany(r.Context())
	if err != nil {
		fmt.Println(err)
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, tenants)
}

func (c *tenantController) GetOne(w http.ResponseWriter, r *http.Request) {
	tenantSlug := r.PathValue("tenantSlug")
	tenant, err := c.core.TenantGetOne(r.Context(),
		c.core.NewTenantGetOneArgs(nil, nil, &tenantSlug),
	)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			notFound(w, r)
			return
		}

		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, tenant)
}

type tenantCreateArgs struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *tenantController) Create(w http.ResponseWriter, r *http.Request) {
	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
	body := &tenantCreateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	args := c.core.NewTenantCreateArgs(body.Name, body.Slug, tracer.AuthAccount.ID)

	tenant, err := c.core.TenantCreate(r.Context(), args)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, tenant)
}

type tenantUpdateArgs struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *tenantController) Update(w http.ResponseWriter, r *http.Request) {
	body := &tenantUpdateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		jsonParseError(w, r)
		return
	}

	tenantSlug := r.PathValue("tenantID")
	tenant, err := c.core.TenantGetOne(r.Context(),
		c.core.NewTenantGetOneArgs(nil, nil, &tenantSlug),
	)
	if err != nil || tenant == nil {
		notFound(w, r)
		return
	}

	updatedTenant, err := c.core.TenantUpdate(r.Context(),
		c.core.NewTenantUpdateArgs(tenant.ID, body.Name, body.Slug, tenant.Owner),
	)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, updatedTenant)
}
