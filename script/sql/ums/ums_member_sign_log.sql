drop table if exists ums_member_sign_log;
create table gozero.ums_member_sign_log
(
    id            bigint auto_increment
        primary key,
    member_id     bigint                             not null comment '会员ID',
    sign_date     date                               not null comment '签到日期',
    continue_days int      default 1                 not null comment '连续签到天数',
    points        int      default 0                 not null comment '获得积分',
    create_time   datetime default CURRENT_TIMESTAMP not null,
    constraint uk_member_date
        unique (member_id, sign_date)
)
    comment '会员签到记录表';


-- 插入会员签到记录数据
insert into gozero.ums_member_sign_log (id, member_id, sign_date, continue_days, points)
values (1, 1001, '2024-01-01', 1, 10),
       (2, 1002, '2024-01-02', 2, 20),
       (3, 1003, '2024-01-03', 3, 30),
       (4, 1004, '2024-01-04', 4, 40),
       (5, 1005, '2024-01-05', 5, 50);