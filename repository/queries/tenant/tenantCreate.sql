
INSERT INTO account.tenant (
	  name
	, slug
	, owner
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
)

RETURNING
	  id
	, uuid
	, name
	, slug
	, owner
	, created
	, created_by
	, modified
	, modified_by;