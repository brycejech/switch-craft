package globalaccount

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *globalAccountController) Delete(w http.ResponseWriter, r *http.Request) {
	accountIDStr := r.PathValue("accountID")
	if accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var (
		accountID int64
		err       error
	)
	if accountID, err = strconv.ParseInt(accountIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	if err := c.core.GlobalAccountDelete(r.Context(), accountID); err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.OK(w, r)
}
