package restutils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
