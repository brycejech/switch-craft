
UPDATE
	account.org_group

SET
	  name = $3
	, description = $4
	, modified = (now() at time zone 'utc')
	, modified_by = $5

WHERE
	    org_id = $1
	AND id=$2

RETURNING
	  org_id
	, id
	, uuid
	, name
	, description
	, created
	, created_by
	, modified
	, modified_by;