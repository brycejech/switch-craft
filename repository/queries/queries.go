package queries

import (
	"embed"
)

/* === ACCOUNT QUERIES === */

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

/* === TENANT QUERIES === */

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

/* === MIGRATION QUERIES === */

//go:embed migrations
var Migrations embed.FS
