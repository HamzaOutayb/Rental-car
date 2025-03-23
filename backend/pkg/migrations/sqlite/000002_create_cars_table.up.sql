-- +migrate Up
CREATE TABLE cars (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price TEXT NOT NULL,
    images TEXT,
    avaibility TEXT,
    return_date TEXT,
    brand_id INTEGER,
    type_id INTEGER,
    Contact_id INTEGER,
    main INTEGER NOT NULL DEFAULT 0,
    main_type INTEGER NOT NULL DEFAULT 0,
    local_id INTEGER NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE CASCADE
    FOREIGN KEY (type_id) REFERENCES type(id) ON DELETE CASCADE
    FOREIGN KEY (Contact_id) REFERENCES Contacts(id) ON DELETE CASCADE
    FOREIGN KEY (local_id) REFERENCES local(id) ON DELETE CASCADE
);