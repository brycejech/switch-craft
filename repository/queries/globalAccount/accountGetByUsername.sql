
SELECT
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
	, modified_by

FROM
	account.account

WHERE
	username=$1

LIMIT 1;