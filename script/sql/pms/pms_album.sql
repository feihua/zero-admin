create table pms_album
(
    id          bigint auto_increment
        primary key,
    name        varchar(64)   not null,
    cover_pic   varchar(1000) not null,
    pic_count   int           not null,
    sort        int           not null,
    description varchar(1000) not null
)
    comment '相册表';

