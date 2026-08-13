drop table if exists oms_order_delivery;
create table oms_order_delivery
(
    id                bigint auto_increment
        primary key,
    order_id          bigint                                not null comment '订单ID',
    order_no          varchar(32)                           not null comment '订单编号',
    receiver_name     varchar(50)                           not null comment '收货人姓名',
    receiver_phone    varchar(20)                           not null comment '收货人电话',
    receiver_province varchar(50)                           not null comment '省份',
    receiver_city     varchar(50)                           not null comment '城市',
    receiver_district varchar(50)                           not null comment '区县',
    receiver_address  varchar(200)                          not null comment '详细地址',
    delivery_company  varchar(50) default ''                not null comment '物流公司',
    delivery_no       varchar(50) default ''                not null comment '物流单号',
    create_time       datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time       datetime                              null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted        tinyint     default 0                 not null comment '是否删除',
    constraint uk_order
        unique (order_id, is_deleted)
)
    comment '订单收货地址表';

create index idx_delivery
    on oms_order_delivery (delivery_no, is_deleted);


INSERT INTO oms_order_delivery (
    order_id, order_no, receiver_name, receiver_phone, receiver_province,
    receiver_city, receiver_district, receiver_address, delivery_company,
    delivery_no, create_time, update_time
) VALUES
-- Delivery info for order ORD20231201001
(1, 'ORD20231201001', '张三', '13800138001', '北京市',
 '北京市', '朝阳区', '某某街道101号', '',
 '', '2023-12-01 10:30:00', NULL),

-- Delivery info for order ORD20231201002
(2, 'ORD20231201002', '李四', '13800138002', '上海市',
 '上海市', '浦东新区', '某某路202号', '',
 '', '2023-12-01 14:20:00', NULL),

-- Delivery info for order ORD20231202001
(3, 'ORD20231202001', '王五', '13800138003', '广东省',
 '广州市', '天河区', '某某大道303号', '顺丰速运',
 'SF123456789CN', '2023-12-02 09:15:00', '2023-12-02 14:30:00'),

-- Delivery info for order ORD20231203001
(4, 'ORD20231203001', '赵六', '13800138004', '广东省',
 '深圳市', '南山区', '某某科技园404号', '圆通速递',
 'YT987654321CN', '2023-12-03 16:45:00', '2023-12-04 09:00:00'),

-- Delivery info for order ORD20231204001
(5, 'ORD20231204001', '钱七', '13800138005', '浙江省',
 '杭州市', '西湖区', '某某广场505号', '',
 '', '2023-12-04 11:20:00', NULL),

-- Delivery info for order ORD20231205001
(6, 'ORD20231205001', '孙八', '13800138006', '江苏省',
 '南京市', '鼓楼区', '某某大厦606号', 'EMS',
 'EMS1122334455', '2023-12-05 15:30:00', '2023-12-06 10:00:00'),

-- Delivery info for order ORD20231206001
(7, 'ORD20231206001', '周九', '13800138007', '浙江省',
 '宁波市', '鄞州区', '某某商业街707号', '中通快递',
 'ZT5566778899', '2023-12-06 11:15:00', '2023-12-07 15:00:00');
