-- +migrate Up
CREATE TABLE
    type (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        type INTEGER NOT NULL,
        count INTEGER NOT NULL DEFAULT 0,
    );