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

/* === MIGRATION QUERIES === */

//go:embed migrations
var Migrations embed.FS
