create table pms_product_category_attribute_relation
(
    id                   bigint auto_increment
        primary key,
    product_category_id  bigint not null comment '商品分类id',
    product_attribute_id bigint not null comment '商品属性id'
)
    comment '产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）';

INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (1, 24, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (5, 26, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (7, 28, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (9, 25, 24);
INSERT INTO pms_product_category_attribute_relation (id, product_category_id, product_attribute_id) VALUES (10, 25, 25);
