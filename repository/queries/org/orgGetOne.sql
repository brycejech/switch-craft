
SELECT
	  id
	, uuid
	, name
	, slug
	, owner
	, created
	, created_by
	, modified
	, modified_by

FROM
	account.org

WHERE
	    (COALESCE($1, '') = '' OR id=$1::bigint)
	AND (COALESCE($2, '') = '' OR uuid=$2::uuid)
	AND (COALESCE($3, '') = '' OR slug=$3::text)