
WITH deleted AS
	(
		DELETE FROM account.account WHERE ($1::bigint IS NULL OR tenant_id=$1::bigint) AND id=$2 RETURNING *
	)
	
SELECT
	count(*)::bigint AS num_deleted
	
FROM
	deleted;