
WITH deleted AS (
	DELETE FROM
		account.org_group_account
		
	WHERE
		    org_id=$1
		AND group_id=$2
		AND account_id=$3
	
	RETURNING id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;