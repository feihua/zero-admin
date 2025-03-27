drop table if exists ums_member;
create table ums_member
(
    id                     bigint auto_increment
        primary key,
    member_level_id        bigint                                 not null comment '会员等级id',
    member_name            varchar(64)                            not null comment '用户名',
    password               varchar(64)                            not null comment '密码',
    nickname               varchar(64)                            not null comment '昵称',
    phone                  varchar(64)                            not null comment '手机号码',
    member_status          tinyint                                not null comment '帐号启用状态:0->禁用；1->启用',
    icon                   varchar(500) default ''                not null comment '头像',
    gender                 tinyint      default 0                 not null comment '性别：0->未知；1->男；2->女',
    birthday               date                                   not null comment '生日',
    city                   varchar(64)  default ''                not null comment '所做城市',
    job                    varchar(100) default ''                not null comment '职业',
    personalized_signature varchar(200) default ''                not null comment '个性签名',
    source_type            tinyint                                not null comment '用户来源：1->移动端; 2->pc端',
    integration            int                                    not null comment '积分',
    growth                 int                                    not null comment '成长值',
    lottery_count          int                                    not null comment '剩余抽奖次数',
    history_integration    int                                    not null comment '历史积分数量',
    create_time            datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time            datetime                               null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_phone
        unique (phone),
    constraint idx_username
        unique (member_name)
)
    comment '会员信息';

INSERT INTO ums_member (id, member_level_id, member_name, password, nickname, phone, member_status, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, lottery_count, history_integration) VALUES (1, 4, 'test', '123456', '测试会员1', '18613030352', 1, '', 1, current_date, '深圳', '自由职业', '个性签名测试', 1, 100, 20, 10, 1000);
INSERT INTO ums_member (id, member_level_id, member_name, password, nickname, phone, member_status, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, lottery_count, history_integration) VALUES (2, 4, 'koobe', '123456', '测试会员2', '13144410811', 1, '', 1, current_date, '广州', '失业中', '不想写个性签名', 2, 80, 120, 7, 800);
