create table cms_subject_category
(
    id            bigint auto_increment
        primary key,
    name          varchar(100) not null,
    icon          varchar(500) not null comment '分类图标',
    subject_count int          not null comment '专题数量',
    show_status tinyint not null,
    sort          int          not null
)
    comment '专题分类表' charset = utf8;

INSERT INTO cms_subject_category (id, name, icon, subject_count, show_status, sort) VALUES (1, '服装专题', ' ', 1, 1, 1);
INSERT INTO cms_subject_category (id, name, icon, subject_count, show_status, sort) VALUES (2, '手机专题', ' ', 1, 1, 1);
