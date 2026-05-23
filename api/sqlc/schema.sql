-- Enable foreign key support (highly recommended to run on every connection)
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS paths (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    path        TEXT NOT NULL,
    parent_id   INTEGER,
    information TEXT, -- SQLite stores JSON as TEXT, but supports full JSON functions
    
    FOREIGN KEY (parent_id) REFERENCES paths(id) ON DELETE CASCADE
);

-- Separate non-unique indexes for fast lookups
CREATE INDEX IF NOT EXISTS idx_paths_path ON paths(path);
CREATE INDEX IF NOT EXISTS idx_paths_parent_id ON paths(parent_id);