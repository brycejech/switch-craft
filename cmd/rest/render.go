package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"switchcraft/types"
)

type HTTPStatusCode int

var logger = types.NewLogger(types.LogLevelInfo)

func render(w http.ResponseWriter, r *http.Request, status HTTPStatusCode, data any) {
	trace, ok := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
	if !ok {
		fmt.Println("rest.render invalid operation context")
	}

	if s, ok := data.(string); ok {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(int(status))
		w.Write([]byte(s))

		logger.Info(trace, "Request end", map[string]any{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": status,
		})

		return
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))

		logger.Info(trace, "Request end", map[string]any{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": http.StatusInternalServerError,
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write(bytes)

	logger.Info(trace, "Request end", map[string]any{
		"method": r.Method,
		"path":   r.URL.Path,
		"status": status,
	})
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusBadRequest, "Bad request")
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusUnauthorized, "Unauthorized")
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusInternalServerError, "Internal server error")
}

func jsonParseError(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusBadRequest, "JSON parse error")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusNotFound, "Not found")
}
