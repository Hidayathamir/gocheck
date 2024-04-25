-- +migrate Up
CREATE TABLE "users" (
    "id" bigserial,
    "username" text,
    "full_name" text,
    "balance" bigint,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- +migrate Down