
DELETE FROM
	account.org_group_account

WHERE
	    org_id=$1
	AND group_id=$2;