create table oms_order_return_reason
(
    id          bigint auto_increment
        primary key,
    name        varchar(100) null comment '退货类型',
    sort        int          null,
    status      int(1)       null comment '状态：0->不启用；1->启用',
    create_time datetime     null comment '添加时间'
)
    comment '退货原因表';

INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (1, '质量问题', 1, 1, '2018-10-17 10:00:45');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (2, '尺码太大', 1, 1, '2018-10-17 10:01:03');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (3, '颜色不喜欢', 1, 1, '2018-10-17 10:01:13');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (4, '7天无理由退货', 1, 1, '2018-10-17 10:01:47');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (5, '价格问题', 1, 0, '2018-10-17 10:01:57');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (12, '发票问题', 0, 1, '2018-10-19 16:28:36');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (13, '其他问题', 0, 1, '2018-10-19 16:28:51');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (14, '物流问题', 0, 1, '2018-10-19 16:29:01');
INSERT INTO gozero.oms_order_return_reason (id, name, sort, status, create_time) VALUES (15, '售后问题', 0, 1, '2018-10-19 16:29:11');