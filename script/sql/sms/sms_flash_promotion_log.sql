create table sms_flash_promotion_log
(
    id        bigint auto_increment comment '编号'
        primary key,
    member_id bigint not null comment '会员id',
    product_id     bigint       not null comment '商品id',
    member_phone   varchar(64)  not null comment '会员电话',
    product_name   varchar(100) not null comment '商品名称',
    subscribe_time datetime     not null comment '会员订阅时间',
    send_time      datetime     not null comment '发送时间'
)
    comment '限时购通知记录';

