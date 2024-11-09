
INSERT INTO application.application (
	  tenant_id
	, name
	, slug
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
)

RETURNING
	  tenant_id
	, id
	, uuid
	, name
	, slug
	, created
	, created_by
	, modified
	, modified_by;
