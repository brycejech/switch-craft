
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

	    org_id IS NULL
	AND ($1::bigint IS NULL    OR id=$1::bigint)
	AND (COALESCE($2, '') = '' OR uuid=$2::uuid)
	AND (COALESCE($3, '') = '' OR username=$3::text)

LIMIT 1;