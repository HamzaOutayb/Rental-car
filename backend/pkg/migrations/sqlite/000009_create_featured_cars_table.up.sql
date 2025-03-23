-- +migrate Up
CREATE TABLE featured_cars (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    car_id INTEGER,
    feature_type TEXT NOT NULL,  -- 'main', 'type_main'
    type_id INTEGER,             -- NULL for main,
    display_order INTEGER,
    FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES type(id) ON DELETE CASCADE
);