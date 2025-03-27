drop table if exists oms_order_operate_history;
create table oms_order_operate_history
(
    id           bigint auto_increment
        primary key,
    order_id     bigint                                 not null comment '订单id',
    operate_man  varchar(100)                           not null comment '操作人：用户；系统；后台管理员',
    create_time  datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    order_status tinyint                                not null comment '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
    note         varchar(500) default ''                not null comment '备注'
)
    comment '订单操作历史记录';

INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (12, '后台管理员', 2, '完成发货');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (13, '后台管理员', 2, '完成发货');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (12, '后台管理员', 4, '订单关闭:买家退货');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (13, '后台管理员', 4, '订单关闭:买家退货');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (17, '后台管理员', 4, '订单关闭:xxx');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (13, '后台管理员', 2, '完成发货');
INSERT INTO oms_order_operate_history (order_id, operate_man, order_status, note) VALUES (12, '后台管理员', 4, '修改备注信息：111');

