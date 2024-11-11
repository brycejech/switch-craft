
WITH deleted AS (
	DELETE FROM
		application.feature_flag

	WHERE
		    tenant_id = $1
		AND application_id = $2
		AND id = $3
	
	RETURNING
		id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;