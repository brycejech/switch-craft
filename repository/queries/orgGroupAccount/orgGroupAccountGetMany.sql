
SELECT
	  a.org_id
	, a.id
	, a.uuid
	, a.is_instance_admin
	, a.first_name
	, a.last_name
	, a.email
	, a.username
	, a.password
	, a.created
	, a.created_by
	, a.modified
	, a.modified_by

FROM
	account.org_group_account AS oga

INNER JOIN account.account AS a
	ON
		(
					a.id = oga.account_id
			AND a.org_id = oga.org_id
		)

WHERE
	    oga.org_id=$1
	AND oga.group_id=$2;