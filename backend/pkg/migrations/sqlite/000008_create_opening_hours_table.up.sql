-- +migrate Up
CREATE TABLE
    opening_hours (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        day_of_week TEXT NOT NULL,
        open_time TEXT,
        close_time TEXT,
    );