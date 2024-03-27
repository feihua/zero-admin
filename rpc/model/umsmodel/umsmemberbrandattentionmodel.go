package umsmodel

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberBrandAttentionModel = (*customUmsMemberBrandAttentionModel)(nil)

type (
	// UmsMemberBrandAttentionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberBrandAttentionModel.
	UmsMemberBrandAttentionModel interface {
		umsMemberBrandAttentionModel
		FindAll(ctx context.Context, in *umsclient.MemberBrandAttentionListReq) (*[]UmsMemberBrandAttention, int64, error)
		DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error
		Clear(ctx context.Context, MemberId int64) error
	}

	customUmsMemberBrandAttentionModel struct {
		*defaultUmsMemberBrandAttentionModel
	}
)

// NewUmsMemberBrandAttentionModel returns a model for the database table.
func NewUmsMemberBrandAttentionModel(conn sqlx.SqlConn) UmsMemberBrandAttentionModel {
	return &customUmsMemberBrandAttentionModel{
		defaultUmsMemberBrandAttentionModel: newUmsMemberBrandAttentionModel(conn),
	}
}
func (m *customUmsMemberBrandAttentionModel) FindAll(ctx context.Context, in *umsclient.MemberBrandAttentionListReq) (*[]UmsMemberBrandAttention, int64, error) {
	where := "1=1"

	if in.MemberId != 0 {
		where = where + fmt.Sprintf(" AND member_id = '%d'", in.MemberId)
	}

	query := fmt.Sprintf("select %s from %s where %s", umsMemberBrandAttentionRows, m.table, where)

	var resp []UmsMemberBrandAttention
	err := m.conn.QueryRows(&resp, query)
	if err != nil {
		return nil, 0, err
	}
	query = fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err = m.conn.QueryRow(&count, query)

	if err != nil {
		return nil, 0, err
	}

	return &resp, count, nil
}

func (m *customUmsMemberBrandAttentionModel) DeleteByIdAndMemberId(ctx context.Context, id int64, MemberId int64) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = ? and id = ?", m.table)

	_, err := m.conn.ExecCtx(ctx, query, MemberId, id)
	return err
}

func (m *customUmsMemberBrandAttentionModel) Clear(ctx context.Context, MemberId int64) error {

	//清空关注列表
	query := fmt.Sprintf("DELETE FROM %s WHERE member_id = ? ", m.table)

	_, err := m.conn.ExecCtx(ctx, query, MemberId)
	return err
}
