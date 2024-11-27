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
/* === ORG QUERIES === */
/* ---------------------- */

//go:embed org/orgCreate.sql
var OrgCreate string

//go:embed org/orgGetMany.sql
var OrgGetMany string

//go:embed org/orgGetOne.sql
var OrgGetOne string

//go:embed org/orgUpdate.sql
var OrgUpdate string

//go:embed org/orgDelete.sql
var OrgDelete string

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
