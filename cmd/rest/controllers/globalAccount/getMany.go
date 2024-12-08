package globalaccount

import (
	"net/http"
	"switchcraft/cmd/rest/restutils"
)

func (c *globalAccountController) GetMany(w http.ResponseWriter, r *http.Request) {
	accounts, err := c.core.GlobalAccountGetMany(r.Context())
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, accounts)
}
