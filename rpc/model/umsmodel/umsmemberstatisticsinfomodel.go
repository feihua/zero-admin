package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberStatisticsInfoModel = (*customUmsMemberStatisticsInfoModel)(nil)

type (
	// UmsMemberStatisticsInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberStatisticsInfoModel.
	UmsMemberStatisticsInfoModel interface {
		umsMemberStatisticsInfoModel
	}

	customUmsMemberStatisticsInfoModel struct {
		*defaultUmsMemberStatisticsInfoModel
	}
)

// NewUmsMemberStatisticsInfoModel returns a model for the database table.
func NewUmsMemberStatisticsInfoModel(conn sqlx.SqlConn) UmsMemberStatisticsInfoModel {
	return &customUmsMemberStatisticsInfoModel{
		defaultUmsMemberStatisticsInfoModel: newUmsMemberStatisticsInfoModel(conn),
	}
}
