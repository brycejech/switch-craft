
WITH deleted AS (
	DELETE FROM account.account WHERE org_id=$1 AND id=$2 RETURNING id
)
	
SELECT
	count(*)::bigint AS num_deleted
	
FROM
	deleted;