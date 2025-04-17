-- +migrate Up
CREATE TABLE shop_info (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    addr TEXT,
    IG TEXT,
    facebook TEXT,
    watssapp TEXT,
    email TEXT,
    phone_num TEXT,
    schedule_id INTEGER,
    FOREIGN KEY (schedule_id) REFERENCES schedule(id) ON DELETE CASCADE
);