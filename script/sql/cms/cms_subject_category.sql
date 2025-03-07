create table cms_subject_category
(
    id            bigint auto_increment
        primary key comment '主键ID',
    name          varchar(100)                          not null comment '专题分类名称',
    icon          varchar(500)                          not null comment '分类图标',
    subject_count int                                   not null comment '专题数量',
    show_status   tinyint                               not null comment '显示状态：0->不显示；1->显示',
    sort          int                                   not null comment '排序',
    create_by     varchar(50)                           not null comment '创建者',
    create_time   datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by     varchar(50) default ''                not null comment '更新者',
    update_time   datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '专题分类表' charset = utf8;

INSERT INTO cms_subject_category (id, name, icon, subject_count, show_status, sort, create_by)
VALUES (1, '服装专题', ' ', 1, 1, 1, 'admin');
INSERT INTO cms_subject_category (id, name, icon, subject_count, show_status, sort, create_by)
VALUES (2, '手机专题', ' ', 1, 1, 1, 'admin');
