drop table if exists sms_coupon_record;
create table sms_coupon_record
(
    id              bigint auto_increment comment '主键ID'
        primary key,
    coupon_id       bigint                                   not null comment '优惠券ID',
    member_id       bigint                                   not null comment '用户ID',
    get_time        datetime                                 not null comment '领取时间',
    get_type        tinyint                                  not null comment '获取类型：0->后台赠送；1->主动获取',
    use_time        datetime                                 null comment '使用时间',
    order_id        bigint         default 0                 not null comment '使用订单ID',
    order_amount    decimal(10, 2) default 0                 not null comment '订单金额',
    discount_amount decimal(10, 2) default 0                 not null comment '优惠金额',
    status          tinyint        default 0                 not null comment '状态：0-未使用，1-已使用，2-已过期，3-已失效',
    invalid_time    datetime                                 null comment '失效时间',
    invalid_reason  varchar(255)   default ''                not null comment '失效原因',
    create_time     datetime       default CURRENT_TIMESTAMP not null comment '创建时间',
    is_deleted      tinyint        default 0                 not null comment '是否删除'
)
    comment '优惠券领取记录表';

