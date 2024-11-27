
INSERT INTO application.application (
	  org_id
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
	  org_id
	, id
	, uuid
	, name
	, slug
	, created
	, created_by
	, modified
	, modified_by;
