
UPDATE account.account

SET
	  is_instance_admin = $2
	, first_name = $3
	, last_name = $4
	, email = $5
	, username = $6
	, modified = (now() at time zone 'utc')
	, modified_by = $7

WHERE
	    org_id IS NULL
	AND id = $1

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