package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/sms"
)

var _ SmsHomeRecommendSubjectModel = (*customSmsHomeRecommendSubjectModel)(nil)

type (
	// SmsHomeRecommendSubjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeRecommendSubjectModel.
	SmsHomeRecommendSubjectModel interface {
		smsHomeRecommendSubjectModel
		Count(ctx context.Context, in *sms.HomeRecommendSubjectListReq) (int64, error)
		FindAll(ctx context.Context, in *sms.HomeRecommendSubjectListReq) (*[]SmsHomeRecommendSubject, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindOneBySubjectId(ctx context.Context, subjectId int64) (*SmsHomeRecommendSubject, error)
	}

	customSmsHomeRecommendSubjectModel struct {
		*defaultSmsHomeRecommendSubjectModel
	}
)

// NewSmsHomeRecommendSubjectModel returns a model for the database table.
func NewSmsHomeRecommendSubjectModel(conn sqlx.SqlConn) SmsHomeRecommendSubjectModel {
	return &customSmsHomeRecommendSubjectModel{
		defaultSmsHomeRecommendSubjectModel: newSmsHomeRecommendSubjectModel(conn),
	}
}

func (m *customSmsHomeRecommendSubjectModel) FindAll(ctx context.Context, in *sms.HomeRecommendSubjectListReq) (*[]SmsHomeRecommendSubject, error) {

	where := "1=1"
	if len(in.SubjectName) > 0 {
		where = where + fmt.Sprintf(" AND subject_name like '%%%s%%'", in.SubjectName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsHomeRecommendSubjectRows, m.table, where)
	var resp []SmsHomeRecommendSubject
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsHomeRecommendSubjectModel) Count(ctx context.Context, in *sms.HomeRecommendSubjectListReq) (int64, error) {
	where := "1=1"
	if len(in.SubjectName) > 0 {
		where = where + fmt.Sprintf(" AND subject_name like '%%%s%%'", in.SubjectName)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customSmsHomeRecommendSubjectModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}

func (m *customSmsHomeRecommendSubjectModel) FindOneBySubjectId(ctx context.Context, subjectId int64) (*SmsHomeRecommendSubject, error) {

	where := fmt.Sprintf("subject_id = %d", subjectId)
	query := fmt.Sprintf("select %s from %s where %s ", smsHomeRecommendSubjectRows, m.table, where)
	var resp SmsHomeRecommendSubject
	err := m.conn.QueryRow(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
