package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/internal/config"
)

type ServiceContext struct {
	c config.Config

	SmsCouponHistoryModel                 *smsmodel.SmsCouponHistoryModel
	SmsCouponModel                        *smsmodel.SmsCouponModel
	SmsCouponProductCategoryRelationModel *smsmodel.SmsCouponProductCategoryRelationModel
	SmsCouponProductRelationModel         *smsmodel.SmsCouponProductRelationModel
	SmsFlashPromotionLogModel             *smsmodel.SmsFlashPromotionLogModel
	SmsFlashPromotionModel                *smsmodel.SmsFlashPromotionModel
	SmsFlashPromotionProductRelationModel *smsmodel.SmsFlashPromotionProductRelationModel
	SmsFlashPromotionSessionModel         *smsmodel.SmsFlashPromotionSessionModel
	SmsHomeAdvertiseModel                 *smsmodel.SmsHomeAdvertiseModel
	SmsHomeBrandModel                     *smsmodel.SmsHomeBrandModel
	SmsHomeNewProductModel                *smsmodel.SmsHomeNewProductModel
	SmsHomeRecommendProductModel          *smsmodel.SmsHomeRecommendProductModel
	SmsHomeRecommendSubjectModel          *smsmodel.SmsHomeRecommendSubjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		SmsCouponHistoryModel:                 smsmodel.NewSmsCouponHistoryModel(sqlConn),
		SmsCouponModel:                        smsmodel.NewSmsCouponModel(sqlConn),
		SmsCouponProductCategoryRelationModel: smsmodel.NewSmsCouponProductCategoryRelationModel(sqlConn),
		SmsCouponProductRelationModel:         smsmodel.NewSmsCouponProductRelationModel(sqlConn),
		SmsFlashPromotionLogModel:             smsmodel.NewSmsFlashPromotionLogModel(sqlConn),
		SmsFlashPromotionModel:                smsmodel.NewSmsFlashPromotionModel(sqlConn),
		SmsFlashPromotionProductRelationModel: smsmodel.NewSmsFlashPromotionProductRelationModel(sqlConn),
		SmsFlashPromotionSessionModel:         smsmodel.NewSmsFlashPromotionSessionModel(sqlConn),
		SmsHomeAdvertiseModel:                 smsmodel.NewSmsHomeAdvertiseModel(sqlConn),
		SmsHomeBrandModel:                     smsmodel.NewSmsHomeBrandModel(sqlConn),
		SmsHomeNewProductModel:                smsmodel.NewSmsHomeNewProductModel(sqlConn),
		SmsHomeRecommendProductModel:          smsmodel.NewSmsHomeRecommendProductModel(sqlConn),
		SmsHomeRecommendSubjectModel:          smsmodel.NewSmsHomeRecommendSubjectModel(sqlConn),
	}
}
