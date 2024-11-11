
UPDATE application.feature_flag

SET
	  name = $4
	, slug = $5
	, is_enabled = $6
	, modified = (now() at time zone 'utc')
	, modified_by = $7

WHERE
	    tenant_id = $1
	AND application_id = $2
	AND id = $3

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