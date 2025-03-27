drop table if exists ums_member_receive_address;
create table ums_member_receive_address
(
    id             bigint auto_increment
        primary key,
    member_id      bigint                             not null comment '会员id',
    member_name    varchar(100)                       not null comment '收货人名称',
    phone_number   varchar(64)                        not null comment '收货人电话',
    default_status tinyint                            not null comment '是否为默认',
    post_code      varchar(100)                       not null comment '邮政编码',
    province       varchar(100)                       not null comment '省份/直辖市',
    city           varchar(100)                       not null comment '城市',
    region         varchar(100)                       not null comment '区',
    detail_address varchar(128)                       not null comment '详细地址(街道)',
    create_time    datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time    datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '会员收货地址';

INSERT INTO ums_member_receive_address (id, member_id, member_name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (1, 1, 'test', '18613030352', 1, '518000', '广东省', '深圳市', '南山区', '西丽街道1688街道88号');
INSERT INTO ums_member_receive_address (id, member_id, member_name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (2, 1, 'test', '18613030111', 0, '318000', '广东省', '茂名市', '高州', '曹江林村123号');
INSERT INTO ums_member_receive_address (id, member_id, member_name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (3, 2, 'koobe', '13144410811', 1, '226600', '广东省', '茂名市', '高州', '测试详细地址');
INSERT INTO ums_member_receive_address (id, member_id, member_name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (4, 2, 'koobe', '13144410222', 0, '647000', '广东省', '广州市', '天河区', '测试详细地址2');
