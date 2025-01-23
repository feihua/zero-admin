create table pms_comment
(
    id                bigint auto_increment
        primary key,
    product_id        bigint        not null comment '商品id',
    member_nick_name  varchar(255)  not null comment '评价者昵称',
    product_name      varchar(255)  not null comment '商品名称',
    star              int(3)        not null comment '评价星数：0->5',
    member_ip         varchar(64)   not null comment '评价的ip',
    create_time       datetime      not null comment '评价时间',
    show_status       tinyint       not null comment '是否显示，0-不显示，1-显示',
    product_attribute varchar(255)  not null comment '购买时的商品属性',
    collect_count     int           not null comment '点赞数',
    read_count        int           not null comment '阅读数',
    content           text          not null comment '内容',
    pics              varchar(1000) not null comment '上传图片地址，以逗号隔开',
    member_icon       varchar(255)  not null comment '评论用户头像',
    replay_count      int           not null comment '回复数量'
)
    comment '商品评价表';

