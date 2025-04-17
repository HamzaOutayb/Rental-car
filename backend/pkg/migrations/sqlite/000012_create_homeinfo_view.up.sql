-- +migrate Up
CREATE VIEW shop_view AS
SELECT
    si.id AS ShopID,
    b.name AS Brands,
    si.name AS Name,
    si.addr AS Addr,
    si.IG,
    si.facebook AS Facebook,
    si.watssapp AS Watssapp,
    si.email AS Email,
    si."phone-num" AS Phone_num,
    s.day_of_week AS Days_of_week,
    s.open_time AS Open_time,
    s.close_time AS Close_time
FROM shop_info si
LEFT JOIN cars c ON si.id = c.local_id
LEFT JOIN brands b ON c.brand_id = b.id
LEFT JOIN schedule s ON si.schedule_id = s.id
LEFT JOIN type t ON c.type_id = t.id
WHERE b.count > 0 AND t.count > 0;