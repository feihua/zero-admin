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
	pmsAlbumFieldNames          = builderx.FieldNames(&PmsAlbum{})
	pmsAlbumRows                = strings.Join(pmsAlbumFieldNames, ",")
	pmsAlbumRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsAlbumFieldNames, "id", "create_time", "update_time"), ",")
	pmsAlbumRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsAlbumFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsAlbumModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsAlbum struct {
		Id          int64  `db:"id"`
		Name        string `db:"name"`
		CoverPic    string `db:"cover_pic"`
		PicCount    int64  `db:"pic_count"`
		Sort        int64  `db:"sort"`
		Description string `db:"description"`
	}
)

func NewPmsAlbumModel(conn sqlx.SqlConn) *PmsAlbumModel {
	return &PmsAlbumModel{
		conn:  conn,
		table: "pms_album",
	}
}

func (m *PmsAlbumModel) Insert(data PmsAlbum) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, pmsAlbumRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.CoverPic, data.PicCount, data.Sort, data.Description)
	return ret, err
}

func (m *PmsAlbumModel) FindOne(id int64) (*PmsAlbum, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsAlbumRows, m.table)
	var resp PmsAlbum
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

func (m *PmsAlbumModel) FindAll(Current int64, PageSize int64) (*[]PmsAlbum, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsAlbumRows, m.table)
	var resp []PmsAlbum
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

func (m *PmsAlbumModel) Count() (int64, error) {
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

func (m *PmsAlbumModel) Update(data PmsAlbum) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsAlbumRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.CoverPic, data.PicCount, data.Sort, data.Description, data.Id)
	return err
}

func (m *PmsAlbumModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
