create table ums_integration_change_history
(
    id           bigint auto_increment
        primary key,
    member_id    bigint                                 not null comment '会员id',
    change_type  tinyint                                not null comment '改变类型：0->增加；1->减少',
    change_count int                                    not null comment '积分改变数量',
    operate_man  varchar(100)                           not null comment '操作人员',
    operate_note varchar(200) default ''                not null comment '操作备注',
    source_type  tinyint                                not null comment '积分来源：0->购物；1->管理员修改',
    create_time  datetime     default CURRENT_TIMESTAMP not null comment '创建时间'
)
    comment '积分变化历史记录表';

