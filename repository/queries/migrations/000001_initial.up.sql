BEGIN TRANSACTION;

CREATE SCHEMA account;

CREATE TABLE account.account (
	  id          bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid        uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, first_name  varchar(32)  NOT NULL
	, last_name   varchar(32)  NOT NULL
	, email       varchar(64)  NOT NULL
	, username    varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

CREATE TABLE account.tenant (
	  id    bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid  uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name  varchar(64)  NOT NULL UNIQUE
	, slug  varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

CREATE SCHEMA application;

CREATE TABLE application.application (
	  id         bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid       uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, tenant_id  bigint       NOT NULL REFERENCES account.tenant(id) ON DELETE CASCADE ON UPDATE CASCADE
	, name       varchar(64)  NOT NULL UNIQUE
	, slug       varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

CREATE TABLE application.feature_flag (
	  id              int          NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid            uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, application_id  bigint       NOT NULL REFERENCES application.application(id) ON DELETE CASCADE ON UPDATE CASCADE
	, name            varchar(64)  NOT NULL UNIQUE
	, slug            varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

END TRANSACTION;