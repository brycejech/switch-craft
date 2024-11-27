package rest

import (
	"net/http"
	"switchcraft/core"
	"switchcraft/types"
)

func addRoutes(logger *types.Logger, core *core.Core, router *http.ServeMux) {
	authController := newAuthController(logger, core)
	orgController := newOrgController(logger, core)

	authMiddleware := createAuthMiddleware(logger, core)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, 200, map[string]any{
			"message": "Welcome to the SwitchCraft REST API",
		})
	})

	router.HandleFunc("POST /authn", authController.Login)

	router.HandleFunc("POST /org", authMiddleware(orgController.Create))
	router.HandleFunc("GET /org", authMiddleware(orgController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}", authMiddleware(orgController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}", authMiddleware(orgController.Update))
}
