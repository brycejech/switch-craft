package globalaccount

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type updateGlobalAccountArgs struct {
	ID              int64  `json:"id"`
	IsInstanceAdmin bool   `json:"isInstanceAdmin"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Username        string `json:"username"`
}

func (c *globalAccountController) Update(w http.ResponseWriter, r *http.Request) {
	accountIDStr := r.PathValue("accountID")
	if accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &updateGlobalAccountArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
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

	account, err := c.core.GlobalAccountGetOne(r.Context(),
		c.core.NewGlobalAccountGetOneArgs(&accountID, nil, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)

	if account.ID != body.ID {
		restutils.BadRequest(w, r)
		c.logger.Warn(tracer, "Global account update ID mismatch detected", map[string]any{
			"user":            tracer.AuthAccount.Username,
			"requestBody":     body,
			"existingAccount": account,
		})
		return
	}

	updatedAccount, err := c.core.GlobalAccountUpdate(r.Context(),
		c.core.NewGlobalAccountUpdateArgs(
			account.ID,
			body.IsInstanceAdmin,
			body.FirstName,
			body.LastName,
			body.Email,
			body.Username,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, updatedAccount)
}
