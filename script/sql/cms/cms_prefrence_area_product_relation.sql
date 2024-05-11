create table cms_prefrence_area_product_relation
(
    id                bigint auto_increment
        primary key,
    prefrence_area_id bigint not null,
    product_id        bigint not null
)
    comment '优选专区和产品关系表' charset = utf8;

INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (1, 1, 12);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (2, 1, 13);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (3, 1, 14);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (4, 1, 18);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (5, 1, 7);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (6, 2, 7);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (7, 1, 22);
INSERT INTO cms_prefrence_area_product_relation (id, prefrence_area_id, product_id) VALUES (24, 1, 23);
