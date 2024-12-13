package rest

import (
	"net/http"
	"switchcraft/cmd/rest/controllers/application"
	"switchcraft/cmd/rest/controllers/auth"
	"switchcraft/cmd/rest/controllers/featureflag"
	"switchcraft/cmd/rest/controllers/globalaccount"
	"switchcraft/cmd/rest/controllers/org"
	"switchcraft/cmd/rest/controllers/orgaccount"
	"switchcraft/cmd/rest/controllers/orggroup"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/core"
	"switchcraft/types"
)

func addRoutes(logger *types.Logger, core *core.Core, router *http.ServeMux) {
	var (
		authController          = auth.NewAuthController(logger, core)
		globalAccountController = globalaccount.NewGlobalAccountController(logger, core)
		orgController           = org.NewOrgController(logger, core)
		orgAccountController    = orgaccount.NewOrgAccountController(logger, core)
		orgGroupController      = orggroup.NewOrgGroupController(logger, core)
		appController           = application.NewAppController(logger, core)
		featFlagController      = featureflag.NewFeatureFlagController(logger, core)
	)

	authMiddleware := createAuthMiddleware(logger, core)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		restutils.Render(w, r, 200, map[string]any{
			"message": "Welcome to the SwitchCraft REST API",
		})
	})

	router.HandleFunc("POST /authn", authController.Login)

	/* === GLOBAL ACCOUNT ROUTES === */
	router.HandleFunc("POST /account", authMiddleware(globalAccountController.Create))
	router.HandleFunc("GET /account", authMiddleware(globalAccountController.GetMany))
	router.HandleFunc("GET /account/{accountID}", authMiddleware(globalAccountController.GetOne))
	router.HandleFunc("PUT /account/{accountID}", authMiddleware(globalAccountController.Update))
	router.HandleFunc("DELETE /account/{accountID}", authMiddleware(globalAccountController.Delete))

	/* === ORGANIZATION ROUTES === */
	router.HandleFunc("POST /org", authMiddleware(orgController.Create))
	router.HandleFunc("GET /org", authMiddleware(orgController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}", authMiddleware(orgController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}", authMiddleware(orgController.Update))

	/* === ORG ACCOUNT ROUTES === */
	router.HandleFunc("POST /org/{orgSlug}/account", authMiddleware(orgAccountController.Create))
	router.HandleFunc("GET /org/{orgSlug}/account", authMiddleware(orgAccountController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.Update))
	router.HandleFunc("DELETE /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.Delete))

	/* === ORG GROUP ROUTES === */
	router.HandleFunc("POST /org/{orgSlug}/group", authMiddleware(orgGroupController.Create))
	router.HandleFunc("GET /org/{orgSlug}/group", authMiddleware(orgGroupController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}/group/{groupID}", authMiddleware(orgGroupController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}/group/{groupID}", authMiddleware(orgGroupController.Update))
	router.HandleFunc("DELETE /org/{orgSlug}/group/{groupID}", authMiddleware(orgGroupController.Delete))

	/* === APPLICATION ROUTES === */
	router.HandleFunc("POST /org/{orgSlug}/app", authMiddleware(appController.Create))
	router.HandleFunc("GET /org/{orgSlug}/app", authMiddleware(appController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}/app/{appSlug}", authMiddleware(appController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}/app/{appSlug}", authMiddleware(appController.Update))
	router.HandleFunc("DELETE /org/{orgSlug}/app/{appSlug}", authMiddleware(appController.Delete))

	/* === FEATURE FLAG ROUTES === */
	router.HandleFunc("POST /org/{orgSlug}/app/{appSlug}/flag", authMiddleware(featFlagController.Create))
	router.HandleFunc("GET /org/{orgSlug}/app/{appSlug}/flag", authMiddleware(featFlagController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}/app/{appSlug}/flag/{flagID}", authMiddleware(featFlagController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}/app/{appSlug}/flag/{flagID}", authMiddleware(featFlagController.Update))
	router.HandleFunc("DELETE /org/{orgSlug}/app/{appSlug}/flag/{flagID}", authMiddleware(featFlagController.Delete))
}
