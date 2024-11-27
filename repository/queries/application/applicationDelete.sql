
WITH deleted AS (
	DELETE FROM application.application WHERE org_id=$1 AND id=$2 RETURNING id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;