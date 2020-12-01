create table ums_integration_consume_setting
(
    id                    bigint auto_increment
        primary key,
    deduction_per_amount  int    null comment '每一元需要抵扣的积分数量',
    max_percent_per_order int    null comment '每笔订单最高抵用百分比',
    use_unit              int    null comment '每次使用积分最小单位100',
    coupon_status         int(1) null comment '是否可以和优惠券同用；0->不可以；1->可以'
)
    comment '积分消费设置';

INSERT INTO gozero.ums_integration_consume_setting (id, deduction_per_amount, max_percent_per_order, use_unit, coupon_status) VALUES (1, 100, 50, 100, 1);