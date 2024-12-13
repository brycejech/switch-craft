
SELECT
	  org_id
	, id
	, uuid
	, name
	, description
	, created
	, created_by
	, modified
	, modified_by

FROM
	account.org_group

WHERE
	    org_id = $1
	AND ($2::bigint IS NULL    OR id=$2::bigint)
	AND (COALESCE($3, '') = '' OR uuid=$3::uuid);