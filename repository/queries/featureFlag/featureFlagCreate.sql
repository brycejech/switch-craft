
INSERT INTO application.feature_flag (
	  tenant_id
	, application_id
	, name
	, slug
	, is_enabled
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
	, $5
	, $6
)

RETURNING
		tenant_id
	, application_id
	, id
	, uuid
	, name
	, slug
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;