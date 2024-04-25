-- +migrate Up
CREATE TABLE "transactions" (
    "id" bigserial,
    "sender_id" bigint,
    "recipient_id" bigint,
    "amount" bigint,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- +migrate Down