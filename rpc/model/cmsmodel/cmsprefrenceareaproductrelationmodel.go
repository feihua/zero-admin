package cmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ CmsPrefrenceAreaProductRelationModel = (*customCmsPrefrenceAreaProductRelationModel)(nil)

type (
	// CmsPrefrenceAreaProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaProductRelationModel.
	CmsPrefrenceAreaProductRelationModel interface {
		cmsPrefrenceAreaProductRelationModel
		FindAll(ctx context.Context, productId int64) (*[]CmsPrefrenceAreaProductRelation, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		DeleteByProductId(ctx context.Context, productId int64) error
	}

	customCmsPrefrenceAreaProductRelationModel struct {
		*defaultCmsPrefrenceAreaProductRelationModel
	}
)

// NewCmsPrefrenceAreaProductRelationModel returns a model for the database table.
func NewCmsPrefrenceAreaProductRelationModel(conn sqlx.SqlConn) CmsPrefrenceAreaProductRelationModel {
	return &customCmsPrefrenceAreaProductRelationModel{
		defaultCmsPrefrenceAreaProductRelationModel: newCmsPrefrenceAreaProductRelationModel(conn),
	}
}

func (m *customCmsPrefrenceAreaProductRelationModel) FindAll(ctx context.Context, productId int64) (*[]CmsPrefrenceAreaProductRelation, error) {

	query := fmt.Sprintf("select %s from %s where product_id =?", cmsPrefrenceAreaProductRelationRows, m.table)
	var resp []CmsPrefrenceAreaProductRelation
	err := m.conn.QueryRows(&resp, query, productId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customCmsPrefrenceAreaProductRelationModel) DeleteByIds(ctx context.Context, ids []int64) error {
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	// 执行删除语句
	_, err := m.conn.ExecCtx(ctx, query, args...)
	return err
}
func (m *customCmsPrefrenceAreaProductRelationModel) DeleteByProductId(ctx context.Context, productId int64) error {
	query := fmt.Sprintf("delete from %s where `product_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, productId)
	return err
}
