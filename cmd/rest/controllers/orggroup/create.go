package orggroup

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

type createOrgGroupArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *orgGroupController) Create(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &createOrgGroupArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	group, err := c.core.OrgGroupCreate(r.Context(),
		c.core.NewOrgGroupCreateArgs(
			orgSlug,
			body.Name,
			body.Description,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, group)
}
