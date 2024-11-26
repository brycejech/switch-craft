package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	tenantID, err := strconv.ParseInt(r.PathValue("tenantID"), 10, 64)
	if err != nil {
		badRequest(w, r)
		return
	}

	tenant, err := c.core.TenantGetOne(r.Context(), tenantID)
	if err != nil {
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

	tenantIDFromURL, err := strconv.ParseInt(r.PathValue("tenantID"), 10, 64)
	if err != nil {
		badRequest(w, r)
		return
	}
	tenant, err := c.core.TenantGetOne(r.Context(), tenantIDFromURL)
	if err != nil || tenant == nil {
		notFound(w, r)
		return
	}

	args := c.core.NewTenantUpdateArgs(tenantIDFromURL, body.Name, body.Slug, tenant.Owner)
	updatedTenant, err := c.core.TenantUpdate(r.Context(), args)
	if err != nil {
		internalServerError(w, r)
		return
	}

	render(w, r, http.StatusOK, updatedTenant)
}
