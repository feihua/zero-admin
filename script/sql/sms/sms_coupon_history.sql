drop table if exists sms_coupon_history;
create table sms_coupon_history
(
    id              bigint auto_increment
        primary key,
    coupon_id       bigint       not null comment '优惠券id',
    member_id       bigint       not null comment '会员id',
    coupon_code     varchar(64)  not null comment '优惠码',
    member_nickname varchar(64)  not null comment '领取人昵称',
    get_type        tinyint      not null comment '获取类型：0->后台赠送；1->主动获取',
    use_status      tinyint      not null comment '使用状态：0->未使用；1->已使用；2->已过期',
    use_time        datetime     null comment '使用时间',
    order_id        bigint       null comment '订单编号',
    order_sn        varchar(100) null comment '订单号码',
    create_time     datetime     not null comment '领取时间'
)
    comment '优惠券使用、领取历史表';

create index idx_coupon_id
    on sms_coupon_history (coupon_id);

create index idx_member_id
    on sms_coupon_history (member_id);


INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (1, 1, '4931048380330001', 'test', 1, current_time, 0);
INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (2, 1, '4931048380330002', 'test', 1, current_time, 0);
INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (3, 1, '4931048380330003', 'test', 1, current_time, 0);
INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (1, 2, '4931048380330001', 'koobe', 1, current_time, 0);
INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (2, 2, '4931048380330002', 'koobe', 1, current_time, 0);
INSERT INTO sms_coupon_history (coupon_id, member_id, coupon_code, member_nickname, get_type, create_time, use_status) VALUES (3, 2, '4931048380330003', 'koobe', 1, current_time, 0);


