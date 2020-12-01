create table ums_growth_change_history
(
    id           bigint auto_increment
        primary key,
    member_id    bigint       null,
    create_time  datetime     null,
    change_type  int(1)       null comment '改变类型：0->增加；1->减少',
    change_count int          null comment '积分改变数量',
    operate_man  varchar(100) null comment '操作人员',
    operate_note varchar(200) null comment '操作备注',
    source_type  int(1)       null comment '积分来源：0->购物；1->管理员修改'
)
    comment '成长值变化历史记录表';

INSERT INTO gozero.ums_growth_change_history (id, member_id, create_time, change_type, change_count, operate_man, operate_note, source_type) VALUES (1, 1, '2018-08-29 17:16:35', 0, 1000, 'test', '测试使用', 1);