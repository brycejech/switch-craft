
UPDATE account.account

SET
	  is_instance_admin = $3
	, first_name = $4
	, last_name = $5
	, email = $6
	, username = $7
	, modified = (now() at time zone 'utc')
	, modified_by = $8

WHERE
	    ($1::bigint IS NULL OR org_id = $1::bigint)
	AND id = $2

RETURNING
	  org_id
	, id
	, uuid
	, is_instance_admin
	, first_name
	, last_name
	, email
	, username
	, password
	, created
	, created_by
	, modified
	, modified_by;