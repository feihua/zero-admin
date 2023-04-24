package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/cms/internal/config"
	"zero-admin/rpc/model/cmsmodel"
)

type ServiceContext struct {
	Config config.Config

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
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:                               c,
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
