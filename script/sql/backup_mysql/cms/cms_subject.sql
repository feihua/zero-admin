drop table if exists cms_subject;
create table cms_subject
(
    id               bigint auto_increment
        primary key comment '专题id',
    category_id      bigint                                not null comment '专题分类id',
    title            varchar(100)                          not null comment '专题标题',
    pic              varchar(500)                          not null comment '专题主图',
    product_count    int                                   not null comment '关联产品数量',
    recommend_status tinyint                               not null comment '推荐状态：0->不推荐；1->推荐',
    collect_count    int                                   not null comment '收藏数',
    read_count       int                                   not null comment '阅读数',
    comment_count    int                                   not null comment '评论数',
    album_pics       varchar(1000)                         not null comment '画册图片用逗号分割',
    description      varchar(1000)                         not null comment '专题内容',
    show_status      tinyint                               not null comment '显示状态：0->不显示；1->显示',
    content          text                                  not null comment '专题内容',
    forward_count    int                                   not null comment '转发数',
    category_name    varchar(200)                          not null comment '专题分类名称',
    sort             int         default 1                 not null comment '排序',
    create_by        varchar(50)                           not null comment '创建者',
    create_time      datetime    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by        varchar(50) default ''                not null comment '更新者',
    update_time      datetime                              null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '专题表' charset = utf8;

INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (1, 1, 'polo衬衫的也时尚', '1', 1, 1, '2018-11-11 13:26:55', 1, 1, 1, '1', ' ', 1, '1', 1, '服装专题', 'admin');
INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (2, 2, '大牌手机低价秒', '1', 1, 1, '2018-11-12 13:27:00', 1, 1, 1, '1', ' ', 1, '1', 1, '手机专题', 'admin');
INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (3, 2, '晓龙845新品上市', '1', 1, 1, '2018-11-13 13:27:05', 1, 1, 1, '1', ' ', 1, '1', 1, '手机专题', 'admin');
INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (4, 1, '夏天应该穿什么', '1', 1, 1, '2018-11-01 13:27:09', 1, 1, 1, '1', ' ', 1, '1', 1, '服装专题', 'admin');
INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (5, 1, '夏季精选', '1', 1, 1, '2018-11-06 13:27:18', 1, 1, 1, '1', ' ', 1, '1', 1, '服装专题', 'admin');
INSERT INTO cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count,
                         read_count, comment_count, album_pics, description, show_status, content, forward_count,
                         category_name, create_by)
VALUES (6, 2, '品牌手机降价', '1', 1, 1, '2018-11-07 13:27:21', 1, 1, 1, '1', ' ', 1, '1', 1, '手机专题', 'admin');
