PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE battery_history
(
    id         TEXT NOT NULL,
    battery_id TEXT NOT NULL,
    timestamp  TEXT NOT NULL,
    comment    TEXT NOT NULL DEFAULT '',
    created    TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    updated    TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    PRIMARY KEY (id)
) STRICT;
CREATE TABLE battery
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
INSERT INTO battery VALUES('5a19768e-c5c4-479f-8709-a96ca946d705','Test Battery 01','Brand #1',0,'','','','','','2025-01-07T08:25:55.180Z','2025-01-07T08:25:55.180Z');
INSERT INTO battery VALUES('e03f5ddf-efed-4362-ba42-75b1b714d43a','Test Battery 02','Brand #1',0,'','','','','','2025-01-07T08:26:20.162Z','2025-01-07T08:26:20.162Z');
COMMIT;
