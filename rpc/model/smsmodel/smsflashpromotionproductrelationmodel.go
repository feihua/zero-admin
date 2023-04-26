package smsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SmsFlashPromotionProductRelationModel = (*customSmsFlashPromotionProductRelationModel)(nil)

type (
	// SmsFlashPromotionProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionProductRelationModel.
	SmsFlashPromotionProductRelationModel interface {
		smsFlashPromotionProductRelationModel
	}

	customSmsFlashPromotionProductRelationModel struct {
		*defaultSmsFlashPromotionProductRelationModel
	}
)

// NewSmsFlashPromotionProductRelationModel returns a model for the database table.
func NewSmsFlashPromotionProductRelationModel(conn sqlx.SqlConn) SmsFlashPromotionProductRelationModel {
	return &customSmsFlashPromotionProductRelationModel{
		defaultSmsFlashPromotionProductRelationModel: newSmsFlashPromotionProductRelationModel(conn),
	}
}
