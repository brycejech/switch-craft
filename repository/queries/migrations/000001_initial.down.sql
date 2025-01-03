BEGIN TRANSACTION;

DROP TABLE application.org_group_feature_flag;

DROP TABLE account.org_group_account;
DROP TABLE account.org_group;

DROP TABLE application.feature_flag;
DROP TABLE application.application;
DROP SCHEMA application;

ALTER TABLE account.org
	  DROP COLUMN owner
	, DROP COLUMN created_by
	, DROP COLUMN modified_by;

DROP TABLE account.account;
DROP TABLE account.org;
DROP SCHEMA account;

END TRANSACTION;