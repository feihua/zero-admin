create table ums_member_receive_address
(
    id             bigint auto_increment
        primary key,
    member_id      bigint       null,
    name           varchar(100) null comment '收货人名称',
    phone_number   varchar(64)  null,
    default_status int(1)       null comment '是否为默认',
    post_code      varchar(100) null comment '邮政编码',
    province       varchar(100) null comment '省份/直辖市',
    city           varchar(100) null comment '城市',
    region         varchar(100) null comment '区',
    detail_address varchar(128) null comment '详细地址(街道)'
)
    comment '会员收货地址表';

INSERT INTO gozero.ums_member_receive_address (id, member_id, name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (1, 1, '大梨', '18033441849', 0, '518000', '广东省', '深圳市', '南山区', '科兴科学园');
INSERT INTO gozero.ums_member_receive_address (id, member_id, name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (3, 1, '大梨', '18033441849', 0, '518000', '广东省', '深圳市', '福田区', '清水河街道');
INSERT INTO gozero.ums_member_receive_address (id, member_id, name, phone_number, default_status, post_code, province, city, region, detail_address) VALUES (4, 1, '大梨', '18033441849', 1, '518000', '广东省', '深圳市', '福田区', '东晓街道');