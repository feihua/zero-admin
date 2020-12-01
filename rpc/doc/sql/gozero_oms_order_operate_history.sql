create table oms_order_operate_history
(
    id           bigint auto_increment
        primary key,
    order_id     bigint       null comment '订单id',
    operate_man  varchar(100) null comment '操作人：用户；系统；后台管理员',
    create_time  datetime     null comment '操作时间',
    order_status int(1)       null comment '订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
    note         varchar(500) null comment '备注'
)
    comment '订单操作历史记录';

INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (5, 12, '后台管理员', '2018-10-12 14:01:29', 2, '完成发货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (6, 13, '后台管理员', '2018-10-12 14:01:29', 2, '完成发货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (7, 12, '后台管理员', '2018-10-12 14:13:10', 4, '订单关闭:买家退货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (8, 13, '后台管理员', '2018-10-12 14:13:10', 4, '订单关闭:买家退货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (9, 22, '后台管理员', '2018-10-15 16:31:48', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (10, 22, '后台管理员', '2018-10-15 16:35:08', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (11, 22, '后台管理员', '2018-10-15 16:35:59', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (12, 17, '后台管理员', '2018-10-15 16:43:40', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (13, 25, '后台管理员', '2018-10-15 16:52:14', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (14, 26, '后台管理员', '2018-10-15 16:52:14', 4, '订单关闭:xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (15, 23, '后台管理员', '2018-10-16 14:41:28', 2, '完成发货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (16, 13, '后台管理员', '2018-10-16 14:42:17', 2, '完成发货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (17, 18, '后台管理员', '2018-10-16 14:42:17', 2, '完成发货');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (18, 26, '后台管理员', '2018-10-30 14:37:44', 4, '订单关闭:关闭订单');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (19, 25, '后台管理员', '2018-10-30 15:07:01', 0, '修改收货人信息');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (20, 25, '后台管理员', '2018-10-30 15:08:13', 0, '修改费用信息');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (21, 25, '后台管理员', '2018-10-30 15:08:31', 0, '修改备注信息：xxx');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (22, 25, '后台管理员', '2018-10-30 15:08:39', 4, '订单关闭:2222');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (23, 12, '后台管理员', '2019-11-09 16:50:28', 4, '修改备注信息：111');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (24, 30, '后台管理员', '2020-02-25 16:52:37', 0, '修改费用信息');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (25, 30, '后台管理员', '2020-02-25 16:52:51', 0, '修改费用信息');
INSERT INTO gozero.oms_order_operate_history (id, order_id, operate_man, create_time, order_status, note) VALUES (26, 30, '后台管理员', '2020-02-25 16:54:03', 2, '完成发货');