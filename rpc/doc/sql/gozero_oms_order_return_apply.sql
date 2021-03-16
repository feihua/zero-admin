create table oms_order_return_apply
(
    id                 bigint auto_increment
        primary key,
    order_id           bigint         null comment '订单id',
    company_address_id bigint         null comment '收货地址表id',
    product_id         bigint         null comment '退货商品id',
    order_sn           varchar(64)    null comment '订单编号',
    create_time        datetime       null comment '申请时间',
    member_username    varchar(64)    null comment '会员用户名',
    return_amount      decimal(10, 2) null comment '退款金额',
    return_name        varchar(100)   null comment '退货人姓名',
    return_phone       varchar(100)   null comment '退货人电话',
    status             int(1)         null comment '申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝',
    handle_time        datetime       null comment '处理时间',
    product_pic        varchar(500)   null comment '商品图片',
    product_name       varchar(200)   null comment '商品名称',
    product_brand      varchar(200)   null comment '商品品牌',
    product_attr       varchar(500)   null comment '商品销售属性：颜色：红色；尺码：xl;',
    product_count      int            null comment '退货数量',
    product_price      decimal(10, 2) null comment '商品单价',
    product_real_price decimal(10, 2) null comment '商品实际支付单价',
    reason             varchar(200)   null comment '原因',
    description        varchar(500)   null comment '描述',
    proof_pics         varchar(1000)  null comment '凭证图片，以逗号隔开',
    handle_note        varchar(500)   null comment '处理备注',
    handle_man         varchar(100)   null comment '处理人员',
    receive_man        varchar(100)   null comment '收货人',
    receive_time       datetime       null comment '收货时间',
    receive_note       varchar(500)   null comment '收货备注'
)
    comment '订单退货申请';

INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (3, 12, 1, 26, '201809150101000001', '2018-10-17 14:34:57', 'test', 3585.00, '大梨', '18000000000', 0, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg,http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '已经处理了', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (4, 12, 2, 27, '201809150101000001', '2018-10-17 14:40:21', 'test', 3585.98, '大梨', '18000000000', 1, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '', '已经处理了', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (5, 12, 3, 28, '201809150101000001', '2018-10-17 14:44:18', 'test', 3585.98, '大梨', '18000000000', 2, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '', '已经处理了', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (8, 13, 1, 28, '201809150102000002', '2018-10-17 14:44:18', 'test', 3585.00, '大梨', '18000000000', 3, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '', '理由不够充分', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (9, 14, 2, 26, '201809130101000001', '2018-10-17 14:34:57', 'test', 3500.00, '大梨', '18000000000', 2, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '', '呵呵', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (10, 14, 1, 27, '201809130101000001', '2018-10-17 14:40:21', 'test', 3585.00, '大梨', '18000000000', 3, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg', '小米8', '小米', '颜色：黑色;内存：32G', 1, 2699.00, 2022.81, '质量问题', '不够高端', '', '就是不退', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (11, 14, 2, 28, '201809130101000001', '2018-10-17 14:44:18', 'test', 591.05, '大梨', '18000000000', 1, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5a9d248cN071f4959.jpg', '红米5A', '小米', '颜色：金色;内存：16G', 1, 649.00, 591.05, '质量问题', '颜色太土', '', '可以退款', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');
INSERT INTO gozero.oms_order_return_apply (id, order_id, company_address_id, product_id, order_sn, create_time, member_username, return_amount, return_name, return_phone, status, handle_time, product_pic, product_name, product_brand, product_attr, product_count, product_price, product_real_price, reason, description, proof_pics, handle_note, handle_man, receive_man, receive_time, receive_note) VALUES (12, 15, 3, 26, '201809130102000002', '2018-10-17 14:34:57', 'test', 3500.00, '大梨', '18000000000', 2, '2021-03-16 20:56:40', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg', '华为 HUAWEI P20', '华为', '颜色：金色;内存：16G', 1, 3788.00, 3585.98, '质量问题', '老是卡', '', '退货了', 'admin', 'admin', '2021-03-16 20:56:16', '收货了');