drop table if exists `sms_seckill_activity`;
create table sms_seckill_activity
(
    id          bigint auto_increment comment '编号'
        primary key,
    name        varchar(100)                           not null comment '活动名称',
    description varchar(500) default ''                not null comment '活动描述',
    start_time  datetime                               not null comment '开始时间',
    end_time    datetime                               not null comment '结束时间',
    status      tinyint      default 0                 not null comment '活动状态：0-未开始，1-进行中，2-已结束，3-已取消',
    is_enabled  tinyint      default 1                 not null comment '是否启用',
    create_by   bigint                                 not null comment '创建人ID',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   bigint                                 null comment '更新人ID',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted  tinyint      default 0                 not null comment '是否删除'
)
    comment '秒杀活动表';

-- 插入秒杀活动数据
insert into sms_seckill_activity (id, name, description, start_time, end_time, status, is_enabled, create_by, is_deleted)
values (1, '双十一秒杀活动', '双十一期间的限时秒杀活动', '2025-11-11 00:00:00', '2025-11-11 23:59:59', 0, 1, 1, 0),
       (2, '黑色星期五秒杀', '黑色星期五特惠秒杀活动', '2025-11-24 00:00:00', '2025-11-24 23:59:59', 0, 1, 2, 0),
       (3, '圣诞节秒杀', '圣诞节期间的限时秒杀活动', '2025-12-25 00:00:00', '2025-12-25 23:59:59', 0, 1, 3, 0);