create table oms_company_address
(
    id             bigint auto_increment
        primary key,
    address_name   varchar(200) null comment '地址名称',
    send_status    int(1)       null comment '默认发货地址：0->否；1->是',
    receive_status int(1)       null comment '是否默认收货地址：0->否；1->是',
    name           varchar(64)  null comment '收发货人姓名',
    phone          varchar(64)  null comment '收货人电话',
    province       varchar(64)  null comment '省/直辖市',
    city           varchar(64)  null comment '市',
    region         varchar(64)  null comment '区',
    detail_address varchar(200) null comment '详细地址'
)
    comment '公司收发货地址表';

INSERT INTO gozero.oms_company_address (id, address_name, send_status, receive_status, name, phone, province, city, region, detail_address) VALUES (1, '深圳发货点', 1, 1, '大梨', '18000000000', '广东省', '深圳市', '南山区', '科兴科学园');
INSERT INTO gozero.oms_company_address (id, address_name, send_status, receive_status, name, phone, province, city, region, detail_address) VALUES (2, '北京发货点', 0, 0, '大梨', '18000000000', '北京市', ' ', '南山区', '科兴科学园');
INSERT INTO gozero.oms_company_address (id, address_name, send_status, receive_status, name, phone, province, city, region, detail_address) VALUES (3, '南京发货点', 0, 0, '大梨', '18000000000', '江苏省', '南京市', '南山区', '科兴科学园');