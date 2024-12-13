
WITH deleted AS (
	DELETE FROM account.org_group WHERE org_id=$1 AND id=$2 RETURNING id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;