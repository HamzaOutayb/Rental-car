-- +migrate Up
CREATE TABLE local (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    addr TEXT,
    IG TEXT,
    facebook TEXT,
    watssapp TEXT,
    email TEXT,
    phone-num TEXT,
    time_id INTEGER,
    FOREIGN KEY (time_id) REFERENCES opening_hours(id) ON DELETE CASCADE,
);