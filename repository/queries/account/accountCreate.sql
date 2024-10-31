
INSERT INTO account.account(
	  first_name
	, last_name
	, email
	, username
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
	, $5
)

RETURNING
	  id
	, uuid
	, first_name
	, last_name
	, email
	, username
	, created
	, created_by
	, modified
	, modified_by;
