drop table if exists ums_member_info;
create table ums_member_info
(
    id                 bigint auto_increment comment '主键ID'
        primary key,
    member_id          bigint                                   not null comment '会员ID',
    level_id           bigint                                   not null comment '等级ID',
    nickname           varchar(50)                              not null comment '昵称',
    mobile             varchar(11)                              not null comment '手机号码',
    source             tinyint        default 0                 not null comment '注册来源：0-PC，1-APP，2-小程序',
    password           varchar(100)                             not null comment '密码',
    avatar             varchar(255)   default ''                not null comment '头像',
    signature          varchar(200)   default ''                not null comment '个性签名',
    gender             tinyint        default 0                 not null comment '性别：0-未知，1-男，2-女',
    birthday           date                                     null comment '生日',
    growth_point       int            default 0                 not null comment '成长值',
    points             int            default 0                 not null comment '积分',
    total_points       int            default 0                 not null comment '累计获得积分',
    spend_amount       decimal(10, 2) default 0.00              not null comment '累计消费金额',
    order_count        int            default 0                 not null comment '订单数',
    coupon_count       int            default 0                 not null comment '优惠券数量',
    comment_count      int            default 0                 not null comment '评价数',
    return_count       int            default 0                 not null comment '退货数',
    lottery_times      int            default 0                 not null comment '剩余抽奖次数',
    first_login_status tinyint        default 1                 not null comment '1:未登录过,2:已登录过',
    last_login         datetime                                 null comment '最后登录',
    is_enabled         tinyint        default 1                 not null comment '是否启用：0-禁用，1-启用',
    create_time        datetime       default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time        datetime                                 null comment '更新时间',
    is_deleted         tinyint        default 0                 not null comment '是否删除',
    constraint uk_mobile
        unique (mobile, is_deleted),
    constraint uk_user
        unique (member_id, is_deleted)
)
    comment '会员信息表';

create index idx_level
    on ums_member_info (level_id, is_deleted);

create index idx_source
    on ums_member_info (source, is_deleted);

create index idx_status
    on ums_member_info (is_enabled, is_deleted);


-- 插入会员数据
insert into ums_member_info (id, member_id, level_id, nickname, mobile, source, password, avatar, signature, gender, birthday, growth_point, points, total_points, spend_amount, order_count, coupon_count, comment_count, return_count, lottery_times, last_login, is_enabled)
values (1, 1001, 1, '张三', '13800138001', 0, '123456', 'https://example.com/avatar/001.jpg', '生活就是购物~', 1, '1990-01-15', 100, 500, 1000, 999.99, 10, 5, 8, 1, 2, '2024-01-15 08:30:00', 1),
       (2, 1002, 2, '李四', '13800138002', 1, '123456', 'https://example.com/avatar/002.jpg', '潮流时尚，品质生活', 2, '1995-03-20', 500, 1200, 2000, 2999.99, 25, 8, 20, 0, 5, '2024-01-15 12:45:00', 1),
       (3, 1003, 3, '王五', '13800138003', 2, '123456', 'https://example.com/avatar/003.jpg', 'VIP专享生活', 1, '1988-07-08', 1000, 3000, 5000, 8999.99, 50, 15, 45, 2, 10, '2024-01-15 15:20:00', 1),
       (4, 1004, 4, '赵六', '13800138004', 1, '123456', 'https://example.com/avatar/004.jpg', '享受生活每一天', 2, '1992-12-25', 2000, 5000, 8000, 15999.99, 80, 20, 75, 3, 15, '2024-01-15 18:10:00', 1),
       (5, 1005, 5, '钱七', '13800138005', 0, '123456', 'https://example.com/avatar/005.jpg', '至尊购物体验', 1, '1985-05-01', 5000, 10000, 15000, 29999.99, 150, 30, 120, 5, 20, '2024-01-15 20:30:00', 1);