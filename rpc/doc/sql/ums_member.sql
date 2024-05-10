create table ums_member
(
    id                     bigint auto_increment
        primary key,
    member_level_id        bigint                             not null comment '会员等级id',
    member_name            varchar(64)                        not null comment '用户名',
    password               varchar(64)                        not null comment '密码',
    nickname               varchar(64)                        not null comment '昵称',
    phone                  varchar(64)                        not null comment '手机号码',
    member_status          tinyint                            not null comment '帐号启用状态:0->禁用；1->启用',
    icon                   varchar(500)                       null comment '头像',
    gender                 tinyint                            null comment '性别：0->未知；1->男；2->女',
    birthday               date                               null comment '生日',
    city                   varchar(64)                        null comment '所做城市',
    job                    varchar(100)                       null comment '职业',
    personalized_signature varchar(200)                       null comment '个性签名',
    source_type            tinyint                            not null comment '用户来源',
    integration            int                                not null comment '积分',
    growth                 int                                not null comment '成长值',
    lottery_count          int                                not null comment '剩余抽奖次数',
    history_integration    int                                not null comment '历史积分数量',
    create_time            datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time            datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_phone
        unique (phone),
    constraint idx_username
        unique (member_name)
)
    comment '会员表';

INSERT INTO ums_member (id, member_level_id, member_name, password, nickname, phone, member_status, create_time, icon,
                        gender, birthday, city, job, personalized_signature, source_type, integration, growth,
                        lottery_count, history_integration)
VALUES (1, 4, 'koobe', '123456', 'koobe', '18613030352', 1, current_time, '', 1, '2021-03-16', '深圳', 'go开发', 'test',
        1, 100, 20, 1000, 1000);
INSERT INTO ums_member (id, member_level_id, member_name, password, nickname, phone, member_status, create_time, icon,
                        gender, birthday, city, job, personalized_signature, source_type, integration, growth,
                        lottery_count, history_integration)
VALUES (2, 1, 'koobe1', '123456', 'koobe1', '18613030351', 0, current_time, '', 0, '2022-06-22', '', '', '', 0, 0, 0, 0,
        0);
