BEGIN TRANSACTION;

CREATE SCHEMA account;

CREATE TABLE account.tenant (
	  id     bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid   uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name   varchar(64)  NOT NULL UNIQUE
	, slug   varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, modified     timestamp with time zone
);

CREATE TABLE account.account (
	  tenant_id   bigint        REFERENCES account.tenant(id) ON DELETE CASCADE ON UPDATE CASCADE
	, id          bigint        NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid        uuid          NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, first_name  varchar(32)   NOT NULL
	, last_name   varchar(32)   NOT NULL
	, email       varchar(64)   NOT NULL
	, username    varchar(64)   NOT NULL UNIQUE
	, password    varchar(128)

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

ALTER TABLE account.tenant
	  ADD COLUMN owner        bigint  REFERENCES account.account(id)
	, ADD COLUMN created_by   bigint  REFERENCES account.account(id)
	, ADD COLUMN modified_by  bigint  REFERENCES account.account(id);

CREATE SCHEMA application;

CREATE TABLE application.application (
	  tenant_id  bigint       NOT NULL REFERENCES account.tenant(id) ON DELETE CASCADE ON UPDATE CASCADE
	
	, id         bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid       uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name       varchar(64)  NOT NULL UNIQUE
	, slug       varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

CREATE TABLE application.feature_flag (
	  tenant_id       bigint       NOT NULL REFERENCES account.tenant(id) ON DELETE CASCADE ON UPDATE CASCADE
	, application_id  bigint       NOT NULL REFERENCES application.application(id) ON DELETE CASCADE ON UPDATE CASCADE
	
	, id              int          NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid            uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name            varchar(64)  NOT NULL UNIQUE
	, slug            varchar(64)  NOT NULL UNIQUE
	, is_enabled      boolean      NOT NULL DEFAULT FALSE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

END TRANSACTION;