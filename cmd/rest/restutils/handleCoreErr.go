package restutils

import (
	"errors"
	"net/http"
	"switchcraft/types"
)

func HandleCoreErr(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		// Expected err, got nil
		// Better to report to caller an unknown error rather than success
		InternalServerError(w, r)
		return
	}

	if errors.Is(err, types.ErrNotFound) {
		NotFound(w, r)
	} else if errors.Is(err, types.ErrItemExists) {
		BadRequest(w, r)
	} else if errors.Is(err, types.ErrLinkedItemNotFound) {
		BadRequest(w, r)
	} else {
		InternalServerError(w, r)
	}
}
