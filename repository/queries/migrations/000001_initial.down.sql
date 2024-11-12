BEGIN TRANSACTION;

DROP TABLE application.feature_flag;
DROP TABLE application.application;
DROP SCHEMA application;

ALTER TABLE account.tenant
	  DROP COLUMN owner
	, DROP COLUMN created_by
	, DROP COLUMN modified_by;

DROP TABLE account.account;
DROP TABLE account.tenant;
DROP SCHEMA account;

END TRANSACTION;