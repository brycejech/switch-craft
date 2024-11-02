
SELECT
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
	, modified_by

FROM
	account.account

WHERE tenant_id=$1;