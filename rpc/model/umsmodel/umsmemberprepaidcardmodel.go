package umsmodel

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberPrepaidCardModel = (*customUmsMemberPrepaidCardModel)(nil)

type (
	// UmsMemberPrepaidCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberPrepaidCardModel.
	UmsMemberPrepaidCardModel interface {
		umsMemberPrepaidCardModel
		FindAll(Current int64, PageSize int64) (*[]UmsMemberPrepaidCard, error)
		Count() (int64, error)
	}

	customUmsMemberPrepaidCardModel struct {
		*defaultUmsMemberPrepaidCardModel
	}
)

// NewUmsMemberPrepaidCardModel returns a model for the database table.
func NewUmsMemberPrepaidCardModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberPrepaidCardModel {
	return &customUmsMemberPrepaidCardModel{
		defaultUmsMemberPrepaidCardModel: newUmsMemberPrepaidCardModel(conn, c),
	}
}

func (m *customUmsMemberPrepaidCardModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberPrepaidCard, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberPrepaidCardRows, m.table)
	var resp []UmsMemberPrepaidCard
	err := m.QueryRowsNoCache(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberPrepaidCardModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
