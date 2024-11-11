
SELECT
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
	, modified_by

FROM
	application.feature_flag

WHERE
	    tenant_id = $1
	AND application_id = $2;