
WITH deleted AS (
	DELETE FROM account.tenant WHERE id=$1 RETURNING id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;