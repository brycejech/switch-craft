
SELECT
	  org_id
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
	org_id = $1;