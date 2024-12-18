
UPDATE account.account

SET org_id=$1

WHERE id=$2

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