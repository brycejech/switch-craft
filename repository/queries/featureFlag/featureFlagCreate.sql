
INSERT INTO application.feature_flag (
	  tenant_id
	, application_id
	, name
	, is_enabled
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
	, $5
)

RETURNING
		tenant_id
	, application_id
	, id
	, uuid
	, name
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;