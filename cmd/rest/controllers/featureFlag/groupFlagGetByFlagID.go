package featureflag

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *featureFlagController) GroupFlagGetMany(w http.ResponseWriter, r *http.Request) {
	var (
		orgSlug   = r.PathValue("orgSlug")
		appSlug   = r.PathValue("appSlug")
		flagIDStr = r.PathValue("flagID")
		flagID    int64
		err       error
	)
	if orgSlug == "" || appSlug == "" || flagIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	if flagID, err = strconv.ParseInt(flagIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	groupFlags, err := c.core.GroupFlagsGetByFlagID(r.Context(),
		c.core.NewGroupFlagsGetByFlagIDArgs(
			orgSlug,
			appSlug,
			flagID,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, groupFlags)
}
