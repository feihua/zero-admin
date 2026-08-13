drop table if exists sms_coupon;
create table sms_coupon
(
    id             bigint auto_increment comment '优惠券ID'
        primary key,
    type_id        bigint                                 not null comment '优惠券类型ID',
    name           varchar(100)                           not null comment '优惠券名称',
    code           varchar(32)                            not null comment '优惠券码',
    amount         decimal(10, 2)                         not null comment '优惠金额/折扣率',
    min_amount     decimal(10, 2)                         not null comment '最低使用金额',
    start_time     datetime                               not null comment '生效时间',
    end_time       datetime                               not null comment '失效时间',
    total_count    int                                    not null comment '发放总量',
    received_count int          default 0                 not null comment '已领取数量',
    used_count     int          default 0                 not null comment '已使用数量',
    per_limit      int          default 1                 not null comment '每人限领数量',
    status         tinyint      default 0                 not null comment '状态：0-未开始，1-进行中，2-已结束，3-已取消',
    is_enabled     tinyint      default 1                 not null comment '是否启用',
    description    varchar(500) default ''                not null comment '使用说明',
    create_by      bigint                                 not null comment '创建人ID',
    create_time    datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      bigint                                 null comment '更新人ID',
    update_time    datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint      default 0                 not null comment '是否删除'
)
    comment '优惠券表';

-- 插入优惠券数据
insert into sms_coupon (type_id, name, code, amount, min_amount, start_time, end_time, total_count, received_count, used_count, per_limit, status, is_enabled, description, create_by, is_deleted)
values (1, '满减优惠券', 'FULLREDUCE', 20.00, 150.00, '2023-11-01 00:00:00', '2029-11-30 23:59:59', 300, 0, 0, 1, 1, 1, '满150减20', 3, 0),
       (3, '新用户优惠券', 'NEWUSER2023', 50.00, 200.00, '2023-11-01 00:00:00', '2029-12-31 23:59:59', 1000, 0, 0, 1, 0, 1, '适用于新用户首次购物', 1, 0),
       (5, '双十一折扣券', 'DOUBLE11', 10.00, 100.00, '2023-11-11 00:00:00', '2029-11-11 23:59:59', 500, 0, 0, 2, 1, 1, '双十一当天使用', 2, 0);
