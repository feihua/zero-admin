package svc

import (
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	c       config.Config
	DbEngin *gorm.DB

	PmsAlbumModel                            pmsmodel.PmsAlbumModel
	PmsAlbumPicModel                         pmsmodel.PmsAlbumPicModel
	PmsBrandModel                            pmsmodel.PmsBrandModel
	PmsCommentModel                          pmsmodel.PmsCommentModel
	PmsCommentReplayModel                    pmsmodel.PmsCommentReplayModel
	PmsFeightTemplateModel                   pmsmodel.PmsFeightTemplateModel
	PmsMemberPriceModel                      pmsmodel.PmsMemberPriceModel
	PmsProductAttributeCategoryModel         pmsmodel.PmsProductAttributeCategoryModel
	PmsProductAttributeModel                 pmsmodel.PmsProductAttributeModel
	PmsProductAttributeValueModel            pmsmodel.PmsProductAttributeValueModel
	PmsProductCategoryAttributeRelationModel pmsmodel.PmsProductCategoryAttributeRelationModel
	PmsProductCategoryModel                  pmsmodel.PmsProductCategoryModel
	PmsProductFullReductionModel             pmsmodel.PmsProductFullReductionModel
	PmsProductLadderModel                    pmsmodel.PmsProductLadderModel
	PmsProductModel                          pmsmodel.PmsProductModel
	PmsProductOperateLogModel                pmsmodel.PmsProductOperateLogModel
	PmsProductVertifyRecordModel             pmsmodel.PmsProductVertifyRecordModel
	PmsSkuStockModel                         pmsmodel.PmsSkuStockModel
	PmsCollectModel                          pmsmodel.PmsCollectModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "ums_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	// 如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	// 自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	// db.AutoMigrate(&models.UmsMember{})

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:       c,
		DbEngin: db,

		PmsAlbumModel:                            pmsmodel.NewPmsAlbumModel(sqlConn, c.Cache),
		PmsAlbumPicModel:                         pmsmodel.NewPmsAlbumPicModel(sqlConn, c.Cache),
		PmsBrandModel:                            pmsmodel.NewPmsBrandModel(sqlConn, c.Cache),
		PmsCommentModel:                          pmsmodel.NewPmsCommentModel(sqlConn, c.Cache),
		PmsCommentReplayModel:                    pmsmodel.NewPmsCommentReplayModel(sqlConn, c.Cache),
		PmsFeightTemplateModel:                   pmsmodel.NewPmsFeightTemplateModel(sqlConn, c.Cache),
		PmsMemberPriceModel:                      pmsmodel.NewPmsMemberPriceModel(sqlConn, c.Cache),
		PmsProductAttributeCategoryModel:         pmsmodel.NewPmsProductAttributeCategoryModel(sqlConn, c.Cache),
		PmsProductAttributeModel:                 pmsmodel.NewPmsProductAttributeModel(sqlConn, c.Cache),
		PmsProductAttributeValueModel:            pmsmodel.NewPmsProductAttributeValueModel(sqlConn, c.Cache),
		PmsProductCategoryAttributeRelationModel: pmsmodel.NewPmsProductCategoryAttributeRelationModel(sqlConn, c.Cache),
		PmsProductCategoryModel:                  pmsmodel.NewPmsProductCategoryModel(sqlConn, c.Cache),
		PmsProductFullReductionModel:             pmsmodel.NewPmsProductFullReductionModel(sqlConn, c.Cache),
		PmsProductLadderModel:                    pmsmodel.NewPmsProductLadderModel(sqlConn, c.Cache),
		PmsProductModel:                          pmsmodel.NewPmsProductModel(sqlConn, c.Cache),
		PmsProductOperateLogModel:                pmsmodel.NewPmsProductOperateLogModel(sqlConn, c.Cache),
		PmsProductVertifyRecordModel:             pmsmodel.NewPmsProductVertifyRecordModel(sqlConn, c.Cache),
		PmsSkuStockModel:                         pmsmodel.NewPmsSkuStockModel(sqlConn, c.Cache),
	}
}
