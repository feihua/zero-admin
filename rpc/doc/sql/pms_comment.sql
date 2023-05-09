create table pms_comment
(
    id                bigint auto_increment
        primary key,
    product_id        bigint        not null,
    member_nick_name  varchar(255)  not null,
    product_name      varchar(255)  not null,
    star              int(3)        not null comment '评价星数：0->5',
    member_ip         varchar(64)   not null comment '评价的ip',
    create_time       datetime      not null,
    show_status       int(1)        not null,
    product_attribute varchar(255)  not null comment '购买时的商品属性',
    collect_couont    int           not null,
    read_count        int           not null,
    content           text          not null,
    pics              varchar(1000) not null comment '上传图片地址，以逗号隔开',
    member_icon       varchar(255)  not null comment '评论用户头像',
    replay_count      int           not null
)
    comment '商品评价表';

