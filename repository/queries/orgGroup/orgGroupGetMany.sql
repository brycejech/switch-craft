
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
	org_id = $1;