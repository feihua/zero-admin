drop table if exists ums_member_task;
create table ums_member_task
(
    id             bigint auto_increment comment '主键ID'
        primary key,
    task_name      varchar(100)                           not null comment '任务名称',
    task_desc      varchar(500) default ''                not null comment '任务描述',
    task_growth    int          default 0                 not null comment '赠送成长值',
    task_integral  int          default 0                 not null comment '赠送积分',
    task_type      tinyint      default 0                 not null comment '任务类型：0-新手任务，1-日常任务，2-周常任务，3-月常任务',
    complete_count int          default 1                 not null comment '需要完成次数',
    reward_type    tinyint      default 0                 not null comment '奖励类型：0-积分成长值，1-优惠券，2-抽奖次数',
    reward_params  varchar(200) default ''                not null comment '奖励参数JSON',
    start_time     datetime                               not null comment '任务开始时间',
    end_time       datetime                               not null comment '任务结束时间',
    status         tinyint      default 1                 not null comment '状态：0-禁用，1-启用',
    sort           int          default 0                 not null comment '排序',
    create_by      bigint                                 not null comment '创建人ID',
    create_time    datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      bigint                                 null comment '更新人ID',
    update_time    datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint      default 0                 not null comment '是否删除'
)
    comment '会员任务表';

create index idx_time
    on ums_member_task (start_time, end_time, status);

create index idx_type_status
    on ums_member_task (task_type, status, is_deleted);


-- 插入会员任务数据
insert into ums_member_task (id, task_name, task_desc, task_growth, task_integral, task_type, complete_count, reward_type, reward_params, start_time, end_time, status, sort, create_by, is_deleted)
values (1, '注册奖励', '完成注册即可获得奖励', 100, 50, 0, 1, 0, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 1, 1, 0),
       (2, '每日签到', '每日签到可获得成长值和积分', 10, 5, 1, 1, 0, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 2, 1, 0),
       (3, '每周购物', '每周购物满100元可获得奖励', 50, 20, 2, 1, 0, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 3, 1, 0),
       (4, '每月评价', '每月评价满5次可获得奖励', 30, 15, 3, 5, 0, '{"coupon_id": null}', '2024-01-01 00:00:00', '2024-12-31 23:59:59', 1, 4, 1, 0);

# 注册奖励和每日签到的任务，触发时机：
#  1.实时触发 ：
#    - 当会员注册奖励和每日签到时，立即检查该会员是否达到任务要求。如果达到，则立即发放奖励。


# 每周购物和每月评价的任务，触发时机：
#
# 1. 每周定时任务 ：
#    - 设置一个每周定时任务（例如每周一凌晨），检查所有会员在过去一周内的购物情况。如果会员在过去一周内的购物金额或订单数量达到任务要求，则立即发放奖励。

# 2. 实时触发 ：
#    - 当会员完成一次购物时，立即检查该会员在当前周的购物总额或订单数量是否达到任务要求。如果达到，则立即发放奖励。

drop table if exists ums_member_task_relation;
CREATE TABLE `ums_member_task_relation`
(
    `id`        BIGINT                             NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `member_id` BIGINT                             NOT NULL COMMENT '会员ID',
    `task_id`   BIGINT                             NOT NULL COMMENT '任务ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='会员任务关联表';