create table sms_flash_promotion_product_relation
(
    id                         bigint auto_increment comment '编号'
        primary key,
    flash_promotion_id         bigint not null comment '限时购id',
    flash_promotion_session_id bigint not null comment '编号',
    product_id                 bigint not null comment '商品id',
    flash_promotion_price      bigint not null comment '限时购价格',
    flash_promotion_count      int    not null comment '限时购数量',
    flash_promotion_limit      int    not null comment '每人限购数量',
    sort                       int    not null comment '排序'
)
    comment '商品限时购与商品关系表';

INSERT INTO sms_flash_promotion_product_relation (id, flash_promotion_id, flash_promotion_session_id, product_id, flash_promotion_price, flash_promotion_count, flash_promotion_limit, sort) VALUES (1, 2, 1, 26, 3000.00, 10, 1, 0);
INSERT INTO sms_flash_promotion_product_relation (id, flash_promotion_id, flash_promotion_session_id, product_id, flash_promotion_price, flash_promotion_count, flash_promotion_limit, sort) VALUES (2, 2, 1, 27, 2000.00, 10, 1, 20);
INSERT INTO sms_flash_promotion_product_relation (id, flash_promotion_id, flash_promotion_session_id, product_id, flash_promotion_price, flash_promotion_count, flash_promotion_limit, sort) VALUES (3, 2, 1, 28, 599.00, 19, 1, 0);
INSERT INTO sms_flash_promotion_product_relation (id, flash_promotion_id, flash_promotion_session_id, product_id, flash_promotion_price, flash_promotion_count, flash_promotion_limit, sort) VALUES (4, 2, 1, 29, 4999.00, 10, 1, 100);
INSERT INTO sms_flash_promotion_product_relation (id, flash_promotion_id, flash_promotion_session_id, product_id, flash_promotion_price, flash_promotion_count, flash_promotion_limit, sort) VALUES (9, 2, 2, 26, 2999.00, 100, 1, 0);
