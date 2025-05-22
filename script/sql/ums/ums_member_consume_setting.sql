drop table if exists ums_points_consume_setting;
create table ums_points_consume_setting
(
    id                    bigint auto_increment
        primary key,
    deduction_per_amount  int                                not null comment '每一元需要抵扣的积分数量',
    max_percent_per_order int                                not null comment '每笔订单最高抵用百分比',
    use_unit              int                                not null comment '每次使用积分最小单位100',
    coupon_status         tinyint                            not null comment '是否可以和优惠券同用；0->不可以；1->可以',
    status                tinyint  default 1                 not null comment '状态：0->禁用；1->启用',
    create_by             bigint                             not null comment '创建人ID',
    create_time           datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by             bigint                             null comment '更新人ID',
    update_time           datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '积分消费设置';

-- 向积分消费设置表添加模拟数据
INSERT INTO ums_points_consume_setting
(deduction_per_amount, max_percent_per_order, use_unit, coupon_status, status, create_by)
VALUES (10, 50, 100, 0, 1, 1), -- 每1元抵扣10积分，最高抵扣订单金额的50%，最小使用单位100积分，不可与优惠券同用
       (20, 70, 100, 1, 0, 1), -- 每1元抵扣20积分，最高抵扣订单金额的70%，最小使用单位100积分，可与优惠券同用
       (5, 30, 50, 1, 0, 1); -- 每1元抵扣5积分，最高抵扣订单金额的30%，最小使用单位50积分，可与优惠券同用