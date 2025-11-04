drop table if exists oms_order_promotion;
create table oms_order_promotion
(
    id              bigint auto_increment comment '主键ID'
        primary key,
    order_id        bigint                             not null comment '订单ID',
    order_no        varchar(32)                        not null comment '订单编号',
    promotion_type  tinyint                            not null comment '优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动',
    promotion_id    bigint                             null comment '优惠ID',
    promotion_name  varchar(100)                       not null comment '优惠名称',
    discount_amount decimal(10, 2)                     not null comment '优惠金额',
    create_time     datetime default CURRENT_TIMESTAMP not null,
    is_deleted      tinyint  default 0                 not null comment '是否删除'
)
    comment '订单优惠信息表';

create index idx_order
    on oms_order_promotion (order_id, is_deleted);

-- 模拟数据
insert into gozero.oms_order_promotion (id, order_id, order_no, promotion_type, promotion_id, promotion_name, discount_amount, is_deleted)
values (1, 1, 'ORD20240001', 1, 2001, '新人优惠券', 20.00, 0),
       (2, 1, 'ORD20240002', 2, null, '积分抵扣', 15.50, 0),
       (3, 2, 'ORD20240003', 3, null, '会员折扣', 10.00, 0),
       (4, 3, 'ORD20240004', 4, 3001, '618大促', 50.00, 0),
       (5, 4, 'ORD20240005', 1, 2002, '满减券', 30.00, 0);