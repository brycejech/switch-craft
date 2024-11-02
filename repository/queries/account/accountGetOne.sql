
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

WHERE

	    ($1::bigint IS NULL    OR tenant_id=$1::bigint)
	AND ($2::bigint IS NULL    OR id=$2::bigint)
	AND (COALESCE($3, '') = '' OR uuid=$3::uuid)
	AND (COALESCE($4, '') = '' OR username=$4::text)

LIMIT 1;