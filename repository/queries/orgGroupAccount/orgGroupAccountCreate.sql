
INSERT INTO account.org_group_account (
	  org_id
	, group_id
	, account_id
	, created_by
)

VALUES (
	  $1
	, $2
	, $3
	, $4
)

RETURNING
	  org_id
	, group_id
	, id
	, account_id
	, created
	, created_by;