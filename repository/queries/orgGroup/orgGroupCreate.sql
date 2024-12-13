
INSERT INTO account.org_group (
	  org_id
	, name
	, description
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
	, description
	, created
	, created_by
	, modified
	, modified_by;