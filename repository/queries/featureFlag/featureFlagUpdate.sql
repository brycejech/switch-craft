
UPDATE application.feature_flag

SET
	  name = $4
	, label = $5
	, description = $6
	, is_enabled = $7
	, modified = (now() at time zone 'utc')
	, modified_by = $8

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
	, label
	, description
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by;