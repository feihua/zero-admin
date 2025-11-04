drop table if exists oms_order_return;
create table oms_order_return
(
    id              bigint auto_increment comment '主键ID'
        primary key,
    order_id        bigint                                   not null comment '关联订单ID',
    return_no       varchar(64)                              not null comment '退货单号',
    member_id       bigint                                   not null comment '会员ID',
    status          tinyint        default 0                 not null comment '退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）',
    type            tinyint        default 0                 not null comment '售后类型（0退货退款 1仅退款 2换货）',
    reason          varchar(255)   default ''                not null comment '退货原因',
    description     varchar(500)   default ''                not null comment '问题描述',
    proof_pic       varchar(1000)  default ''                not null comment '凭证图片，逗号分隔',
    refund_amount   decimal(10, 2) default 0.00              not null comment '退款金额',
    return_name     varchar(64)                              not null comment '退货人姓名',
    return_phone    varchar(32)                              not null comment '退货人电话',
    company_address varchar(255)                             not null comment '退货收货地址',
    create_time     datetime       default CURRENT_TIMESTAMP not null comment '申请时间',
    handle_time     datetime                                 null comment '处理时间',
    handle_note     varchar(500)   default ''                not null comment '处理备注',
    handle_man      varchar(100)   default ''                not null comment '处理人员',
    receive_time    datetime                                 null comment '收货时间',
    receive_note    varchar(500)   default ''                not null comment '收货备注',
    receive_man     varchar(100)   default ''                not null comment '收货人',
    refund_time     datetime                                 null comment '退款时间',
    close_time      datetime                                 null comment '关闭时间',
    remark          varchar(255)   default ''                not null comment '备注'
)
    comment '退货/售后主表';

create index idx_member_id
    on oms_order_return (member_id);

create index idx_order_id
    on oms_order_return (order_id);

create index idx_return_no
    on oms_order_return (return_no);

INSERT INTO oms_order_return (
    order_id, return_no, member_id, status, type, reason, description,
    proof_pic, refund_amount, return_name, return_phone, company_address,
    create_time, handle_time, handle_note, handle_man, receive_time,
    receive_note, receive_man, refund_time, close_time, remark
) VALUES
-- Pending review return request
(1001, 'RTN20231201001', 10001, 0, 0, '商品质量问题', '收到的商品有明显划痕和损坏',
 'pic1.jpg,pic2.jpg', 299.00, '张三', '13800138001', '北京市朝阳区某某街道101号',
 '2023-12-01 10:30:00', NULL, '', '', NULL,
 '', '', NULL, NULL, '客户 urgent'),

-- Approved return
(1002, 'RTN20231201002', 10002, 1, 0, '尺寸不合适', '买大了，需要换小一号',
 'pic3.jpg', 159.90, '李四', '13800138002', '上海市浦东新区某某路202号',
 '2023-12-01 14:20:00', '2023-12-01 15:00:00', '已审核通过', '客服小王', NULL,
 '', '', NULL, NULL, '普通客户'),

-- Received return
(1003, 'RTN20231202001', 10003, 2, 1, '发错商品', '收到的不是订购的商品',
 '', 89.50, '王五', '13800138003', '广州市天河区某某大道303号',
 '2023-12-02 09:15:00', '2023-12-02 10:00:00', '审核通过', '客服小李', '2023-12-03 14:30:00',
 '已收到退货商品', '仓库小赵', NULL, NULL, 'VIP客户'),

-- Refunded return
(1004, 'RTN20231203001', 10004, 3, 0, '商品不喜欢', '颜色与描述不符',
 'pic4.jpg,pic5.jpg', 199.99, '赵六', '13800138004', '深圳市南山区某某科技园404号',
 '2023-12-03 16:45:00', '2023-12-03 17:00:00', '同意退款', '客服小陈', '2023-12-04 09:00:00',
 '商品已验收', '仓库小孙', '2023-12-04 10:30:00', NULL, '需加快处理'),

-- Rejected return
(1005, 'RTN20231204001', 10005, 4, 2, '人为损坏', '商品有明显人为使用痕迹',
 'pic6.jpg', 0.00, '钱七', '13800138005', '杭州市西湖区某某广场505号',
 '2023-12-04 11:20:00', '2023-12-04 13:00:00', '不符合退货条件，已拒绝', '客服小周', NULL,
 '', '', NULL, NULL, '注意客户态度'),

-- Closed return
(1006, 'RTN20231205001', 10006, 5, 1, '其他原因', '客户主动取消退货申请',
 '', 0.00, '孙八', '13800138006', '南京市鼓楼区某某大厦606号',
 '2023-12-05 15:30:00', NULL, '', '', NULL,
 '', '', NULL, '2023-12-05 16:00:00', '客户自行解决');
