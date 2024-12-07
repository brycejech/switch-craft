package queries

import (
	"embed"
)

/* --------------------------- */
/* === ORG ACCOUNT QUERIES === */
/* --------------------------- */

//go:embed orgAccount/orgAccountCreate.sql
var OrgAccountCreate string

//go:embed orgAccount/orgAccountGetMany.sql
var OrgAccountGetMany string

//go:embed orgAccount/orgAccountGetOne.sql
var OrgAccountGetOne string

//go:embed orgAccount/orgAccountUpdate.sql
var OrgAccountUpdate string

//go:embed orgAccount/orgAccountDelete.sql
var OrgAccountDelete string

//go:embed orgAccount/accountGetByUsername.sql
var AccountGetByUsername string

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
