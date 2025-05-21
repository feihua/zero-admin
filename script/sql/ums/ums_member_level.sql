drop table if exists ums_member_level;
-- 会员等级表
create table gozero.ums_member_level
(
    id            bigint auto_increment
        primary key comment '主键ID',
    name          varchar(50)                             not null comment '等级名称',
    level         int                                     not null comment '等级',
    growth_point  int                                     not null comment '升级所需成长值',
    discount_rate decimal(5, 2) default 100.00            not null comment '折扣率(0-100)',
    free_freight  tinyint       default 0                 not null comment '是否免运费',
    comment_extra tinyint       default 0                 not null comment '是否可评论获取奖励',
    privileges    varchar(500)  default ''                not null comment '会员特权JSON',
    remark        varchar(200)  default ''                not null comment '备注',
    is_enabled    tinyint       default 1                 not null comment '是否启用',
    create_by     bigint                                  not null comment '创建人ID',
    create_time   datetime      default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     bigint                                  null comment '更新人ID',
    update_time   datetime                                null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted    tinyint       default 0                 not null comment '是否删除',
    constraint uk_level
        unique (level, is_deleted)
)
    comment '会员等级表';

#特权字段说明：
# {
#     "priority_service": 0,      // 优先客服：0-无，1-普通优先，2-高优先级，3-专属客服
#     "birthday_gift": 0,         // 生日礼包：0-无，1-普通礼包，2-高级礼包，3-豪华礼包
#     "exclusive_price": 0,       // 专属价格：0-无，1-有
#     "vip_room": 0,             // VIP专区：0-无，1-有
#     "free_return": 0,          // 免费退换：0-无，1-有
#     "personal_butler": 0,       // 专属管家：0-无，1-有
#     "early_access": 0,         // 新品优先购：0-无，1-提前1天，2-提前3天，3-提前7天
#     "point_rate": 1.0          // 积分倍率：1.0-1倍，1.2-1.2倍，以此类推
# }

-- 添加会员等级
insert into gozero.ums_member_level (id, name, level, growth_point, discount_rate, free_freight, comment_extra, privileges, remark, is_enabled, create_by)
values (1, '普通会员', 1, 0, 100.00, 0, 0, '{
    "priority_service": 0,
    "birthday_gift": 0,
    "exclusive_price": 0,
    "vip_room": 0,
    "free_return": 0,
    "personal_butler": 0,
    "early_access": 0,
    "point_rate": 1.0
 }', '新注册会员默认等级', 1, 1),
       (2, '银卡会员', 2, 1000, 98.00, 0, 1, '{
    "priority_service": 1,
    "birthday_gift": 1,
    "exclusive_price": 0,
    "vip_room": 0,
    "free_return": 0,
    "personal_butler": 0,
    "early_access": 1,
    "point_rate": 1.2
 }', '消费满1000成长值可升级', 1, 1),
       (3, '金卡会员', 3, 3000, 95.00, 1, 1, '{
    "priority_service": 2,
    "birthday_gift": 2,
    "exclusive_price": 1,
    "vip_room": 0,
    "free_return": 1,
    "personal_butler": 0,
    "early_access": 2,
    "point_rate": 1.5
 }', '消费满3000成长值可升级', 1, 1),
       (4, '钻石会员', 4, 10000, 92.00, 1, 1, '{
    "priority_service": 2,
    "birthday_gift": 2,
    "exclusive_price": 1,
    "vip_room": 1,
    "free_return": 1,
    "personal_butler": 0,
    "early_access": 2,
    "point_rate": 2.0
 }', '消费满10000成长值可升级', 1, 1),
       (5, '黑金会员', 5, 50000, 88.00, 1, 1, '{
    "priority_service": 3,
    "birthday_gift": 3,
    "exclusive_price": 1,
    "vip_room": 1,
    "free_return": 1,
    "personal_butler": 1,
    "early_access": 3,
    "point_rate": 3.0
 }', '年消费达到50000可升级', 1, 1);