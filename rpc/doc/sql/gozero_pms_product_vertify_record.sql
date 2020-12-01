create table pms_product_vertify_record
(
    id          bigint auto_increment
        primary key,
    product_id  bigint       null,
    create_time datetime     null,
    vertify_man varchar(64)  null comment '审核人',
    status      int(1)       null,
    detail      varchar(255) null comment '反馈详情'
)
    comment '商品审核记录';

INSERT INTO gozero.pms_product_vertify_record (id, product_id, create_time, vertify_man, status, detail) VALUES (1, 1, '2018-04-27 16:36:41', 'test', 1, '验证通过');
INSERT INTO gozero.pms_product_vertify_record (id, product_id, create_time, vertify_man, status, detail) VALUES (2, 2, '2018-04-27 16:36:41', 'test', 1, '验证通过');