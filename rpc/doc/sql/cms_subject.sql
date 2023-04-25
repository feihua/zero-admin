create table cms_subject
(
    id               bigint auto_increment
        primary key,
    category_id      bigint        null,
    title            varchar(100)  null,
    pic              varchar(500)  null comment '专题主图',
    product_count    int           null comment '关联产品数量',
    recommend_status int(1)        null,
    create_time      datetime      null,
    collect_count    int           null,
    read_count       int           null,
    comment_count    int           null,
    album_pics       varchar(1000) null comment '画册图片用逗号分割',
    description      varchar(1000) null,
    show_status      int(1)        null comment '显示状态：0->不显示；1->显示',
    content          text          null,
    forward_count    int           null comment '转发数',
    category_name    varchar(200)  null comment '专题分类名称'
)
    comment '专题表' charset = utf8;

INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (1, 1, 'polo衬衫的也时尚', null, null, null, '2018-11-11 13:26:55', null, null, null, null, null, null, null, null, '服装专题');
INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (2, 2, '大牌手机低价秒', null, null, null, '2018-11-12 13:27:00', null, null, null, null, null, null, null, null, '手机专题');
INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (3, 2, '晓龙845新品上市', null, null, null, '2018-11-13 13:27:05', null, null, null, null, null, null, null, null, '手机专题');
INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (4, 1, '夏天应该穿什么', null, null, null, '2018-11-01 13:27:09', null, null, null, null, null, null, null, null, '服装专题');
INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (5, 1, '夏季精选', null, null, null, '2018-11-06 13:27:18', null, null, null, null, null, null, null, null, '服装专题');
INSERT INTO gozero.cms_subject (id, category_id, title, pic, product_count, recommend_status, create_time, collect_count, read_count, comment_count, album_pics, description, show_status, content, forward_count, category_name) VALUES (6, 2, '品牌手机降价', null, null, null, '2018-11-07 13:27:21', null, null, null, null, null, null, null, null, '手机专题');
