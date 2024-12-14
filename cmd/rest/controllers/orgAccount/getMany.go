package orgaccount

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *orgAccountController) GetMany(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	accounts, err := c.core.OrgAccountGetMany(r.Context(), orgSlug)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}

func (c *orgAccountController) GetManyByID(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	if orgSlug == "" {
		restutils.NotFound(w, r)
		return
	}

	idStrs := r.URL.Query()["id"]
	if len(idStrs) < 1 {
		restutils.NotFound(w, r)
		return
	}

	ids := make([]int64, len(idStrs))
	for i, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			restutils.BadRequest(w, r)
			return
		}
		ids[i] = id
	}

	accounts, err := c.core.OrgAccountGetManyByID(r.Context(), orgSlug, ids)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}
