package queries

import (
	"embed"
)

/* ----------------------- */
/* === ACCOUNT QUERIES === */
/* ----------------------- */

//go:embed account/accountCreate.sql
var AccountCreate string

//go:embed account/accountGetMany.sql
var AccountGetMany string

//go:embed account/accountGetOne.sql
var AccountGetOne string

//go:embed account/accountUpdate.sql
var AccountUpdate string

//go:embed account/accountDelete.sql
var AccountDelete string

/* ---------------------- */
/* === TENANT QUERIES === */
/* ---------------------- */

//go:embed tenant/tenantCreate.sql
var TenantCreate string

//go:embed tenant/tenantGetMany.sql
var TenantGetMany string

//go:embed tenant/tenantGetOne.sql
var TenantGetOne string

//go:embed tenant/tenantUpdate.sql
var TenantUpdate string

//go:embed tenant/tenantDelete.sql
var TenantDelete string

/* --------------------------- */
/* === APPLICATION QUERIES === */
/* --------------------------- */

//go:embed application/applicationCreate.sql
var AppCreate string

//go:embed application/applicationGetMany.sql
var AppGetMany string

//go:embed application/applicationGetOne.sql
var AppGetOne string

//go:embed application/applicationUpdate.sql
var AppUpdate string

//go:embed application/applicationDelete.sql
var AppDelete string

/* ---------------------------- */
/* === FEATURE FLAG QUERIES === */
/* ---------------------------- */

//go:embed featureFlag/featureFlagCreate.sql
var FeatureFlagCreate string

//go:embed featureFlag/featureFlagGetMany.sql
var FeatureFlagGetMany string

//go:embed featureFlag/featureFlagGetOne.sql
var FeatureFlagGetOne string

//go:embed featureFlag/featureFlagUpdate.sql
var FeatureFlagUpdate string

//go:embed featureFlag/featureFlagDelete.sql
var FeatureFlagDelete string

/* ------------------------- */
/* === MIGRATION QUERIES === */
/* ------------------------- */

//go:embed migrations
var Migrations embed.FS
