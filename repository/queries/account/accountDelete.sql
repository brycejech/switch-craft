
WITH deleted AS
	(
		DELETE FROM
			account.account
			
		WHERE
					tenant_id=$1
			AND id=$2
		
		RETURNING *
	)
	
SELECT
	count(*)::bigint AS num_deleted
	
FROM
	deleted;