-- +migrate Up
CREATE TABLE
    contacts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        telegram TEXT,
        watssapp TEXT,
    );