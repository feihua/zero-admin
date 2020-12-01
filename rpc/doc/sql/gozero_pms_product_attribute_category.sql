create table pms_product_attribute_category
(
    id              bigint auto_increment
        primary key,
    name            varchar(64)   null,
    attribute_count int default 0 null comment '属性数量',
    param_count     int default 0 null comment '参数数量'
)
    comment '产品属性分类表';

INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (1, '服装-T恤', 2, 5);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (2, '服装-裤装', 2, 4);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (3, '手机数码-手机通讯', 2, 4);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (4, '配件', 0, 0);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (5, '居家', 0, 0);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (6, '洗护', 0, 0);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (10, '测试分类', 0, 0);
INSERT INTO gozero.pms_product_attribute_category (id, name, attribute_count, param_count) VALUES (11, '服装-鞋帽', 3, 0);