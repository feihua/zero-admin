create table ums_member_statistics_info
(
    id                    bigint auto_increment
        primary key,
    member_id             bigint         not null,
    consume_amount        decimal(10, 2) not null comment '累计消费金额',
    order_count           int            not null comment '订单数量',
    coupon_count          int            not null comment '优惠券数量',
    comment_count         int            not null comment '评价数',
    return_order_count    int            not null comment '退货数量',
    login_count           int            not null comment '登录次数',
    attend_count          int            not null comment '关注数量',
    fans_count            int            not null comment '粉丝数量',
    collect_product_count int            not null,
    collect_subject_count int            not null,
    collect_topic_count   int            not null,
    collect_comment_count int            not null,
    invite_friend_count   int            not null,
    recent_order_time     datetime       not null comment '最后一次下订单时间'
)
    comment '会员统计信息';

