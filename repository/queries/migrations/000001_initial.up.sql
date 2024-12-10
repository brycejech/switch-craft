BEGIN TRANSACTION;

CREATE SCHEMA account;

CREATE TABLE account.org (
	  id     bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid   uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name   varchar(64)  NOT NULL UNIQUE
	, slug   varchar(64)  NOT NULL UNIQUE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, modified     timestamp with time zone
);

CREATE TABLE account.account (
	  org_id             bigint        REFERENCES account.org(id) ON DELETE CASCADE ON UPDATE CASCADE

	, id                 bigint        NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid               uuid          NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, is_instance_admin  bool          NOT NULL DEFAULT FALSE
	, first_name         varchar(32)   NOT NULL
	, last_name          varchar(32)   NOT NULL
	, email              varchar(64)   NOT NULL
	, username           varchar(64)   NOT NULL UNIQUE
	, password           varchar(128)

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)
);

ALTER TABLE account.org
	  ADD COLUMN owner        bigint  REFERENCES account.account(id)
	, ADD COLUMN created_by   bigint  REFERENCES account.account(id)
	, ADD COLUMN modified_by  bigint  REFERENCES account.account(id);

CREATE SCHEMA application;

CREATE TABLE application.application (
	  org_id  bigint       NOT NULL REFERENCES account.org(id) ON DELETE CASCADE ON UPDATE CASCADE

	, id      bigint       NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid    uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name    varchar(64)  NOT NULL
	, slug    varchar(64)  NOT NULL

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)

	, UNIQUE (org_id, name)
	, UNIQUE (org_id, slug)
);

CREATE TABLE application.feature_flag (
	  org_id          bigint       NOT NULL REFERENCES account.org(id) ON DELETE CASCADE ON UPDATE CASCADE
	, application_id  bigint       NOT NULL REFERENCES application.application(id) ON DELETE CASCADE ON UPDATE CASCADE

	, id              int          NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
	, uuid            uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
	, name            varchar(64)  NOT NULL
	, label           varchar(64)  NOT NULL
	, description     text
	, is_enabled      boolean      NOT NULL DEFAULT FALSE

	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
	, created_by   bigint                    REFERENCES account.account(id)
	, modified     timestamp with time zone
	, modified_by  bigint                    REFERENCES account.account(id)

	, UNIQUE (application_id, name)
);

-- CREATE TABLE account.org_group (
-- 	  org_id bigint NOT NULL REFERENCES account.org(id) ON DELETE CASCADE ON UPDATE CASCADE

-- 	, id              int          NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
-- 	, uuid            uuid         NOT NULL UNIQUE DEFAULT gen_random_uuid()
-- 	, name            varchar(64)  NOT NULL

-- 	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
-- 	, created_by   bigint                    REFERENCES account.account(id)
-- 	, modified     timestamp with time zone
-- 	, modified_by  bigint                    REFERENCES account.account(id)

-- 	, UNIQUE (org_id, name)
-- );

-- CREATE TABLE account.org_group_account (
-- 	  org_id      bigint  NOT NULL REFERENCES account.org(id) ON DELETE CASCADE ON UPDATE CASCADE
-- 	, group_id    bigint  NOT NULL REFERENCES account.org_group(id) ON DELETE CASCADE ON UPDATE CASCADE

-- 	, id          bigint  NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY
-- 	, account_id  bigint  NOT NULL REFERENCES account.account(id) ON DELETE CASCADE ON UPDATE CASCADE

-- 	, created      timestamp with time zone  NOT NULL DEFAULT (now() at time zone 'utc')
-- 	, created_by   bigint                    REFERENCES account.account(id)
-- 	, modified     timestamp with time zone
-- 	, modified_by  bigint                    REFERENCES account.account(id)

-- 	UNIQUE (group_id, account_id)
-- );

END TRANSACTION;