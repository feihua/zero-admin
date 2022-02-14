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
	pmsBrandFieldNames          = builderx.FieldNames(&PmsBrand{})
	pmsBrandRows                = strings.Join(pmsBrandFieldNames, ",")
	pmsBrandRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsBrandFieldNames, "id", "create_time", "update_time"), ",")
	pmsBrandRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsBrandFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsBrandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsBrand struct {
		Id                  int64  `db:"id"`
		Name                string `db:"name"`
		FirstLetter         string `db:"first_letter"` // 首字母
		Sort                int64  `db:"sort"`
		FactoryStatus       int64  `db:"factory_status"` // 是否为品牌制造商：0->不是；1->是
		ShowStatus          int64  `db:"show_status"`
		ProductCount        int64  `db:"product_count"`         // 产品数量
		ProductCommentCount int64  `db:"product_comment_count"` // 产品评论数量
		Logo                string `db:"logo"`                  // 品牌logo
		BigPic              string `db:"big_pic"`               // 专区大图
		BrandStory          string `db:"brand_story"`           // 品牌故事
	}
)

func NewPmsBrandModel(conn sqlx.SqlConn) *PmsBrandModel {
	return &PmsBrandModel{
		conn:  conn,
		table: "pms_brand",
	}
}

func (m *PmsBrandModel) Insert(data PmsBrand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsBrandRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.FirstLetter, data.Sort, data.FactoryStatus, data.ShowStatus, data.ProductCount, data.ProductCommentCount, data.Logo, data.BigPic, data.BrandStory)
	return ret, err
}

func (m *PmsBrandModel) FindOne(id int64) (*PmsBrand, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsBrandRows, m.table)
	var resp PmsBrand
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

func (m *PmsBrandModel) FindAll(Current int64, PageSize int64) (*[]PmsBrand, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsBrandRows, m.table)
	var resp []PmsBrand
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

func (m *PmsBrandModel) Count() (int64, error) {
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

func (m *PmsBrandModel) Update(data PmsBrand) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsBrandRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.FirstLetter, data.Sort, data.FactoryStatus, data.ShowStatus, data.ProductCount, data.ProductCommentCount, data.Logo, data.BigPic, data.BrandStory, data.Id)
	return err
}

func (m *PmsBrandModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
