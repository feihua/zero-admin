create table ums_member
(
    id                     bigint auto_increment
        primary key,
    member_level_id        bigint       null,
    username               varchar(64)  null comment '用户名',
    password               varchar(64)  null comment '密码',
    nickname               varchar(64)  null comment '昵称',
    phone                  varchar(64)  null comment '手机号码',
    status                 int(1)       null comment '帐号启用状态:0->禁用；1->启用',
    create_time            datetime     null comment '注册时间',
    icon                   varchar(500) null comment '头像',
    gender                 int(1)       null comment '性别：0->未知；1->男；2->女',
    birthday               date         null comment '生日',
    city                   varchar(64)  null comment '所做城市',
    job                    varchar(100) null comment '职业',
    personalized_signature varchar(200) null comment '个性签名',
    source_type            int(1)       null comment '用户来源',
    integration            int          null comment '积分',
    growth                 int          null comment '成长值',
    luckey_count           int          null comment '剩余抽奖次数',
    history_integration    int          null comment '历史积分数量',
    constraint idx_phone
        unique (phone),
    constraint idx_username
        unique (username)
)
    comment '会员表';

INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (1, 4, 'test', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'windir', '18061581849', 1, '2018-08-02 10:35:44', null, 1, '2009-06-01', '上海', '学生', 'test', null, 5000, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (3, 4, 'windy', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'windy', '18061581848', 1, '2018-08-03 16:46:38', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (4, 4, 'zhengsan', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'zhengsan', '18061581847', 1, '2018-11-12 14:12:04', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (5, 4, 'lisi', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'lisi', '18061581841', 1, '2018-11-12 14:12:38', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (6, 4, 'wangwu', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'wangwu', '18061581842', 1, '2018-11-12 14:13:09', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (7, 4, 'lion', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'lion', '18061581845', 1, '2018-11-12 14:21:39', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (8, 4, 'shari', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'shari', '18061581844', 1, '2018-11-12 14:22:00', null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO gozero.ums_member (id, member_level_id, username, password, nickname, phone, status, create_time, icon, gender, birthday, city, job, personalized_signature, source_type, integration, growth, luckey_count, history_integration) VALUES (9, 4, 'aewen', '$2a$10$NZ5o7r2E.ayT2ZoxgjlI.eJ6OEYqjH7INR/F.mXDbjZJi9HF0YCVG', 'aewen', '18061581843', 1, '2018-11-12 14:22:55', null, null, null, null, null, null, null, null, null, null, null);