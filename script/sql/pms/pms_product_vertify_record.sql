create table pms_product_vertify_record
(
    id          bigint auto_increment
        primary key,
    product_id  bigint                             not null comment '商品id',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    review_man  varchar(64)                        not null comment '审核人',
    status      tinyint                            not null comment '审核状态：0->未通过；1->通过',
    detail      varchar(255)                       not null comment '反馈详情'
)
    comment '商品审核记录';

INSERT INTO pms_product_vertify_record (id, product_id, create_time, review_man, status, detail) VALUES (1, 1, current_time, 'test', 1, '验证通过');
INSERT INTO pms_product_vertify_record (id, product_id, create_time, review_man, status, detail) VALUES (2, 2, current_time, 'test', 1, '验证通过');
