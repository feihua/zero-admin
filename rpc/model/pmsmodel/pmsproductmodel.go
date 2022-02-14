package pmsmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
)

var (
	pmsProductFieldNames          = builderx.FieldNames(&PmsProduct{})
	pmsProductRows                = strings.Join(pmsProductFieldNames, ",")
	pmsProductRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsProductFieldNames, "id", "create_time", "update_time"), ",")
	pmsProductRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsProductFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	PmsProductModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsProduct struct {
		Id                         int64     `db:"id"`
		BrandId                    int64     `db:"brand_id"`
		ProductCategoryId          int64     `db:"product_category_id"`
		FeightTemplateId           int64     `db:"feight_template_id"`
		ProductAttributeCategoryId int64     `db:"product_attribute_category_id"`
		Name                       string    `db:"name"`
		Pic                        string    `db:"pic"`
		ProductSn                  string    `db:"product_sn"`       // 货号
		DeleteStatus               int64     `db:"delete_status"`    // 删除状态：0->未删除；1->已删除
		PublishStatus              int64     `db:"publish_status"`   // 上架状态：0->下架；1->上架
		NewStatus                  int64     `db:"new_status"`       // 新品状态:0->不是新品；1->新品
		RecommandStatus            int64     `db:"recommand_status"` // 推荐状态；0->不推荐；1->推荐
		VerifyStatus               int64     `db:"verify_status"`    // 审核状态：0->未审核；1->审核通过
		Sort                       int64     `db:"sort"`             // 排序
		Sale                       int64     `db:"sale"`             // 销量
		Price                      float64   `db:"price"`
		PromotionPrice             float64   `db:"promotion_price"` // 促销价格
		GiftGrowth                 int64     `db:"gift_growth"`     // 赠送的成长值
		GiftPoint                  int64     `db:"gift_point"`      // 赠送的积分
		UsePointLimit              int64     `db:"use_point_limit"` // 限制使用的积分数
		SubTitle                   string    `db:"sub_title"`       // 副标题
		Description                string    `db:"description"`     // 商品描述
		OriginalPrice              float64   `db:"original_price"`  // 市场价
		Stock                      int64     `db:"stock"`           // 库存
		LowStock                   int64     `db:"low_stock"`       // 库存预警值
		Unit                       string    `db:"unit"`            // 单位
		Weight                     float64   `db:"weight"`          // 商品重量，默认为克
		PreviewStatus              int64     `db:"preview_status"`  // 是否为预告商品：0->不是；1->是
		ServiceIds                 string    `db:"service_ids"`     // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		Keywords                   string    `db:"keywords"`
		Note                       string    `db:"note"`
		AlbumPics                  string    `db:"album_pics"` // 画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle                string    `db:"detail_title"`
		DetailDesc                 string    `db:"detail_desc"`
		DetailHtml                 string    `db:"detail_html"`           // 产品详情网页内容
		DetailMobileHtml           string    `db:"detail_mobile_html"`    // 移动端网页详情
		PromotionStartTime         time.Time `db:"promotion_start_time"`  // 促销开始时间
		PromotionEndTime           time.Time `db:"promotion_end_time"`    // 促销结束时间
		PromotionPerLimit          int64     `db:"promotion_per_limit"`   // 活动限购数量
		PromotionType              int64     `db:"promotion_type"`        // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName                  string    `db:"brand_name"`            // 品牌名称
		ProductCategoryName        string    `db:"product_category_name"` // 商品分类名称
	}
)

func NewPmsProductModel(conn sqlx.SqlConn) *PmsProductModel {
	return &PmsProductModel{
		conn:  conn,
		table: "pms_product",
	}
}

func (m *PmsProductModel) Insert(data PmsProduct) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsProductRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BrandId, data.ProductCategoryId, data.FeightTemplateId, data.ProductAttributeCategoryId, data.Name, data.Pic, data.ProductSn, data.DeleteStatus, data.PublishStatus, data.NewStatus, data.RecommandStatus, data.VerifyStatus, data.Sort, data.Sale, data.Price, data.PromotionPrice, data.GiftGrowth, data.GiftPoint, data.UsePointLimit, data.SubTitle, data.Description, data.OriginalPrice, data.Stock, data.LowStock, data.Unit, data.Weight, data.PreviewStatus, data.ServiceIds, data.Keywords, data.Note, data.AlbumPics, data.DetailTitle, data.DetailDesc, data.DetailHtml, data.DetailMobileHtml, data.PromotionStartTime, data.PromotionEndTime, data.PromotionPerLimit, data.PromotionType, data.BrandName, data.ProductCategoryName)
	return ret, err
}

func (m *PmsProductModel) FindOne(id int64) (*PmsProduct, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", pmsProductRows, m.table)
	var resp PmsProduct
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

func (m *PmsProductModel) FindAll(Current int64, PageSize int64) (*[]PmsProduct, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductRows, m.table)
	var resp []PmsProduct
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

func (m *PmsProductModel) Count() (int64, error) {
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

func (m *PmsProductModel) Update(data PmsProduct) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, pmsProductRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BrandId, data.ProductCategoryId, data.FeightTemplateId, data.ProductAttributeCategoryId, data.Name, data.Pic, data.ProductSn, data.DeleteStatus, data.PublishStatus, data.NewStatus, data.RecommandStatus, data.VerifyStatus, data.Sort, data.Sale, data.Price, data.PromotionPrice, data.GiftGrowth, data.GiftPoint, data.UsePointLimit, data.SubTitle, data.Description, data.OriginalPrice, data.Stock, data.LowStock, data.Unit, data.Weight, data.PreviewStatus, data.ServiceIds, data.Keywords, data.Note, data.AlbumPics, data.DetailTitle, data.DetailDesc, data.DetailHtml, data.DetailMobileHtml, data.PromotionStartTime, data.PromotionEndTime, data.PromotionPerLimit, data.PromotionType, data.BrandName, data.ProductCategoryName, data.Id)
	return err
}

func (m *PmsProductModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
