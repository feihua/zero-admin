create table ums_member_task
(
    id           bigint auto_increment
        primary key,
    name         varchar(100) null,
    growth       int          null comment '赠送成长值',
    intergration int          null comment '赠送积分',
    type         int(1)       null comment '任务类型：0->新手任务；1->日常任务'
)
    comment '会员任务表';

