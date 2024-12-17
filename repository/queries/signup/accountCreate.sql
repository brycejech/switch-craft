

WITH new_account AS (
	INSERT INTO account.account(
			is_instance_admin
		, first_name
		, last_name
		, email
		, username
		, password
		, created_by
	)

	VALUES (
			false
		, $1
		, $2
		, $3
		, $4
		, $5
		, NULL
	)

	RETURNING id
)

UPDATE
	account.account AS A

SET
	A.created_by = new_account.id

FROM
	new_account

WHERE
	new_account.id = A.id

RETURNING
	  A.org_id
	, A.id
	, A.uuid
	, A.is_instance_admin
	, A.first_name
	, A.last_name
	, A.email
	, A.username
	, A.password
	, A.created
	, A.created_by
	, A.modified
	, A.modified_by;




