create table ums_integration_consume_setting
(
    id                    bigint auto_increment
        primary key,
    deduction_per_amount  int               not null comment '每一元需要抵扣的积分数量',
    max_percent_per_order int               not null comment '每笔订单最高抵用百分比',
    use_unit              int               not null comment '每次使用积分最小单位100',
    is_default            tinyint default 0 not null comment '是否默认：0->否；1->是',
    coupon_status         tinyint           not null comment '是否可以和优惠券同用；0->不可以；1->可以'
)
    comment '积分消费设置';

INSERT INTO ums_integration_consume_setting (id, deduction_per_amount, max_percent_per_order, use_unit, is_default, coupon_status) VALUES (1, 100, 50, 100, 1, 1);
