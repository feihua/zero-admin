package pmsmodel

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
	pmsAlbumPicFieldNames          = builderx.FieldNames(&PmsAlbumPic{})
	pmsAlbumPicRows                = strings.Join(pmsAlbumPicFieldNames, ",")
	pmsAlbumPicRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsAlbumPicFieldNames, "id", "create_time", "update_time"), ",")
	pmsAlbumPicRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsAlbumPicFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsAlbumPicModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsAlbumPic struct {
		Id      int64  `db:"id"`
		AlbumId int64  `db:"album_id"`
		Pic     string `db:"pic"`
	}
)

func NewPmsAlbumPicModel(conn sqlx.SqlConn) *PmsAlbumPicModel {
	return &PmsAlbumPicModel{
		conn:  conn,
		table: "pms_album_pic",
	}
}

func (m *PmsAlbumPicModel) Insert(data PmsAlbumPic) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, pmsAlbumPicRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.AlbumId, data.Pic)
	return ret, err
}

func (m *PmsAlbumPicModel) FindOne(id int64) (*PmsAlbumPic, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsAlbumPicRows, m.table)
	var resp PmsAlbumPic
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

func (m *PmsAlbumPicModel) FindAll(Current int64, PageSize int64) (*[]PmsAlbumPic, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsAlbumPicRows, m.table)
	var resp []PmsAlbumPic
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

func (m *PmsAlbumPicModel) Count() (int64, error) {
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

func (m *PmsAlbumPicModel) Update(data PmsAlbumPic) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsAlbumPicRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.AlbumId, data.Pic, data.Id)
	return err
}

func (m *PmsAlbumPicModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
