drop table if exists ums_member_address;

create table ums_member_address
(
    id             bigint auto_increment
        primary key comment '主键ID',
    member_id      bigint                                not null comment '会员ID',
    receiver_name  varchar(50)                           not null comment '收货人姓名',
    receiver_phone varchar(20)                           not null comment '收货人电话',
    province       varchar(50)                           not null comment '省份',
    city           varchar(50)                           not null comment '城市',
    district       varchar(50)                           not null comment '区县',
    detail_address varchar(200)                          not null comment '详细地址',
    postal_code    varchar(20) default ''                not null comment '邮政编码',
    tag            varchar(20) default ''                not null comment '地址标签：家、公司等',
    is_default     tinyint     default 0                 not null comment '是否默认地址',
    create_time    datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time    datetime                              null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint     default 0                 not null comment '是否删除'
)
    comment '会员收货地址表';

create index idx_default
    on ums_member_address (member_id, is_default, is_deleted);

create index idx_member
    on ums_member_address (member_id, is_deleted);


-- 为会员添加收货地址数据
insert into ums_member_address (id, member_id, receiver_name, receiver_phone, province, city, district, detail_address, postal_code, tag, is_default)
values (1, 1001, '张三', '13800138001', '广东省', '深圳市', '南山区', '科技园科兴科学园B座', '518057', '公司', 0),
       (2, 1001, '张三', '13800138001', '广东省', '深圳市', '福田区', '福中三路1006号诺德中心', '518048', '家', 1),
       (3, 1002, '李四', '13800138002', '广东省', '广州市', '天河区', '珠江新城花城大道85号高德置地广场', '510623', '公司', 1),
       (4, 1002, '李四妈妈', '13800138003', '广东省', '广州市', '越秀区', '解放北路928号', '510030', '父母家', 0),
       (5, 1003, '王五', '13800138004', '广东省', '珠海市', '香洲区', '情侣南路399号', '519000', '家', 1),
       (6, 1003, '王五', '13800138004', '广东省', '珠海市', '横琴新区', '环岛东路2000号', '519031', '公司', 0),
       (7, 1004, '赵六', '13800138005', '广东省', '东莞市', '南城区', '鸿福路200号第一国际', '523000', '公司', 1),
       (8, 1005, '钱七', '13800138006', '广东省', '深圳市', '宝安区', '新安街道创业二路腾讯滨海大厦', '518101', '公司', 1),
       (9, 1005, '钱七', '13800138006', '广东省', '深圳市', '罗湖区', '东门南路金光华广场', '518001', '家', 0);