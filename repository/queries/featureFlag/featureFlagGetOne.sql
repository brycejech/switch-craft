
SELECT
		org_id
	, application_id
	, id
	, uuid
	, name
	, label
	, description
	, is_enabled
	, created
	, created_by
	, modified
	, modified_by

FROM
	application.feature_flag

WHERE
	    org_id=$1
	AND application_id=$2
	AND ($3::bigint IS NULL    OR id=$3::bigint)
	AND (COALESCE($4, '') = '' OR uuid=$4::uuid)
	AND (COALESCE($5, '') = '' OR name=$5::text);