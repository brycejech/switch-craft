package orggroup

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type updateOrgGroupArgs struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *orgGroupController) Update(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	groupIDStr := r.PathValue("groupID")
	if orgSlug == "" || groupIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &updateOrgGroupArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	var (
		groupID int64
		err     error
	)
	if groupID, err = strconv.ParseInt(groupIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	group, err := c.core.OrgGroupGetOne(r.Context(),
		c.core.NewOrgGroupGetOneArgs(orgSlug, &groupID, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)

	if group.ID != body.ID {
		restutils.BadRequest(w, r)
		c.logger.Warn(tracer, "Org group update ID mismatch detected", map[string]any{
			"user":          tracer.AuthAccount.Username,
			"requestBody":   body,
			"existingGroup": group,
		})
		return
	}

	updatedGroup, err := c.core.OrgGroupUpdate(r.Context(),
		c.core.NewOrgGroupUpdateArgs(
			orgSlug,
			group.ID,
			body.Name,
			body.Description,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, updatedGroup)
}
