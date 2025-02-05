create table ums_member_statistics_info
(
    id                    bigint auto_increment
        primary key,
    member_id             bigint   not null comment '会员id',
    consume_amount        bigint   not null comment '累计消费金额',
    order_count           int      not null comment '订单数量',
    coupon_count          int      not null comment '优惠券数量',
    comment_count         int      not null comment '评价数',
    return_order_count    int      not null comment '退货数量',
    login_count           int      not null comment '登录次数',
    attend_count          int      not null comment '关注数量',
    fans_count            int      not null comment '粉丝数量',
    collect_product_count int      not null comment '收藏的商品数量',
    collect_subject_count int      not null comment '收藏的专题活动数量',
    collect_topic_count   int      not null comment '收藏的评论数量',
    collect_comment_count int      not null comment '收藏的专题活动数量',
    invite_friend_count   int      not null comment '邀请好友数',
    recent_order_time     datetime not null comment '最后一次下订单时间'
)
    comment '会员统计信息';

