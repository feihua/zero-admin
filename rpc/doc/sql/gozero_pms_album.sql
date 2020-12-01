create table pms_album
(
    id          bigint auto_increment
        primary key,
    name        varchar(64)   null,
    cover_pic   varchar(1000) null,
    pic_count   int           null,
    sort        int           null,
    description varchar(1000) null
)
    comment '相册表';

