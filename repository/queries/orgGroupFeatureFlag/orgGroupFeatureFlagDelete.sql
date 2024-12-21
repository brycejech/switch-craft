
WITH deleted AS (
	DELETE FROM
		application.org_group_feature_flag

	WHERE
		    org_id = $1
		AND group_id = $2
		AND application_id = $3
		AND flag_id = $4

	RETURNING group_id
)

SELECT
	count(*) AS num_deleted

FROM
	deleted;