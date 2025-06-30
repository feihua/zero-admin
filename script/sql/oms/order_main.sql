drop table if exists oms_order_main;
create table oms_order_main
(
    id                   bigint auto_increment
        primary key,
    order_no             varchar(32)                              not null comment '订单编号',
    user_id              bigint                                   not null comment '用户ID',
    order_status         tinyint        default 1                 not null comment '订单状态：1-待支付,2-已支付,3-已发货,4-已完成,5-已取消,6-已退款,7-售后中',
    total_amount         decimal(10, 2)                           not null comment '订单总金额',
    promotion_amount     decimal(10, 2) default 0.00              not null comment '促销金额',
    coupon_amount        decimal(10, 2) default 0.00              not null comment '优惠券金额',
    points_amount        decimal(10, 2) default 0.00              not null comment '积分金额',
    discount_amount      decimal(10, 2) default 0.00              not null comment '优惠金额',
    freight_amount       decimal(10, 2) default 0.00              not null comment '运费金额',
    pay_amount           decimal(10, 2)                           not null comment '实付金额',
    pay_type             tinyint                                  null comment '支付方式：1-支付宝,2-微信,3-银联',
    pay_time             datetime                                 null comment '支付时间',
    delivery_time        datetime                                 null comment '发货时间',
    receive_time         datetime                                 null comment '收货时间',
    comment_time         datetime                                 null comment '评价时间',
    source_type          tinyint        default 1                 not null comment '订单来源：1-APP,2-PC,3-小程序',
    express_order_number varchar(64)                              not null comment '快递单号',
    use_points           int            default 0                 not null comment '下单时使用的积分',
    receive_status       tinyint                                  not null comment '是否确认收货：0->否,1->是',
    remark               varchar(500)   default ''                not null comment '订单备注',
    create_time          datetime       default CURRENT_TIMESTAMP not null comment '提交时间',
    update_time          datetime                                 null on update CURRENT_TIMESTAMP,
    is_deleted           tinyint        default 0                 not null comment '是否删除',
    constraint uk_order_no
        unique (order_no, is_deleted)
)
    comment '订单主表';

create index idx_time
    on oms_order_main (create_time, is_deleted);

create index idx_user
    on oms_order_main (user_id, order_status, is_deleted);

