CREATE TABLE backups (
    id         INTEGER PRIMARY KEY AUTOINCREMENT
                       NOT NULL,
    shop       TEXT    NOT NULL,
    title      TEXT    NOT NULL,
    content    BLOB    NOT NULL,
    created_at TIME    NOT NULL
);