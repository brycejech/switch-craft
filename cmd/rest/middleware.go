package rest

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/core"
	"switchcraft/types"
	"time"
)

func trace(logger *types.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			traceId     = r.Header.Get("X-Trace-Id")
			startTime   = time.Now()
			authAccount = types.Account{}
		)

		opCtx := types.NewOperationCtx(r.Context(), traceId, startTime, authAccount)
		tracer, _ := opCtx.Value(types.CtxOperationTracer).(types.OperationTracer)

		logger.Info(tracer, "Request start", map[string]any{
			"path":   r.URL.Path,
			"method": r.Method,
		})

		next.ServeHTTP(w, r.WithContext(opCtx))
	})
}

var tokenRegexp = regexp.MustCompile(`^[Bb]earer\s`)

func createAuthMiddleware(logger *types.Logger, core *core.Core) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			tracer, ok := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)
			if !ok {
				restutils.InternalServerError(w, r)
				return
			}

			token := strings.Trim(
				tokenRegexp.ReplaceAllString(r.Header.Get("Authorization"), ""),
				" ",
			)

			if token == "" {
				logger.Error(tracer, "authorization token not provided", nil)
				restutils.BadRequest(w, r)
				return
			}

			account, err := core.AuthValidateJWT(token)
			if err != nil {
				fmt.Println(err)
				restutils.BadRequest(w, r)
				return
			}
			if account == nil {
				restutils.InternalServerError(w, r)
				return
			}

			tracer.AuthAccount = *account

			ctx := context.WithValue(r.Context(), types.CtxOperationTracer, tracer)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
