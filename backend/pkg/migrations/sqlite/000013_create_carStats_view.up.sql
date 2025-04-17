-- +migrate UP
CREATE VIEW car_detailed_view AS
SELECT
    c.id AS car_id,
    c.name AS car_name,
    c.description AS car_description,
    c.price AS car_price,
    c.availability_status,
    b.brand AS brand_name,
    t.type AS type_name,
    co.name AS contact_name,
    co.telegram AS contact_telegram,
    co.watssapp AS contact_watssapp
FROM cars c
LEFT JOIN brands b ON c.brand_id = b.id
LEFT JOIN type t ON c.type_id = t.id
LEFT JOIN contacts co ON c.contact_id = co.id