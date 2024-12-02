package org

import (
	"encoding/json"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type orgCreateArgs struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *orgController) Create(w http.ResponseWriter, r *http.Request) {
	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
	body := &orgCreateArgs{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	args := c.core.NewOrgCreateArgs(body.Name, body.Slug, tracer.AuthAccount.ID)

	org, err := c.core.OrgCreate(r.Context(), args)
	if err != nil {
		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, org)
}
