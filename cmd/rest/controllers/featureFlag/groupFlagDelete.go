package featureflag

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *featureFlagController) GroupFlagDelete(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug    = r.PathValue("orgSlug")
		appSlug    = r.PathValue("appSlug")
		flagIDStr  = r.PathValue("flagID")
		flagID     int64
		groupIDStr = r.PathValue("groupID")
		groupID    int64
		err        error
	)
	if orgSlug == "" || appSlug == "" || flagIDStr == "" || groupIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	if flagID, err = strconv.ParseInt(flagIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}
	if groupID, err = strconv.ParseInt(groupIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	if err = c.core.GroupFlagDelete(r.Context(),
		c.core.NewGroupFlagDeleteArgs(
			orgSlug,
			groupID,
			appSlug,
			flagID,
		),
	); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.OK(w, r)
}
