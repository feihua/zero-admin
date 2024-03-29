// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsTopicCategoryFieldNames          = builder.RawFieldNames(&CmsTopicCategory{})
	cmsTopicCategoryRows                = strings.Join(cmsTopicCategoryFieldNames, ",")
	cmsTopicCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsTopicCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsTopicCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsTopicCategoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsTopicCategoryModel interface {
		Insert(ctx context.Context, data *CmsTopicCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsTopicCategory, error)
		Update(ctx context.Context, data *CmsTopicCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsTopicCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsTopicCategory struct {
		Id           int64  `db:"id"`
		Name         string `db:"name"`
		Icon         string `db:"icon"`          // 分类图标
		SubjectCount int64  `db:"subject_count"` // 专题数量
		ShowStatus   int64  `db:"show_status"`
		Sort         int64  `db:"sort"`
	}
)

func newCmsTopicCategoryModel(conn sqlx.SqlConn) *defaultCmsTopicCategoryModel {
	return &defaultCmsTopicCategoryModel{
		conn:  conn,
		table: "`cms_topic_category`",
	}
}

func (m *defaultCmsTopicCategoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsTopicCategoryModel) FindOne(ctx context.Context, id int64) (*CmsTopicCategory, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsTopicCategoryRows, m.table)
	var resp CmsTopicCategory
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCmsTopicCategoryModel) Insert(ctx context.Context, data *CmsTopicCategory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, cmsTopicCategoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Icon, data.SubjectCount, data.ShowStatus, data.Sort)
	return ret, err
}

func (m *defaultCmsTopicCategoryModel) Update(ctx context.Context, data *CmsTopicCategory) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsTopicCategoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Icon, data.SubjectCount, data.ShowStatus, data.Sort, data.Id)
	return err
}

func (m *defaultCmsTopicCategoryModel) tableName() string {
	return m.table
}
