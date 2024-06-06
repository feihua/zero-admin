create table pms_brand
(
    id                    bigint auto_increment
        primary key,
    name                  varchar(64)                        not null comment '品牌名称',
    first_letter          varchar(8)                         not null comment '首字母',
    sort                  int                                not null comment '排序',
    factory_status        tinyint                            not null comment '是否为品牌制造商：0->不是；1->是',
    show_status           tinyint                            not null comment '品牌显示状态',
    recommend_status      tinyint  default 0                 not null comment '推荐状态',
    product_count         int                                not null comment '产品数量',
    product_comment_count int                                not null comment '产品评论数量',
    logo                  varchar(255)                       not null comment '品牌logo',
    big_pic               varchar(255)                       not null comment '专区大图',
    brand_story           text                               not null comment '品牌故事',
    create_by             varchar(50)                        not null comment '创建者',
    create_time           datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_by             varchar(50)                        null comment '更新者',
    update_time           datetime                           null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '品牌表';

INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (1, '万和', 'W', 0, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg(5).jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png',
        'Victoria''s Secret的故事', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (2, '三星', 'S', 100, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/57201b47N7bf15715.jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/sanxing_banner_01.png', '三星的故事',
        'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (3, '华为', 'H', 100, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5abf6f26N31658aa2.jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png',
        'Victoria''s Secret的故事', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (4, '格力', 'G', 30, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg (3).jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png',
        'Victoria''s Secret的故事', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (5, '方太', 'F', 20, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg (4).jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png',
        'Victoria''s Secret的故事', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (6, '小米', 'M', 500, 1, 1, 100, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5565f5a2N0b8169ae.jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/xiaomi_banner_01.png', '小米手机的故事',
        'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (21, 'OPPO', 'O', 0, 1, 1, 88, 500,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/timg(6).jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png', 'string', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (49, '七匹狼', 'S', 200, 1, 1, 77, 400,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180518/1522738681.jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/huawei_banner_01.png', 'BOOB的故事',
        'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (50, '海澜之家', 'H', 200, 1, 1, 66, 300,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/5a5c69b9N5d6c5696.jpg',
        'http://img10.360buyimg.com/cms/jfs/t1/133148/4/1605/470028/5edaf5ccEd7a687a9/e0a007631361ff75.jpg',
        '海澜之家的故事', 'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (51, '苹果', 'A', 200, 1, 1, 55, 200,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20200607/49b30bb0377030d1.jpg',
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221108/apple_banner_01.png', '苹果的故事',
        'admin');
INSERT INTO pms_brand (id, name, first_letter, sort, factory_status, show_status, product_count, product_comment_count,
                       logo, big_pic, brand_story, create_by)
VALUES (58, 'NIKE', 'N', 0, 1, 1, 33, 100,
        'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/timg (51).jpg', '', 'NIKE的故事', 'admin');
