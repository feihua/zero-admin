create table pms_album_pic
(
    id       bigint auto_increment
        primary key,
    album_id bigint        null,
    pic      varchar(1000) null
)
    comment '画册图片表';

