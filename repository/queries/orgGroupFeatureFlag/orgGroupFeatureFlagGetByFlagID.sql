
SELECT
	  org_id
	, group_id
	, application_id
	, flag_id
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by

FROM
	application.org_group_feature_flag

WHERE
	    org_id = $1
	AND application_id = $2
	AND flag_id = $3;