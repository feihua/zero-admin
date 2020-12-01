create table pms_comment
(
    id                bigint auto_increment
        primary key,
    product_id        bigint        null,
    member_nick_name  varchar(255)  null,
    product_name      varchar(255)  null,
    star              int(3)        null comment '评价星数：0->5',
    member_ip         varchar(64)   null comment '评价的ip',
    create_time       datetime      null,
    show_status       int(1)        null,
    product_attribute varchar(255)  null comment '购买时的商品属性',
    collect_couont    int           null,
    read_count        int           null,
    content           text          null,
    pics              varchar(1000) null comment '上传图片地址，以逗号隔开',
    member_icon       varchar(255)  null comment '评论用户头像',
    replay_count      int           null
)
    comment '商品评价表';

