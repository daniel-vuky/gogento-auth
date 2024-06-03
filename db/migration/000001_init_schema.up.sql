CREATE TYPE "gender" AS ENUM (
    '1',
    '2',
    '3'
);

CREATE TABLE "customers" (
    "customer_id" bigserial PRIMARY KEY,
    "email" varchar(255) UNIQUE NOT NULL,
    "firstname" varchar(32) NOT NULL,
    "lastname" varchar(32) NOT NULL,
    "gender" gender,
    "dob" timestamptz,
    "hashed_password" varchar NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "admin_users" (
    "admin_id" bigserial PRIMARY KEY,
    "role_id" bigint NOT NULL,
    "email" varchar(255) UNIQUE NOT NULL,
    "firstname" varchar(32) NOT NULL,
    "lastname" varchar(32) NOT NULL,
    "hashed_password" varchar NOT NULL,
    "is_active" bool NOT NULL DEFAULT true,
    "lock_expires" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "authorization_roles" (
    "role_id" bigserial PRIMARY KEY,
    "role_name" varchar(255) NOT NULL,
    "description" varchar(255),
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

INSERT INTO authorization_roles ("role_name") VALUES ('Administrator');

CREATE TABLE "authorization_rules" (
    "rule_id" bigserial PRIMARY KEY,
    "role_id" bigint NOT NULL,
    "is_administrator" bool NOT NULL DEFAULT true,
    "permission_code" varchar(128) NOT NULL,
    "is_allowed" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

INSERT INTO authorization_rules
    ("role_id", "is_administrator", "permission_code", "is_allowed")
    VALUES (1, true, 'all', true);

CREATE TABLE "refresh_tokens" (
    "refresh_token_id" bigserial PRIMARY KEY,
    "customer_id" bigint NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" bool NOT NULL DEFAULT false,
    "expired_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE UNIQUE INDEX ON "authorization_rules" ("role_id", "permission_code");

ALTER TABLE customers
    ADD CONSTRAINT CUSTOMERS_EMAIL_NOT_EMPTY CHECK (email <> '');

ALTER TABLE "admin_users"
    ADD FOREIGN KEY ("role_id")
        REFERENCES "authorization_roles" ("role_id")
        ON DELETE SET NULL ON UPDATE NO ACTION;

ALTER TABLE "authorization_rules"
    ADD FOREIGN KEY ("role_id")
        REFERENCES "authorization_roles" ("role_id")
        ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "refresh_tokens"
    ADD FOREIGN KEY ("customer_id")
        REFERENCES "customers" ("customer_id")
        ON DELETE CASCADE ON UPDATE NO ACTION;


