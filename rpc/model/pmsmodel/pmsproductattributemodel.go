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
	pmsProductAttributeFieldNames          = builderx.FieldNames(&PmsProductAttribute{})
	pmsProductAttributeRows                = strings.Join(pmsProductAttributeFieldNames, ",")
	pmsProductAttributeRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductAttributeFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductAttributeRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductAttributeFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductAttributeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProductAttribute struct {
		Id                         int64  `db:"id"`
		ProductAttributeCategoryId int64  `db:"product_attribute_category_id"`
		Name                       string `db:"name"`
		SelectType                 int64  `db:"select_type"`     // 属性选择类型：0->唯一；1->单选；2->多选
		InputType                  int64  `db:"input_type"`      // 属性录入方式：0->手工录入；1->从列表中选取
		InputList                  string `db:"input_list"`      // 可选值列表，以逗号隔开
		Sort                       int64  `db:"sort"`            // 排序字段：最高的可以单独上传图片
		FilterType                 int64  `db:"filter_type"`     // 分类筛选样式：1->普通；1->颜色
		SearchType                 int64  `db:"search_type"`     // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
		RelatedStatus              int64  `db:"related_status"`  // 相同属性产品是否关联；0->不关联；1->关联
		HandAddStatus              int64  `db:"hand_add_status"` // 是否支持手动新增；0->不支持；1->支持
		Type                       int64  `db:"type"`            // 属性的类型；0->规格；1->参数
	}
)

func NewPmsProductAttributeModel(conn sqlx.SqlConn) *PmsProductAttributeModel {
	return &PmsProductAttributeModel{
		conn:  conn,
		table: "pms_product_attribute",
	}
}

func (m *PmsProductAttributeModel) Insert(data PmsProductAttribute) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductAttributeRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ProductAttributeCategoryId, data.Name, data.SelectType, data.InputType, data.InputList, data.Sort, data.FilterType, data.SearchType, data.RelatedStatus, data.HandAddStatus, data.Type)
	return ret, err
}

func (m *PmsProductAttributeModel) FindOne(id int64) (*PmsProductAttribute, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductAttributeRows, m.table)
	var resp PmsProductAttribute
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

func (m *PmsProductAttributeModel) FindAll(Current int64, PageSize int64) (*[]PmsProductAttribute, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductAttributeRows, m.table)
	var resp []PmsProductAttribute
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

func (m *PmsProductAttributeModel) Count() (int64, error) {
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

func (m *PmsProductAttributeModel) Update(data PmsProductAttribute) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductAttributeRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ProductAttributeCategoryId, data.Name, data.SelectType, data.InputType, data.InputList, data.Sort, data.FilterType, data.SearchType, data.RelatedStatus, data.HandAddStatus, data.Type, data.Id)
	return err
}

func (m *PmsProductAttributeModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
