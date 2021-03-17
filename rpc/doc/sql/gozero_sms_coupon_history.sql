create table sms_coupon_history
(
    id              bigint auto_increment
        primary key,
    coupon_id       bigint       null,
    member_id       bigint       null,
    coupon_code     varchar(64)  null,
    member_nickname varchar(64)  null comment '领取人昵称',
    get_type        int(1)       null comment '获取类型：0->后台赠送；1->主动获取',
    create_time     datetime     null,
    use_status      int(1)       null comment '使用状态：0->未使用；1->已使用；2->已过期',
    use_time        datetime     null comment '使用时间',
    order_id        bigint       null comment '订单编号',
    order_sn        varchar(100) null comment '订单号码'
)
    comment '优惠券使用、领取历史表';

create index idx_coupon_id
    on sms_coupon_history (coupon_id);

create index idx_member_id
    on sms_coupon_history (member_id);

INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (2, 2, 1, '4931048380330002', 'windir', 1, '2018-08-29 14:04:12', 1, '2018-11-12 14:38:47', 12, '201809150101000001');
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (3, 3, 1, '4931048380330003', 'windir', 1, '2018-08-29 14:04:29', 0, '2018-11-12 14:38:47', 12, ' ');
INSERT INTO gozero.sms_coupon_history (id, coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status, use_time, order_id, order_sn) VALUES (4, 4, 1, '4931048380330004', 'windir', 1, '2018-08-29 14:04:32', 0, '2018-11-12 14:38:47 ', 12, ' ');
