package org

import (
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
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	org, err := c.core.OrgCreate(r.Context(),
		c.core.NewOrgCreateArgs(body.Name, body.Slug, tracer.AuthAccount.ID),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, org)
}
