
WITH deleted AS (
	DELETE FROM account.account WHERE org_id IS NULL AND id=$1 RETURNING id
)
	
SELECT
	count(*)::bigint AS num_deleted
	
FROM
	deleted;