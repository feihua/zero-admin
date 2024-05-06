package svc

import (
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/config"
	"github.com/feihua/zero-admin/rpc/model/cmsmodel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB

	CmsHelpModel                         cmsmodel.CmsHelpModel
	CmsHelpCategoryModel                 cmsmodel.CmsHelpCategoryModel
	CmsMemberReportModel                 cmsmodel.CmsMemberReportModel
	CmsPrefrenceAreaModel                cmsmodel.CmsPrefrenceAreaModel
	CmsPrefrenceAreaProductRelationModel cmsmodel.CmsPrefrenceAreaProductRelationModel
	CmsSubjectModel                      cmsmodel.CmsSubjectModel
	CmsSubjectCategoryModel              cmsmodel.CmsSubjectCategoryModel
	CmsSubjectCommentModel               cmsmodel.CmsSubjectCommentModel
	CmsSubjectProductRelationModel       cmsmodel.CmsSubjectProductRelationModel
	CmsTopicModel                        cmsmodel.CmsTopicModel
	CmsTopicCategoryModel                cmsmodel.CmsTopicCategoryModel
	CmsTopicCommentModel                 cmsmodel.CmsTopicCommentModel
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

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:                               c,
		DB:                                   DB,
		CmsHelpModel:                         cmsmodel.NewCmsHelpModel(sqlConn),
		CmsHelpCategoryModel:                 cmsmodel.NewCmsHelpCategoryModel(sqlConn),
		CmsMemberReportModel:                 cmsmodel.NewCmsMemberReportModel(sqlConn),
		CmsPrefrenceAreaModel:                cmsmodel.NewCmsPrefrenceAreaModel(sqlConn),
		CmsPrefrenceAreaProductRelationModel: cmsmodel.NewCmsPrefrenceAreaProductRelationModel(sqlConn),
		CmsSubjectModel:                      cmsmodel.NewCmsSubjectModel(sqlConn),
		CmsSubjectCategoryModel:              cmsmodel.NewCmsSubjectCategoryModel(sqlConn),
		CmsSubjectCommentModel:               cmsmodel.NewCmsSubjectCommentModel(sqlConn),
		CmsSubjectProductRelationModel:       cmsmodel.NewCmsSubjectProductRelationModel(sqlConn),
		CmsTopicModel:                        cmsmodel.NewCmsTopicModel(sqlConn),
		CmsTopicCategoryModel:                cmsmodel.NewCmsTopicCategoryModel(sqlConn),
		CmsTopicCommentModel:                 cmsmodel.NewCmsTopicCommentModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Debugf(format, args)
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
