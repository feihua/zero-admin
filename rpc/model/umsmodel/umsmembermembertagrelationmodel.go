package umsmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	umsMemberMemberTagRelationFieldNames          = builderx.FieldNames(&UmsMemberMemberTagRelation{})
	umsMemberMemberTagRelationRows                = strings.Join(umsMemberMemberTagRelationFieldNames, ",")
	umsMemberMemberTagRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberMemberTagRelationFieldNames, "id", "create_time", "update_time"), ",")
	umsMemberMemberTagRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberMemberTagRelationFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	UmsMemberMemberTagRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberMemberTagRelation struct {
		Id       int64 `db:"id"`
		MemberId int64 `db:"member_id"`
		TagId    int64 `db:"tag_id"`
	}
)

func NewUmsMemberMemberTagRelationModel(conn sqlx.SqlConn) *UmsMemberMemberTagRelationModel {
	return &UmsMemberMemberTagRelationModel{
		conn:  conn,
		table: "ums_member_member_tag_relation",
	}
}

func (m *UmsMemberMemberTagRelationModel) Insert(data UmsMemberMemberTagRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, umsMemberMemberTagRelationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.MemberId, data.TagId)
	return ret, err
}

func (m *UmsMemberMemberTagRelationModel) FindOne(id int64) (*UmsMemberMemberTagRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", umsMemberMemberTagRelationRows, m.table)
	var resp UmsMemberMemberTagRelation
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberMemberTagRelationModel) FindAll(Current int64, PageSize int64) (*[]UmsMemberMemberTagRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberMemberTagRelationRows, m.table)
	var resp []UmsMemberMemberTagRelation
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *UmsMemberMemberTagRelationModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

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

func (m *UmsMemberMemberTagRelationModel) Update(data UmsMemberMemberTagRelation) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, umsMemberMemberTagRelationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.MemberId, data.TagId, data.Id)
	return err
}

func (m *UmsMemberMemberTagRelationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
