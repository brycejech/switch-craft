package featureflag

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *featureFlagController) Delete(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	appSlug := r.PathValue("appSlug")
	flagIDStr := r.PathValue("flagID")
	if orgSlug == "" || appSlug == "" || flagIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var (
		flagID int64
		err    error
	)
	if flagID, err = strconv.ParseInt(flagIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	err = c.core.FeatFlagDelete(r.Context(), orgSlug, appSlug, flagID)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.OK(w, r)
}
