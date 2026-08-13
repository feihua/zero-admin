drop table if exists oms_company_address;
create table oms_company_address
(
    id             bigint auto_increment
        primary key comment '主键ID',
    address_name   varchar(200)                       not null comment '地址名称',
    name           varchar(64)                        not null comment '收发货人姓名',
    phone          varchar(64)                        not null comment '收货人电话',
    province       varchar(64)                        not null comment '省/直辖市',
    city           varchar(64)                        not null comment '市',
    region         varchar(64)                        not null comment '区',
    detail_address varchar(200)                       not null comment '详细地址',
    send_status    tinyint                            not null comment '默认发货地址：0->否；1->是',
    receive_status tinyint                            not null comment '默认收货地址：0->否；1->是',
    create_by      bigint                             not null comment '创建人ID',
    create_time    datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by      bigint                             null comment '更新人ID',
    update_time    datetime                           null on update CURRENT_TIMESTAMP comment '更新时间',
    is_deleted     tinyint  default 0                 not null comment '是否删除'
)
    comment '公司收发货地址表';

-- 插入公司收发货地址数据
insert into oms_company_address (id, address_name, name, phone, province, city, region, detail_address, send_status, receive_status, create_by, is_deleted)
values (1, '总部地址', '张三', '13800138000', '北京市', '北京市', '朝阳区', '建国路88号', 1, 0, 1, 0),
       (2, '分公司地址', '李四', '13900139000', '上海市', '上海市', '浦东新区', '世纪大道100号', 0, 1, 2, 0),
       (3, '仓库地址', '王五', '13700137000', '广东省', '广州市', '天河区', '体育西路200号', 1, 1, 3, 0),
       (4, '配送中心', '赵六', '13600136000', '浙江省', '杭州市', '西湖区', '文三路300号', 0, 0, 4, 0),
       (5, '备用地址', '孙七', '13500135000', '江苏省', '南京市', '玄武区', '中山路400号', 0, 0, 5, 0);