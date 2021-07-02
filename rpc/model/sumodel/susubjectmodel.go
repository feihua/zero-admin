package sumodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	suSubjectFieldNames          = builderx.RawFieldNames(&SuSubject{})
	suSubjectRows                = strings.Join(suSubjectFieldNames, ",")
	suSubjectRowsExpectAutoSet   = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	suSubjectRowsWithPlaceHolder = strings.Join(stringx.Remove(suSubjectFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSuSubjectIdPrefix   = "cache#suSubject#id#"
	cacheSuSubjectCodePrefix = "cache#suSubject#code#"
)

type (
	SuSubjectModel interface {
		Insert(data SuSubject) (sql.Result, error)
		FindOne(id int64) (*SuSubject, error)
		FindOneByCode(code string) (*SuSubject, error)
		Update(data SuSubject) error
		Delete(id int64) error
	}

	defaultSuSubjectModel struct {
		sqlc.CachedConn
		table string
	}

	SuSubject struct {
		Id                 int64          `db:"id"`
		Uuid               string         `db:"uuid"`                 // uuid
		Name               sql.NullString `db:"name"`                 // 课程名称
		Status             sql.NullInt64  `db:"status"`               // 课程状态
		Code               string         `db:"code"`                 // 课程码
		MaxPersion         sql.NullInt64  `db:"max_persion"`          // 最大人数
		MainTeacherId      int64          `db:"main_teacher_id"`      // 主讲老师id
		AssistantTeacherId sql.NullInt64  `db:"assistant_teacher_id"` // 助教老师id
		Introduce          sql.NullString `db:"introduce"`            // 介绍
		Backup             sql.NullString `db:"backup"`               // 备注
		CreateTime         sql.NullTime   `db:"create_time"`
		UpdateTime         sql.NullTime   `db:"update_time"`
		DeleteTime         sql.NullTime   `db:"delete_time"`
	}
)

func NewSuSubjectModel(conn sqlx.SqlConn, c cache.CacheConf) SuSubjectModel {
	return &defaultSuSubjectModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`su_subject`",
	}
}

func (m *defaultSuSubjectModel) Insert(data SuSubject) (sql.Result, error) {
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, data.Code)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, suSubjectRowsExpectAutoSet)
		return conn.Exec(query, data.Uuid, data.Name, data.Status, data.Code, data.MaxPersion, data.MainTeacherId, data.AssistantTeacherId, data.Introduce, data.Backup, data.DeleteTime)
	}, suSubjectCodeKey)
	return ret, err
}

func (m *defaultSuSubjectModel) FindOne(id int64) (*SuSubject, error) {
	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	var resp SuSubject
	err := m.QueryRow(&resp, suSubjectIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSuSubjectModel) FindOneByCode(code string) (*SuSubject, error) {
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, code)
	var resp SuSubject
	err := m.QueryRowIndex(&resp, suSubjectCodeKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", suSubjectRows, m.table)
		if err := conn.QueryRow(&resp, query, code); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSuSubjectModel) Update(data SuSubject) error {
	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, suSubjectRowsWithPlaceHolder)
		return conn.Exec(query, data.Uuid, data.Name, data.Status, data.Code, data.MaxPersion, data.MainTeacherId, data.AssistantTeacherId, data.Introduce, data.Backup, data.DeleteTime, data.Id)
	}, suSubjectIdKey)
	return err
}

func (m *defaultSuSubjectModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	suSubjectIdKey := fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, id)
	suSubjectCodeKey := fmt.Sprintf("%s%v", cacheSuSubjectCodePrefix, data.Code)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, suSubjectIdKey, suSubjectCodeKey)
	return err
}

func (m *defaultSuSubjectModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSuSubjectIdPrefix, primary)
}

func (m *defaultSuSubjectModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", suSubjectRows, m.table)
	return conn.QueryRow(v, query, primary)
}
