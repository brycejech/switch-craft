
UPDATE account.account

SET
	  first_name = $2
	, last_name = $3
	, email = $4
	, username = $5
	, modified = (now() at time zone 'utc')
	, modified_by = $6

WHERE
	id = $1

RETURNING
	  id
	, uuid
	, first_name
	, last_name
	, email
	, username
	, created
	, created_by
	, modified
	, modified_by;