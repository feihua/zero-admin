drop table if exists ums_member_growth_log;
create table ums_member_growth_log
(
    id            bigint auto_increment
        primary key,
    member_id     bigint                                 not null comment '会员ID',
    change_type   tinyint                                not null comment '变更类型：1-添加成长值，2-减少成长值',
    change_growth int                                    not null comment '变更成长值',
    source_type   tinyint                                not null comment '来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改',
    description   varchar(200)                           not null comment '描述',
    operate_man   varchar(100) default ''                not null comment '操作人员',
    operate_note  varchar(200) default ''                not null comment '操作备注',
    create_time   datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '会员成长值记录表';

create index idx_member
    on ums_member_growth_log (member_id);


-- 插入会员成长值记录数据
insert into ums_member_growth_log (id, member_id, change_type, change_growth, source_type, description, operate_man, operate_note)
values (1, 1001, 1, 100, 1, '订单完成增加成长值', '管理员A', '订单ID: 12345'),
       (2, 1002, 2, 50, 2, '活动参与减少成长值', '管理员B', '活动ID: 67890'),
       (3, 1003, 1, 200, 3, '签到增加成长值', '管理员C', '签到ID: 11223'),
       (4, 1004, 2, 30, 4, '管理员修改减少成长值', '管理员D', '手动调整'),
       (5, 1005, 1, 150, 0, '其他来源增加成长值', '管理员E', '特殊奖励');