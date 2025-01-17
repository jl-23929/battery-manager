CREATE TABLE IF NOT EXISTS battery
(
    battery_id        TEXT NOT NULL,
    battery_name      TEXT NOT NULL,
    brand             TEXT NOT NULL,
    capacity          INTEGER NOT NULL DEFAULT 0,
    purchase_date     TEXT NOT NULL DEFAULT '',
    purchase_price    TEXT NOT NULL DEFAULT '',
    purchase_currency TEXT NOT NULL DEFAULT '',
    competition_class TEXT NOT NULL DEFAULT '',
    status            TEXT NOT NULL DEFAULT '',
    created           TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    updated           TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    PRIMARY KEY (battery_id)
    ) STRICT;
