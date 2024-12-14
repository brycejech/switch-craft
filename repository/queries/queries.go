package queries

import (
	"embed"
)

/* ------------------------------ */
/* === GLOBAL ACCOUNT QUERIES === */
/* ------------------------------ */

//go:embed globalAccount/globalAccountCreate.sql
var GlobalAccountCreate string

//go:embed globalAccount/globalAccountGetMany.sql
var GlobalAccountGetMany string

//go:embed globalAccount/globalAccountGetOne.sql
var GlobalAccountGetOne string

//go:embed globalAccount/globalAccountUpdate.sql
var GlobalAccountUpdate string

//go:embed globalAccount/globalAccountDelete.sql
var GlobalAccountDelete string

//go:embed globalAccount/accountGetByUsername.sql
var AccountGetByUsername string

/* --------------------------- */
/* === ORG ACCOUNT QUERIES === */
/* --------------------------- */

//go:embed orgAccount/orgAccountCreate.sql
var OrgAccountCreate string

//go:embed orgAccount/orgAccountGetMany.sql
var OrgAccountGetMany string

//go:embed orgAccount/orgAccountGetManyById.sql
var OrgAccountGetManyByID string

//go:embed orgAccount/orgAccountGetOne.sql
var OrgAccountGetOne string

//go:embed orgAccount/orgAccountUpdate.sql
var OrgAccountUpdate string

//go:embed orgAccount/orgAccountDelete.sql
var OrgAccountDelete string

/* ------------------------- */
/* === ORG GROUP QUERIES === */
/* ------------------------- */

//go:embed orgGroup/orgGroupCreate.sql
var OrgGroupCreate string

//go:embed orgGroup/orgGroupGetMany.sql
var OrgGroupGetMany string

//go:embed orgGroup/orgGroupGetOne.sql
var OrgGroupGetOne string

//go:embed orgGroup/orgGroupUpdate.sql
var OrgGroupUpdate string

//go:embed orgGroup/orgGroupDelete.sql
var OrgGroupDelete string

/* ------------------------- */
/* === ORG GROUP QUERIES === */
/* ------------------------- */

//go:embed orgGroupAccount/orgGroupAccountCreate.sql
var OrgGroupAccountCreate string

//go:embed orgGroupAccount/orgGroupAccountGetMany.sql
var OrgGroupAccountGetMany string

//go:embed orgGroupAccount/orgGroupAccountDeleteOne.sql
var OrgGroupAccountDeleteOne string

//go:embed orgGroupAccount/orgGroupAccountDeleteAll.sql
var OrgGroupAccountDeleteAll string

/* ------------------- */
/* === ORG QUERIES === */
/* ------------------- */

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
