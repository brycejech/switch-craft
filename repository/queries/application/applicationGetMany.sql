
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
	tenant_id = $1;