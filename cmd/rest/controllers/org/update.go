package org

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type orgUpdateArgs struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *orgController) Update(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &orgUpdateArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	org, err := c.core.OrgGetOne(r.Context(),
		c.core.NewOrgGetOneArgs(nil, nil, &orgSlug),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	updatedOrg, err := c.core.OrgUpdate(r.Context(),
		c.core.NewOrgUpdateArgs(org.ID, body.Name, body.Slug, org.Owner),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, updatedOrg)
}
