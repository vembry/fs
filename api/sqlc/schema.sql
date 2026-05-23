CREATE TABLE paths (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    path        TEXT NOT NULL,
    parent_id   INTEGER,
    information TEXT
);
