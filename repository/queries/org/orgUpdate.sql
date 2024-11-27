
UPDATE
	account.org

SET
	  name = $2
	, slug = $3
	, owner = $4
	, modified = (now() at time zone 'utc')
	, modified_by = $5

WHERE
	id = $1

RETURNING
	  id
	, uuid
	, name
	, slug
	, owner
	, created
	, created_by
	, modified
	, modified_by;