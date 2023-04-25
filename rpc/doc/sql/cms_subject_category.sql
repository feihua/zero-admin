create table cms_subject_category
(
    id            bigint auto_increment
        primary key,
    name          varchar(100) null,
    icon          varchar(500) null comment '分类图标',
    subject_count int          null comment '专题数量',
    show_status   int(2)       null,
    sort          int          null
)
    comment '专题分类表' charset = utf8;

INSERT INTO gozero.cms_subject_category (id, name, icon, subject_count, show_status, sort) VALUES (1, '服装专题', null, null, null, null);
INSERT INTO gozero.cms_subject_category (id, name, icon, subject_count, show_status, sort) VALUES (2, '手机专题', null, null, null, null);
