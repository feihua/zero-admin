package svc

import (
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	c  config.Config
	DB *gorm.DB

	//PmsAlbumModel                            pmsmodel.PmsAlbumModel
	//PmsAlbumPicModel                         pmsmodel.PmsAlbumPicModel
	//PmsBrandModel                            pmsmodel.PmsBrandModel
	//PmsCommentModel                          pmsmodel.PmsCommentModel
	//PmsCommentReplayModel                    pmsmodel.PmsCommentReplayModel
	//PmsFeightTemplateModel                   pmsmodel.PmsFeightTemplateModel
	//PmsMemberPriceModel                      pmsmodel.PmsMemberPriceModel
	//PmsProductAttributeCategoryModel         pmsmodel.PmsProductAttributeCategoryModel
	//PmsProductAttributeModel                 pmsmodel.PmsProductAttributeModel
	//PmsProductAttributeValueModel            pmsmodel.PmsProductAttributeValueModel
	//PmsProductCategoryAttributeRelationModel pmsmodel.PmsProductCategoryAttributeRelationModel
	//PmsProductCategoryModel                  pmsmodel.PmsProductCategoryModel
	//PmsProductFullReductionModel             pmsmodel.PmsProductFullReductionModel
	//PmsProductLadderModel                    pmsmodel.PmsProductLadderModel
	//PmsProductModel                          pmsmodel.PmsProductModel
	//PmsProductOperateLogModel                pmsmodel.PmsProductOperateLogModel
	//PmsProductVertifyRecordModel             pmsmodel.PmsProductVertifyRecordModel
	//PmsSkuStockModel                         pmsmodel.PmsSkuStockModel
	//PmsCollectModel                          pmsmodel.PmsProductCollectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}

	logx.Debug("mysql已连接")
	query.SetDefault(DB)

	//sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:  c,
		DB: DB,

		//PmsAlbumModel:                            pmsmodel.NewPmsAlbumModel(sqlConn),
		//PmsAlbumPicModel:                         pmsmodel.NewPmsAlbumPicModel(sqlConn),
		//PmsBrandModel:                            pmsmodel.NewPmsBrandModel(sqlConn),
		//PmsCommentModel:                          pmsmodel.NewPmsCommentModel(sqlConn),
		//PmsCommentReplayModel:                    pmsmodel.NewPmsCommentReplayModel(sqlConn),
		//PmsFeightTemplateModel:                   pmsmodel.NewPmsFeightTemplateModel(sqlConn),
		//PmsMemberPriceModel:                      pmsmodel.NewPmsMemberPriceModel(sqlConn),
		//PmsProductAttributeCategoryModel:         pmsmodel.NewPmsProductAttributeCategoryModel(sqlConn),
		//PmsProductAttributeModel:                 pmsmodel.NewPmsProductAttributeModel(sqlConn),
		//PmsProductAttributeValueModel:            pmsmodel.NewPmsProductAttributeValueModel(sqlConn),
		//PmsProductCategoryAttributeRelationModel: pmsmodel.NewPmsProductCategoryAttributeRelationModel(sqlConn),
		//PmsProductCategoryModel:                  pmsmodel.NewPmsProductCategoryModel(sqlConn),
		//PmsProductFullReductionModel:             pmsmodel.NewPmsProductFullReductionModel(sqlConn),
		//PmsProductLadderModel:                    pmsmodel.NewPmsProductLadderModel(sqlConn),
		//PmsProductModel:                          pmsmodel.NewPmsProductModel(sqlConn),
		//PmsProductOperateLogModel:                pmsmodel.NewPmsProductOperateLogModel(sqlConn),
		//PmsProductVertifyRecordModel:             pmsmodel.NewPmsProductVertifyRecordModel(sqlConn),
		//PmsSkuStockModel:                         pmsmodel.NewPmsSkuStockModel(sqlConn),
		//PmsCollectModel:                          pmsmodel.NewPmsProductCollectModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
