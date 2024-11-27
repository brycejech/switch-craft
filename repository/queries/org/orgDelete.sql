
WITH deleted AS (
	DELETE FROM account.org WHERE id=$1 RETURNING id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;