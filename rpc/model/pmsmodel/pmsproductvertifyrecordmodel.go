package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsProductVertifyRecordModel = (*customPmsProductVertifyRecordModel)(nil)

type (
	// PmsProductVertifyRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductVertifyRecordModel.
	PmsProductVertifyRecordModel interface {
		pmsProductVertifyRecordModel
	}

	customPmsProductVertifyRecordModel struct {
		*defaultPmsProductVertifyRecordModel
	}
)

// NewPmsProductVertifyRecordModel returns a model for the database table.
func NewPmsProductVertifyRecordModel(conn sqlx.SqlConn) PmsProductVertifyRecordModel {
	return &customPmsProductVertifyRecordModel{
		defaultPmsProductVertifyRecordModel: newPmsProductVertifyRecordModel(conn),
	}
}
