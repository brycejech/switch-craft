package rest

import (
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

func addRoutes(logger *types.Logger, core *core.Core, router *http.ServeMux) {
	authController := newAuthController(logger, core)
	tenantController := newTenantController(logger, core)

	authMiddleware := createAuthMiddleware(logger, core)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, 200, map[string]any{
			"message": "Welcome to the SwitchCraft REST API",
		})
	})

	router.HandleFunc("POST /authn", authController.Login)

	router.HandleFunc("POST /tenant", authMiddleware(tenantController.Create))
	router.HandleFunc("GET /tenant", authMiddleware(tenantController.GetMany))
	router.HandleFunc("GET /tenant/{tenantSlug}", authMiddleware(tenantController.GetOne))
	router.HandleFunc("PUT /tenant/{tenantSlug}", authMiddleware(tenantController.Update))
}
