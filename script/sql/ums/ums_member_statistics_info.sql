drop table if exists ums_member_statistics_info;
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

-- 插入会员统计信息数据
insert into gozero.ums_member_statistics_info (id, member_id, consume_amount, order_count, coupon_count, comment_count, return_order_count, login_count, attend_count, fans_count, collect_product_count, collect_subject_count, collect_topic_count, collect_comment_count, invite_friend_count, recent_order_time)
values (1, 1001, 5000, 10, 5, 20, 2, 50, 10, 100, 15, 5, 3, 8, 2, '2024-01-01 10:00:00'),
       (2, 1002, 3000, 8, 3, 15, 1, 30, 5, 50, 10, 3, 2, 5, 1, '2024-01-02 11:00:00'),
       (3, 1003, 7000, 15, 7, 25, 3, 70, 20, 150, 20, 10, 5, 12, 4, '2024-01-03 12:00:00'),
       (4, 1004, 2000, 5, 2, 10, 0, 20, 3, 30, 5, 2, 1, 3, 0, '2024-01-04 13:00:00'),
       (5, 1005, 10000, 20, 10, 30, 5, 100, 25, 200, 25, 15, 8, 20, 5, '2024-01-05 14:00:00');