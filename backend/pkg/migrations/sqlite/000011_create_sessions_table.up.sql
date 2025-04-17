-- +migrate Up
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    sessions_exp TIME DEFAULT (TIME('now'))
);