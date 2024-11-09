
SELECT
	  tenant_id
	, id
	, uuid
	, name
	, slug
	, created
	, created_by
	, modified
	, modified_by

FROM
	application.application

WHERE
	    tenant_id=$1
	AND ($2::bigint IS NULL    OR id=$2::bigint)
	AND (COALESCE($3, '') = '' OR uuid=$3::uuid)
	AND (COALESCE($4, '') = '' OR slug=$4::text)