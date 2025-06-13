drop table if exists pms_product_spec;
CREATE TABLE `pms_product_spec` (
                                `id`              BIGINT          NOT NULL AUTO_INCREMENT,
                                `category_id`     BIGINT          NOT NULL COMMENT '分类ID',
                                `name`            VARCHAR(50)     NOT NULL COMMENT '规格名称',
                                `sort`            INT             NOT NULL DEFAULT 0 COMMENT '排序',
                                `create_time`     DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
                                `update_time`     DATETIME(3)     NULL ON UPDATE CURRENT_TIMESTAMP(3),
                                `creator_id`      BIGINT          NOT NULL COMMENT '创建人ID',
                                `updater_id`      BIGINT          NULL COMMENT '更新人ID',
                                `is_deleted`      TINYINT(1)      NOT NULL DEFAULT 0 COMMENT '是否删除',
                                PRIMARY KEY (`id`),
                                KEY `idx_category` (`category_id`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品规格表';