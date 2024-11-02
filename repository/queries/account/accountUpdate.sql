
UPDATE account.account

SET
	  first_name = $3
	, last_name = $4
	, email = $5
	, username = $6
	, modified = (now() at time zone 'utc')
	, modified_by = $7

WHERE
	    ($1::bigint IS NULL OR tenant_id = $1::bigint)
	AND id = $2

RETURNING
	  tenant_id
	, id
	, uuid
	, first_name
	, last_name
	, email
	, username
	, password
	, created
	, created_by
	, modified
	, modified_by;