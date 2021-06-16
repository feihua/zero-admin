create table pay_wx_record
(
    id          int auto_increment comment '主键',
    businessId  varchar(32)   not null comment '业务编号',
    amount      varchar(32)   not null comment '金额',
    pay_type    varchar(12)   not null comment '支付类型(APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付)',
    remarks     varchar(128)  not null comment '备注',
    return_code varchar(12)   not null,
    return_msg  varchar(12)   not null,
    result_code varchar(12)   not null,
    result_msg  varchar(12)   not null,
    pay_status  int default 0 not null comment '0：初始化 1：已发送 2：成功 3：失败',
    create_time date          not null comment '创建时间',
    update_time date          not null comment '更新时间',
    constraint pay_record_id_uindex
        unique (id)
)
    comment '微信支付记录';

alter table pay_wx_record
    add primary key (id);

