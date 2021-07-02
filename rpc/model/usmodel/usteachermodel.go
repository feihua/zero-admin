package usmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usTeacherFieldNames          = builderx.RawFieldNames(&UsTeacher{})
	usTeacherRows                = strings.Join(usTeacherFieldNames, ",")
	usTeacherRowsExpectAutoSet   = strings.Join(stringx.Remove(usTeacherFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usTeacherRowsWithPlaceHolder = strings.Join(stringx.Remove(usTeacherFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	usTeacherRowsForInsert = strings.Join(stringx.Remove(usTeacherFieldNames, "`id`"), ",")
	usTeacherRowsForUpdate = strings.Join(stringx.Remove(usTeacherFieldNames, "`id`", "`create_time`"), "=?,") + "=?"
	usTeacherRowsForDeleteSoft = strings.Join([]string{"delete_time"}, "=?,") + "=?"

	cacheUsTeacherIdPrefix = "cache#usTeacher#id#"
)

type (
	UsTeacherModel interface {
		Insert(data UsTeacher) (sql.Result, error)
		InsertBySession(data UsTeacher, session sqlx.Session) (sql.Result, error)
		FindOne(id int64) (*UsTeacher, error)
		Update(data UsTeacher) error
		Delete(id int64) error
		DeleteIdCache(id int64) error
	}

	defaultUsTeacherModel struct {
		sqlc.CachedConn
		table string
	}

	UsTeacher struct {
		Id         int64          `db:"id"`
		Academy    sql.NullString `db:"academy"` // 学院
		School     sql.NullString `db:"school"`  // 学校
		UserId     int64          `db:"user_id"`
		CreateTime sql.NullTime   `db:"create_time"`
		UpdateTime sql.NullTime   `db:"update_time"`
		DeleteTime sql.NullTime   `db:"delete_time"`
	}
)

func NewUsTeacherModel(conn sqlx.SqlConn, c cache.CacheConf) UsTeacherModel {
	return &defaultUsTeacherModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`us_teacher`",
	}
}

func (m *defaultUsTeacherModel) DeleteIdCache(id int64) error {
	usTeacherIdKey := fmt.Sprintf("%s%v", cacheUsTeacherIdPrefix, id)
	return m.DelCache(usTeacherIdKey)
}

func (m *defaultUsTeacherModel) Insert(data UsTeacher) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?)", m.table, usTeacherRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := m.ExecNoCache(query, data.Academy, data.School, data.UserId, data.CreateTime, data.UpdateTime, data.DeleteTime)

	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteIdCache(id)
	}

	return ret, err
}

//for Transact
func (m *defaultUsTeacherModel) InsertBySession(data UsTeacher, session sqlx.Session) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ( ?, ?, ?, ?, ?, ?)", m.table, usTeacherRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := session.Exec(query, data.Academy, data.School, data.UserId, data.CreateTime, data.UpdateTime, data.DeleteTime)
	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteIdCache(id)
	}

	return ret, err
}

func (m *defaultUsTeacherModel) FindOne(id int64) (*UsTeacher, error) {
	usTeacherIdKey := fmt.Sprintf("%s%v", cacheUsTeacherIdPrefix, id)
	var resp UsTeacher
	err := m.QueryRow(&resp, usTeacherIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `delete_time` is null limit 1", usTeacherRows, m.table)
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

func (m *defaultUsTeacherModel) Update(data UsTeacher) error {
	usTeacherIdKey := fmt.Sprintf("%s%v", cacheUsTeacherIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usTeacherRowsForUpdate)
		data.UpdateTime.Time = time.Now()
		data.UpdateTime.Valid = true
		return conn.Exec(query, data.Academy, data.School, data.UserId, data.UpdateTime, data.DeleteTime, data.Id)
	}, usTeacherIdKey)
	return err
}

func (m *defaultUsTeacherModel) Delete(id int64) error {

	usTeacherIdKey := fmt.Sprintf("%s%v", cacheUsTeacherIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usTeacherIdKey)
	return err
}

func (m *defaultUsTeacherModel) DeleteSoft(id int64) error {
	usUsersIdKey := fmt.Sprintf("%s%v", cacheUsUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usTeacherRowsForDeleteSoft)
		currTime := time.Now()
		return conn.Exec(query, currTime, id)
	}, usUsersIdKey)
	return err
}

func (m *defaultUsTeacherModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsTeacherIdPrefix, primary)
}

func (m *defaultUsTeacherModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usTeacherRows, m.table)
	return conn.QueryRow(v, query, primary)
}
