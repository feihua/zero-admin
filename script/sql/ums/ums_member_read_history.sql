create table ums_member_read_history
(
    id                bigint auto_increment
        primary key,
    member_id         bigint                              not null comment '会员id',
    member_nick_name  varchar(50)                         not null comment '会员姓名',
    member_icon       varchar(200)                        not null comment '会员头像',
    product_id        bigint                              not null comment '商品id',
    product_name      varchar(50)                         not null comment '商品名称',
    product_pic       varchar(200)                        not null comment '商品图片',
    product_sub_title varchar(128)                        null comment '商品标题',
    product_price bigint not null comment '商品价格',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '浏览时间'
)
    comment '用户商品浏览历史记录';

