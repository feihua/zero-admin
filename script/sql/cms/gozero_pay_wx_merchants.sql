create table pay_wx_merchants
(
    id          int auto_increment,
    mer_id      varchar(24)  not null comment '本地系统商户Id,分配给调用方',
    app_id      varchar(24)  not null comment '应用ID 微信开放平台审核通过的应用APPID',
    mch_id      varchar(128) not null comment '微信支付分配的商户号',
    mch_key     varchar(128) not null comment '微信支付分配的商户秘钥',
    pay_type    varchar(12)  not null comment 'APP:APP支付 JSAPI:小程序,公众号 MWEB:H5支付',
    notify_url  varchar(128) not null comment '微信支付回调地址',
    remarks     varchar(64)  not null comment '备注',
    create_time date         not null,
    update_time date         not null,
    constraint pay_wx_merchants_id_uindex
        unique (id)
)
    comment '微信商户信息';

alter table pay_wx_merchants
    add primary key (id);

