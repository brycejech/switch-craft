package queries

import (
	"embed"
)

/* === MIGRATION QUERIES === */

//go:embed migrations
var Migrations embed.FS
