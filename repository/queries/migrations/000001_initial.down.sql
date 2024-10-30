BEGIN TRANSACTION;

DROP TABLE application.feature_flag;
DROP TABLE application.application;
DROP SCHEMA application;

DROP TABLE account.tenant;
DROP TABLE account.account;
DROP SCHEMA account;

END TRANSACTION;