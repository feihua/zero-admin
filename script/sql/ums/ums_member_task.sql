create table ums_member_task
(
    id            bigint auto_increment
        primary key,
    task_name     varchar(100)                          not null comment '任务名称',
    task_growth   int                                   not null comment '赠送成长值',
    task_integral int                                   not null comment '赠送积分',
    task_type     tinyint                               not null comment '任务类型：0->新手任务；1->日常任务',
    create_by     varchar(50)                           not null comment '创建者',
    create_time   datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     varchar(50) default ''                not null comment '更新者',
    update_time   datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '会员任务表';
