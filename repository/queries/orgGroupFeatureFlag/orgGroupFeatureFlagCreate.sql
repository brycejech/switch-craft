
INSERT INTO application.org_group_feature_flag (
	  org_id
	, group_id
	, application_id
	, flag_id
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
	  org_id
	, group_id
	, application_id
	, flag_id
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;