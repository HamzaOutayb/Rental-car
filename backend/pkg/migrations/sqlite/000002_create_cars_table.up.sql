-- +migrate Up
CREATE TABLE cars (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price REAL NOT NULL,
    availability_status INTEGER NOT NULL DEFAULT 1,
    return_date TEXT, -- 
    brand_id INTEGER,
    type_id INTEGER,
    contact_id INTEGER,
    local_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE CASCADE,
    FOREIGN KEY (type_id) REFERENCES type(id) ON DELETE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts(id) ON DELETE CASCADE,
    FOREIGN KEY (local_id) REFERENCES local(id) ON DELETE CASCADE
);