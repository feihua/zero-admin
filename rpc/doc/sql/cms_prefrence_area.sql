create table cms_prefrence_area
(
    id          bigint auto_increment
        primary key,
    name        varchar(255)   null,
    sub_title   varchar(255)   null,
    pic         varbinary(500) null comment '展示图片',
    sort        int            null,
    show_status int(1)         null
)
    comment '优选专区' charset = utf8;

INSERT INTO gozero.cms_prefrence_area (id, name, sub_title, pic, sort, show_status) VALUES (1, '让音质更出众', '音质不打折 完美现场感', null, null, 1);
INSERT INTO gozero.cms_prefrence_area (id, name, sub_title, pic, sort, show_status) VALUES (2, '让音质更出众22', '让音质更出众22', null, null, null);
INSERT INTO gozero.cms_prefrence_area (id, name, sub_title, pic, sort, show_status) VALUES (3, '让音质更出众33', null, null, null, null);
INSERT INTO gozero.cms_prefrence_area (id, name, sub_title, pic, sort, show_status) VALUES (4, '让音质更出众44', null, null, null, null);
