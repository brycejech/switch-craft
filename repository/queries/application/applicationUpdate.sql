
UPDATE application.application

SET
	  name = $3
	, slug = $4
	, modified = (now() at time zone 'utc')
	, modified_by = $5

WHERE
	    org_id = $1
	AND id = $2

RETURNING
	  org_id
	, id
	, uuid
	, name
	, slug
	, created
	, created_by
	, modified
	, modified_by;