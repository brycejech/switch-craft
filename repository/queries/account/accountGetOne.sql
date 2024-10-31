
SELECT
	  id
	, uuid
	, first_name
	, last_name
	, email
	, username
	, created
	, created_by
	, modified
	, modified_by

FROM
	account.account

WHERE
      ($1::bigint IS NULL    OR id=$1::bigint)
  AND (COALESCE($2, '') = '' OR uuid=$2::uuid)
  AND (COALESCE($3, '') = '' OR username=$3::text)

LIMIT 1;