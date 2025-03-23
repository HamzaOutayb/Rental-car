-- +migrate Up
CREATE TABLE car_images (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    car_id INTEGER,
    image_path TEXT NOT NULL,
    is_primary BOOLEAN DEFAULT 0,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
);