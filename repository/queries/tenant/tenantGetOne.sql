
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
	account.tenant

WHERE
	id = $1;