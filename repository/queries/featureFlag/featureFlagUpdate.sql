
UPDATE application.feature_flag

SET
	  name = $4
	, is_enabled = $5
	, modified = (now() at time zone 'utc')
	, modified_by = $6

WHERE
	    org_id = $1
	AND application_id = $2
	AND id = $3

RETURNING
		org_id
	, application_id
	, id
	, uuid
	, name
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;