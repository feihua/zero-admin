package cmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/cms/cmsclient"
)

var _ CmsSubjectModel = (*customCmsSubjectModel)(nil)

type (
	// CmsSubjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectModel.
	CmsSubjectModel interface {
		cmsSubjectModel
		Count(ctx context.Context, in *cmsclient.SubjectListReq) (int64, error)
		FindAll(ctx context.Context, in *cmsclient.SubjectListReq) (*[]CmsSubject, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindAllByIds(ctx context.Context, ids []int64) (*[]CmsSubject, error)
	}

	customCmsSubjectModel struct {
		*defaultCmsSubjectModel
	}
)

// NewCmsSubjectModel returns a model for the database table.
func NewCmsSubjectModel(conn sqlx.SqlConn) CmsSubjectModel {
	return &customCmsSubjectModel{
		defaultCmsSubjectModel: newCmsSubjectModel(conn),
	}
}

func (m *customCmsSubjectModel) FindAll(ctx context.Context, in *cmsclient.SubjectListReq) (*[]CmsSubject, error) {

	where := "1=1"
	if len(in.Title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.Title)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	if in.ShowStatus != 2 {
		where = where + fmt.Sprintf(" AND show_status = %d", in.RecommendStatus)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", cmsSubjectRows, m.table, where)
	var resp []CmsSubject
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

func (m *customCmsSubjectModel) Count(ctx context.Context, in *cmsclient.SubjectListReq) (int64, error) {
	where := "1=1"
	if len(in.Title) > 0 {
		where = where + fmt.Sprintf(" AND title like '%%%s%%'", in.Title)
	}
	if in.RecommendStatus != 2 {
		where = where + fmt.Sprintf(" AND recommend_status = %d", in.RecommendStatus)
	}
	if in.ShowStatus != 2 {
		where = where + fmt.Sprintf(" AND show_status = %d", in.RecommendStatus)
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

func (m *customCmsSubjectModel) FindAllByIds(ctx context.Context, ids []int64) (*[]CmsSubject, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", cmsSubjectRows, m.table)
	var resp []CmsSubject
	err := m.conn.QueryRows(&resp, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customCmsSubjectModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
