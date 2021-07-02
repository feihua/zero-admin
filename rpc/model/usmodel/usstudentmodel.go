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
	usStudentFieldNames          = builderx.RawFieldNames(&UsStudent{})
	usStudentRows                = strings.Join(usStudentFieldNames, ",")
	usStudentRowsExpectAutoSet   = strings.Join(stringx.Remove(usStudentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usStudentRowsWithPlaceHolder = strings.Join(stringx.Remove(usStudentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	usStudentRowsForInsert                = strings.Join(stringx.Remove(usStudentFieldNames, "`id`"), ",")
	usStudentRowsForUpdate                = strings.Join(stringx.Remove(usStudentFieldNames, "`id`", "`create_time`"), "=?,") + "=?"

	cacheUsStudentIdPrefix = "cache#usStudent#id#"
)

type (
	UsStudentModel interface {
		Insert(data UsStudent) (sql.Result, error)
		InsertBySession(data UsStudent, session sqlx.Session) (sql.Result, error)
		FindOne(id int64) (*UsStudent, error)
		Update(data UsStudent) error
		Delete(id int64) error
		DeleteCache(id int64) error
	}

	defaultUsStudentModel struct {
		sqlc.CachedConn
		table string
	}

	UsStudent struct {
		Id         int64          `db:"id"`
		Academy    sql.NullString `db:"academy"` // 学院
		Class      sql.NullString `db:"class"`   // 班级
		School     sql.NullString `db:"school"`  // 学校名称
		Grade      sql.NullString `db:"grade"`   // 年级
		UserId     int64          `db:"user_id"` // 用户id
		CreateTime sql.NullTime   `db:"create_time"`
		UpdateTime sql.NullTime   `db:"update_time"`
		DeleteTime sql.NullTime   `db:"delete_time"`
	}
)

func NewUsStudentModel(conn sqlx.SqlConn, c cache.CacheConf) UsStudentModel {
	return &defaultUsStudentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`us_student`",
	}
}

func (m *defaultUsStudentModel) DeleteCache(id int64) error {
	usStudentIdKey := fmt.Sprintf("%s%v", cacheUsStudentIdPrefix, id)
	return m.DelCache(usStudentIdKey)
}

func (m *defaultUsStudentModel) Insert(data UsStudent) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, usStudentRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := m.ExecNoCache(query, data.Academy, data.Class, data.School, data.Grade, data.UserId, data.CreateTime,data.UpdateTime,data.DeleteTime)

	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteCache(id)
	}

	return ret, err
}

//for Transact
func (m *defaultUsStudentModel) InsertBySession(data UsStudent, session sqlx.Session) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, usStudentRowsForInsert)
	data.CreateTime.Time = time.Now()
	data.CreateTime.Valid = true
	data.UpdateTime.Time = time.Now()
	data.UpdateTime.Valid = true
	ret, err := session.Exec(query, data.Academy, data.Class, data.School, data.Grade, data.UserId,data.CreateTime,data.UpdateTime, data.DeleteTime)

	if err == nil {
		id, _ := ret.LastInsertId()
		m.DeleteCache(id)
	}
	return ret, err
}

func (m *defaultUsStudentModel) FindOne(id int64) (*UsStudent, error) {
	usStudentIdKey := fmt.Sprintf("%s%v", cacheUsStudentIdPrefix, id)
	var resp UsStudent
	err := m.QueryRow(&resp, usStudentIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usStudentRows, m.table)
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

func (m *defaultUsStudentModel) Update(data UsStudent) error {
	usStudentIdKey := fmt.Sprintf("%s%v", cacheUsStudentIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usStudentRowsForUpdate)
		data.UpdateTime.Time = time.Now()
		data.UpdateTime.Valid = true
		return conn.Exec(query, data.Academy, data.Class, data.School, data.Grade, data.UserId, data.UpdateTime, data.DeleteTime, data.Id)
	}, usStudentIdKey)
	return err
}

func (m *defaultUsStudentModel) Delete(id int64) error {

	usStudentIdKey := fmt.Sprintf("%s%v", cacheUsStudentIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, usStudentIdKey)
	return err
}

func (m *defaultUsStudentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsStudentIdPrefix, primary)
}

func (m *defaultUsStudentModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usStudentRows, m.table)
	return conn.QueryRow(v, query, primary)
}
