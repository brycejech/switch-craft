
INSERT INTO account.account (
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
	, currval(
	    pg_get_serial_sequence('account.account','id')
		)
)

RETURNING
	  org_id
	, id
	, uuid
	, is_instance_admin
	, first_name
	, last_name
	, email
	, username
	, password
	, created
	, created_by
	, modified
	, modified_by;