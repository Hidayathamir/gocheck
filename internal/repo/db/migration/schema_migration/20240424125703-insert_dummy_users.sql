-- +migrate Up
INSERT INTO
    "users" (
        "username",
        "full_name",
        "balance",
        "created_at",
        "updated_at"
    )
VALUES
    (
        'hidayat',
        'hidayat hamir',
        100000000,
        '2024-04-25 15:55:18.912',
        '2024-04-25 15:55:18.912'
    ) RETURNING "id";

INSERT INTO
    "users" (
        "username",
        "full_name",
        "balance",
        "created_at",
        "updated_at"
    )
VALUES
    (
        'hafiz',
        'hafiz arrahman',
        200000000,
        '2024-04-25 15:55:18.913',
        '2024-04-25 15:55:18.913'
    ) RETURNING "id";

INSERT INTO
    "users" (
        "username",
        "full_name",
        "balance",
        "created_at",
        "updated_at"
    )
VALUES
    (
        'aji',
        'aji hidayat',
        300000000,
        '2024-04-25 15:55:18.913',
        '2024-04-25 15:55:18.913'
    ) RETURNING "id";

-- +migrate Down