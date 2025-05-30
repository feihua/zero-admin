drop table if exists ums_member_brand_attention;
create table ums_member_brand_attention
(
    id               bigint auto_increment
        primary key,
    member_id        bigint                                 not null comment '会员id',
    member_nick_name varchar(50)                            not null comment '会员姓名',
    member_icon      varchar(200)                           not null comment '会员头像',
    brand_id         bigint                                 not null comment '品牌id',
    brand_name       varchar(50)                            not null comment '品牌名称',
    brand_logo       varchar(200)                           not null comment '品牌Logo',
    brand_city       varchar(128) default ''                not null comment '品牌所在城市',
    create_time      timestamp    default CURRENT_TIMESTAMP not null comment '关注时间'
) comment '会员关注品牌管理';

INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (1, 'test', '', 6, '小米', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5565f5a2N0b8169ae.jpg', '广东省');
INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (1, 'test', '', 4, '格力', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg (3).jpg', '广东省');
INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (1, 'test', '', 3, '华为', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5abf6f26N31658aa2.jpg', '广东省');
INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (2, 'test', '', 6, '小米', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5565f5a2N0b8169ae.jpg', '广东省');
INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (2, 'test', '', 4, '格力', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg (3).jpg', '广东省');
INSERT INTO ums_member_brand_attention (member_id, member_nick_name, member_icon, brand_id, brand_name, brand_logo, brand_city) VALUES (2, 'test', '', 3, '华为', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5abf6f26N31658aa2.jpg', '广东省');
