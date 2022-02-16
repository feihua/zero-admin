package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/internal/config"
)

type ServiceContext struct {
	c config.Config

	PmsAlbumModel                            *pmsmodel.PmsAlbumModel
	PmsAlbumPicModel                         *pmsmodel.PmsAlbumPicModel
	PmsBrandModel                            *pmsmodel.PmsBrandModel
	PmsCommentModel                          *pmsmodel.PmsCommentModel
	PmsCommentReplayModel                    *pmsmodel.PmsCommentReplayModel
	PmsFeightTemplateModel                   *pmsmodel.PmsFeightTemplateModel
	PmsMemberPriceModel                      *pmsmodel.PmsMemberPriceModel
	PmsProductAttributeCategoryModel         *pmsmodel.PmsProductAttributeCategoryModel
	PmsProductAttributeModel                 *pmsmodel.PmsProductAttributeModel
	PmsProductAttributeValueModel            *pmsmodel.PmsProductAttributeValueModel
	PmsProductCategoryAttributeRelationModel *pmsmodel.PmsProductCategoryAttributeRelationModel
	PmsProductCategoryModel                  *pmsmodel.PmsProductCategoryModel
	PmsProductFullReductionModel             *pmsmodel.PmsProductFullReductionModel
	PmsProductLadderModel                    *pmsmodel.PmsProductLadderModel
	PmsProductModel                          *pmsmodel.PmsProductModel
	PmsProductOperateLogModel                *pmsmodel.PmsProductOperateLogModel
	PmsProductVertifyRecordModel             *pmsmodel.PmsProductVertifyRecordModel
	PmsSkuStockModel                         *pmsmodel.PmsSkuStockModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		PmsAlbumModel:                            pmsmodel.NewPmsAlbumModel(sqlConn),
		PmsAlbumPicModel:                         pmsmodel.NewPmsAlbumPicModel(sqlConn),
		PmsBrandModel:                            pmsmodel.NewPmsBrandModel(sqlConn),
		PmsCommentModel:                          pmsmodel.NewPmsCommentModel(sqlConn),
		PmsCommentReplayModel:                    pmsmodel.NewPmsCommentReplayModel(sqlConn),
		PmsFeightTemplateModel:                   pmsmodel.NewPmsFeightTemplateModel(sqlConn),
		PmsMemberPriceModel:                      pmsmodel.NewPmsMemberPriceModel(sqlConn),
		PmsProductAttributeCategoryModel:         pmsmodel.NewPmsProductAttributeCategoryModel(sqlConn),
		PmsProductAttributeModel:                 pmsmodel.NewPmsProductAttributeModel(sqlConn),
		PmsProductAttributeValueModel:            pmsmodel.NewPmsProductAttributeValueModel(sqlConn),
		PmsProductCategoryAttributeRelationModel: pmsmodel.NewPmsProductCategoryAttributeRelationModel(sqlConn),
		PmsProductCategoryModel:                  pmsmodel.NewPmsProductCategoryModel(sqlConn),
		PmsProductFullReductionModel:             pmsmodel.NewPmsProductFullReductionModel(sqlConn),
		PmsProductLadderModel:                    pmsmodel.NewPmsProductLadderModel(sqlConn),
		PmsProductModel:                          pmsmodel.NewPmsProductModel(sqlConn),
		PmsProductOperateLogModel:                pmsmodel.NewPmsProductOperateLogModel(sqlConn),
		PmsProductVertifyRecordModel:             pmsmodel.NewPmsProductVertifyRecordModel(sqlConn),
		PmsSkuStockModel:                         pmsmodel.NewPmsSkuStockModel(sqlConn),
	}
}
