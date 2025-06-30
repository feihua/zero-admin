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

