drop table if exists oms_order_payment;
create table oms_order_payment
(
    id             bigint auto_increment
        primary key comment '主键ID',
    order_id       bigint                                not null comment '订单ID',
    order_no       varchar(32)                           not null comment '订单编号',
    pay_type       tinyint                               not null comment '支付方式：1-支付宝，2-微信，3-银联',
    transaction_id varchar(50) default ''                not null comment '支付流水号',
    total_amount   decimal(10, 2)                        not null comment '订单金额',
    pay_amount     decimal(10, 2)                        not null comment '支付金额',
    pay_status     tinyint     default 0                 not null comment '支付状态：0-待支付，1-支付成功，2-支付失败',
    pay_time       datetime                              null comment '支付时间',
    create_time    datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time    datetime                              null on update CURRENT_TIMESTAMP,
    is_deleted     tinyint     default 0                 not null comment '是否删除',
    constraint uk_order
        unique (order_id, is_deleted)
)
    comment '订单支付表';

create index idx_transaction
    on oms_order_payment (transaction_id, is_deleted);

