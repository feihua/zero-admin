create table cms_preferred_area
(
    id          bigint auto_increment
        primary key,
    name        varchar(255)                       not null,
    sub_title   varchar(255)                       not null,
    pic         varchar(500)                       not null comment '展示图片',
    sort        int                                not null,
    show_status tinyint                            not null,
    create_by   varchar(50)                        not null comment '创建者',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by   varchar(50)                        null comment '更新者',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '优选专区' charset = utf8;

INSERT INTO cms_preferred_area (id, name, sub_title, pic, sort, show_status, create_by)
VALUES (1, '让音质更出众', '音质不打折 完美现场感', '', 1, 1, 'admin');
INSERT INTO cms_preferred_area (id, name, sub_title, pic, sort, show_status, create_by)
VALUES (2, '让音质更出众22', '让音质更出众22', '', 2, 1, 'admin');
INSERT INTO cms_preferred_area (id, name, sub_title, pic, sort, show_status, create_by)
VALUES (3, '让音质更出众33', ' ', '', 3, 1, 'admin');
INSERT INTO cms_preferred_area (id, name, sub_title, pic, sort, show_status, create_by)
VALUES (4, '让音质更出众44', ' ', '', 4, 1, 'admin');
