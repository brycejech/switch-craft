package globalaccount

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
)

func (c *globalAccountController) GetOne(w http.ResponseWriter, r *http.Request) {
	accountIDStr := r.PathValue("accountID")
	if accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	var accountID int64
	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)
	if err != nil {
		restutils.BadRequest(w, r)
		return
	}

	account, err := c.core.GlobalAccountGetOne(r.Context(),
		c.core.NewGlobalAccountGetOneArgs(&accountID, nil, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, account)
}
