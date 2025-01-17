CREATE TABLE IF NOT EXISTS battery_history
(
    id         TEXT NOT NULL,
    battery_id TEXT NOT NULL,
    timestamp  TEXT NOT NULL,
    comment    TEXT NOT NULL DEFAULT '',
    created    TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    updated    TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ')),
    PRIMARY KEY (id)
) STRICT;
