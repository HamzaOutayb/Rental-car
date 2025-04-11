-- +migrate Up
CREATE TABLE car_conditions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    condition JSON,
    car_id TEXT NOT NULL,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
);