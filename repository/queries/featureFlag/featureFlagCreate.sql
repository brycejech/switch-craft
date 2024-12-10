
INSERT INTO application.feature_flag (
	  org_id
	, application_id
	, name
	, label
	, description
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
	, $7
)

RETURNING
		org_id
	, application_id
	, id
	, uuid
	, name
	, label
	, description
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;