package org

import (
	"errors"
	"net/http"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type orgUpdateArgs struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *orgController) Update(w http.ResponseWriter, r *http.Request) {
	var orgSlug string
	if orgSlug = r.PathValue("orgSlug"); orgSlug == "" {
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
		if errors.Is(err, types.ErrNotFound) {
			restutils.NotFound(w, r)
		} else {
			restutils.InternalServerError(w, r)
		}
		return
	}

	updatedOrg, err := c.core.OrgUpdate(r.Context(),
		c.core.NewOrgUpdateArgs(org.ID, body.Name, body.Slug, org.Owner),
	)
	if err != nil {
		restutils.InternalServerError(w, r)
		return
	}

	restutils.Render(w, r, http.StatusOK, updatedOrg)
}
