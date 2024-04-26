-- +migrate Up
CREATE TABLE "user_credentials" (
    "id" bigserial,
    "user_id" bigint,
    "password" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- +migrate Down