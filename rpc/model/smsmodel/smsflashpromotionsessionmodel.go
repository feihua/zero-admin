package smsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SmsFlashPromotionSessionModel = (*customSmsFlashPromotionSessionModel)(nil)

type (
	// SmsFlashPromotionSessionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsFlashPromotionSessionModel.
	SmsFlashPromotionSessionModel interface {
		smsFlashPromotionSessionModel
	}

	customSmsFlashPromotionSessionModel struct {
		*defaultSmsFlashPromotionSessionModel
	}
)

// NewSmsFlashPromotionSessionModel returns a model for the database table.
func NewSmsFlashPromotionSessionModel(conn sqlx.SqlConn) SmsFlashPromotionSessionModel {
	return &customSmsFlashPromotionSessionModel{
		defaultSmsFlashPromotionSessionModel: newSmsFlashPromotionSessionModel(conn),
	}
}
