package rest

import (
	"net/http"
	"switchcraft/cmd/rest/controllers/application"
	"switchcraft/cmd/rest/controllers/auth"
	featureflag "switchcraft/cmd/rest/controllers/featureFlag"
	"switchcraft/cmd/rest/controllers/org"
	orgaccount "switchcraft/cmd/rest/controllers/orgAccount"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/core"
	"switchcraft/types"
)

func addRoutes(logger *types.Logger, core *core.Core, router *http.ServeMux) {
	authController := auth.NewAuthController(logger, core)
	orgController := org.NewOrgController(logger, core)
	orgAccountController := orgaccount.NewOrgAccountController(logger, core)
	appController := application.NewAppController(logger, core)
	featFlagController := featureflag.NewFeatureFlagController(logger, core)

	authMiddleware := createAuthMiddleware(logger, core)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		restutils.Render(w, r, 200, map[string]any{
			"message": "Welcome to the SwitchCraft REST API",
		})
	})

	router.HandleFunc("POST /authn", authController.Login)

	/* === ORGANIZATION ROUTES === */
	router.HandleFunc("POST /org", authMiddleware(orgController.Create))
	router.HandleFunc("GET /org", authMiddleware(orgController.GetMany))
	router.HandleFunc("GET /org/{orgSlug}", authMiddleware(orgController.GetOne))
	router.HandleFunc("PUT /org/{orgSlug}", authMiddleware(orgController.Update))

	/* === ORG ACCOUNT ROUTES === */
	router.HandleFunc("POST /org/{orgSlug}/account", authMiddleware(orgAccountController.CreateOrgAccount))
	router.HandleFunc("GET /org/{orgSlug}/account", authMiddleware(orgAccountController.GetOrgAccounts))
	router.HandleFunc("GET /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.GetOneOrgAccount))
	router.HandleFunc("PUT /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.UpdateOrgAccount))
	router.HandleFunc("DELETE /org/{orgSlug}/account/{accountID}", authMiddleware(orgAccountController.DeleteOrgAccount))

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
