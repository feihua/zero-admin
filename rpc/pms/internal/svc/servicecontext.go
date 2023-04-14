package svc

import (
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c config.Config

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

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

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
