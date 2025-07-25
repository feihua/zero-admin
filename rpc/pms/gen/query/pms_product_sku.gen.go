// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/pms/gen/model"
)

func newPmsProductSku(db *gorm.DB, opts ...gen.DOOption) pmsProductSku {
	_pmsProductSku := pmsProductSku{}

	_pmsProductSku.pmsProductSkuDo.UseDB(db, opts...)
	_pmsProductSku.pmsProductSkuDo.UseModel(&model.PmsProductSku{})

	tableName := _pmsProductSku.pmsProductSkuDo.TableName()
	_pmsProductSku.ALL = field.NewAsterisk(tableName)
	_pmsProductSku.ID = field.NewInt64(tableName, "id")
	_pmsProductSku.SpuID = field.NewInt64(tableName, "spu_id")
	_pmsProductSku.Name = field.NewString(tableName, "name")
	_pmsProductSku.SkuCode = field.NewString(tableName, "sku_code")
	_pmsProductSku.MainPic = field.NewString(tableName, "main_pic")
	_pmsProductSku.AlbumPics = field.NewString(tableName, "album_pics")
	_pmsProductSku.Price = field.NewFloat64(tableName, "price")
	_pmsProductSku.PromotionPrice = field.NewFloat64(tableName, "promotion_price")
	_pmsProductSku.PromotionStartTime = field.NewTime(tableName, "promotion_start_time")
	_pmsProductSku.PromotionEndTime = field.NewTime(tableName, "promotion_end_time")
	_pmsProductSku.Stock = field.NewInt32(tableName, "stock")
	_pmsProductSku.LowStock = field.NewInt32(tableName, "low_stock")
	_pmsProductSku.SpecData = field.NewString(tableName, "spec_data")
	_pmsProductSku.Weight = field.NewFloat64(tableName, "weight")
	_pmsProductSku.PublishStatus = field.NewInt32(tableName, "publish_status")
	_pmsProductSku.VerifyStatus = field.NewInt32(tableName, "verify_status")
	_pmsProductSku.Sort = field.NewInt32(tableName, "sort")
	_pmsProductSku.Sales = field.NewInt32(tableName, "sales")
	_pmsProductSku.CreateBy = field.NewInt64(tableName, "create_by")
	_pmsProductSku.CreateTime = field.NewTime(tableName, "create_time")
	_pmsProductSku.UpdateBy = field.NewInt64(tableName, "update_by")
	_pmsProductSku.UpdateTime = field.NewTime(tableName, "update_time")
	_pmsProductSku.IsDeleted = field.NewInt32(tableName, "is_deleted")

	_pmsProductSku.fillFieldMap()

	return _pmsProductSku
}

// pmsProductSku 商品SKU表
type pmsProductSku struct {
	pmsProductSkuDo pmsProductSkuDo

	ALL                field.Asterisk
	ID                 field.Int64   // 商品SkuId
	SpuID              field.Int64   // 商品SpuId
	Name               field.String  // SKU名称
	SkuCode            field.String  // SKU编码
	MainPic            field.String  // 主图
	AlbumPics          field.String  // 图片集
	Price              field.Float64 // 价格
	PromotionPrice     field.Float64 // 单品促销价格
	PromotionStartTime field.Time    // 促销开始时间
	PromotionEndTime   field.Time    // 促销结束时间
	Stock              field.Int32   // 库存
	LowStock           field.Int32   // 预警库存
	SpecData           field.String  // 规格数据
	Weight             field.Float64 // 重量(kg)
	PublishStatus      field.Int32   // 上架状态：0-下架，1-上架
	VerifyStatus       field.Int32   // 审核状态：0-未审核，1-审核通过，2-审核不通过
	Sort               field.Int32   // 排序
	Sales              field.Int32   // 销量
	CreateBy           field.Int64   // 创建人ID
	CreateTime         field.Time    // 创建时间
	UpdateBy           field.Int64   // 更新人ID
	UpdateTime         field.Time    // 更新时间
	IsDeleted          field.Int32   // 是否删除

	fieldMap map[string]field.Expr
}

func (p pmsProductSku) Table(newTableName string) *pmsProductSku {
	p.pmsProductSkuDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductSku) As(alias string) *pmsProductSku {
	p.pmsProductSkuDo.DO = *(p.pmsProductSkuDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductSku) updateTableName(table string) *pmsProductSku {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.SpuID = field.NewInt64(table, "spu_id")
	p.Name = field.NewString(table, "name")
	p.SkuCode = field.NewString(table, "sku_code")
	p.MainPic = field.NewString(table, "main_pic")
	p.AlbumPics = field.NewString(table, "album_pics")
	p.Price = field.NewFloat64(table, "price")
	p.PromotionPrice = field.NewFloat64(table, "promotion_price")
	p.PromotionStartTime = field.NewTime(table, "promotion_start_time")
	p.PromotionEndTime = field.NewTime(table, "promotion_end_time")
	p.Stock = field.NewInt32(table, "stock")
	p.LowStock = field.NewInt32(table, "low_stock")
	p.SpecData = field.NewString(table, "spec_data")
	p.Weight = field.NewFloat64(table, "weight")
	p.PublishStatus = field.NewInt32(table, "publish_status")
	p.VerifyStatus = field.NewInt32(table, "verify_status")
	p.Sort = field.NewInt32(table, "sort")
	p.Sales = field.NewInt32(table, "sales")
	p.CreateBy = field.NewInt64(table, "create_by")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateBy = field.NewInt64(table, "update_by")
	p.UpdateTime = field.NewTime(table, "update_time")
	p.IsDeleted = field.NewInt32(table, "is_deleted")

	p.fillFieldMap()

	return p
}

func (p *pmsProductSku) WithContext(ctx context.Context) IPmsProductSkuDo {
	return p.pmsProductSkuDo.WithContext(ctx)
}

func (p pmsProductSku) TableName() string { return p.pmsProductSkuDo.TableName() }

func (p pmsProductSku) Alias() string { return p.pmsProductSkuDo.Alias() }

func (p pmsProductSku) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductSkuDo.Columns(cols...)
}

func (p *pmsProductSku) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductSku) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 23)
	p.fieldMap["id"] = p.ID
	p.fieldMap["spu_id"] = p.SpuID
	p.fieldMap["name"] = p.Name
	p.fieldMap["sku_code"] = p.SkuCode
	p.fieldMap["main_pic"] = p.MainPic
	p.fieldMap["album_pics"] = p.AlbumPics
	p.fieldMap["price"] = p.Price
	p.fieldMap["promotion_price"] = p.PromotionPrice
	p.fieldMap["promotion_start_time"] = p.PromotionStartTime
	p.fieldMap["promotion_end_time"] = p.PromotionEndTime
	p.fieldMap["stock"] = p.Stock
	p.fieldMap["low_stock"] = p.LowStock
	p.fieldMap["spec_data"] = p.SpecData
	p.fieldMap["weight"] = p.Weight
	p.fieldMap["publish_status"] = p.PublishStatus
	p.fieldMap["verify_status"] = p.VerifyStatus
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["sales"] = p.Sales
	p.fieldMap["create_by"] = p.CreateBy
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_by"] = p.UpdateBy
	p.fieldMap["update_time"] = p.UpdateTime
	p.fieldMap["is_deleted"] = p.IsDeleted
}

func (p pmsProductSku) clone(db *gorm.DB) pmsProductSku {
	p.pmsProductSkuDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductSku) replaceDB(db *gorm.DB) pmsProductSku {
	p.pmsProductSkuDo.ReplaceDB(db)
	return p
}

type pmsProductSkuDo struct{ gen.DO }

type IPmsProductSkuDo interface {
	gen.SubQuery
	Debug() IPmsProductSkuDo
	WithContext(ctx context.Context) IPmsProductSkuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductSkuDo
	WriteDB() IPmsProductSkuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductSkuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductSkuDo
	Not(conds ...gen.Condition) IPmsProductSkuDo
	Or(conds ...gen.Condition) IPmsProductSkuDo
	Select(conds ...field.Expr) IPmsProductSkuDo
	Where(conds ...gen.Condition) IPmsProductSkuDo
	Order(conds ...field.Expr) IPmsProductSkuDo
	Distinct(cols ...field.Expr) IPmsProductSkuDo
	Omit(cols ...field.Expr) IPmsProductSkuDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo
	Group(cols ...field.Expr) IPmsProductSkuDo
	Having(conds ...gen.Condition) IPmsProductSkuDo
	Limit(limit int) IPmsProductSkuDo
	Offset(offset int) IPmsProductSkuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductSkuDo
	Unscoped() IPmsProductSkuDo
	Create(values ...*model.PmsProductSku) error
	CreateInBatches(values []*model.PmsProductSku, batchSize int) error
	Save(values ...*model.PmsProductSku) error
	First() (*model.PmsProductSku, error)
	Take() (*model.PmsProductSku, error)
	Last() (*model.PmsProductSku, error)
	Find() ([]*model.PmsProductSku, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductSku, err error)
	FindInBatches(result *[]*model.PmsProductSku, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductSku) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductSkuDo
	Assign(attrs ...field.AssignExpr) IPmsProductSkuDo
	Joins(fields ...field.RelationField) IPmsProductSkuDo
	Preload(fields ...field.RelationField) IPmsProductSkuDo
	FirstOrInit() (*model.PmsProductSku, error)
	FirstOrCreate() (*model.PmsProductSku, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductSku, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductSkuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductSkuDo) Debug() IPmsProductSkuDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductSkuDo) WithContext(ctx context.Context) IPmsProductSkuDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductSkuDo) ReadDB() IPmsProductSkuDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductSkuDo) WriteDB() IPmsProductSkuDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductSkuDo) Session(config *gorm.Session) IPmsProductSkuDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductSkuDo) Clauses(conds ...clause.Expression) IPmsProductSkuDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductSkuDo) Returning(value interface{}, columns ...string) IPmsProductSkuDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductSkuDo) Not(conds ...gen.Condition) IPmsProductSkuDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductSkuDo) Or(conds ...gen.Condition) IPmsProductSkuDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductSkuDo) Select(conds ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductSkuDo) Where(conds ...gen.Condition) IPmsProductSkuDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductSkuDo) Order(conds ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductSkuDo) Distinct(cols ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductSkuDo) Omit(cols ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductSkuDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductSkuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductSkuDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductSkuDo) Group(cols ...field.Expr) IPmsProductSkuDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductSkuDo) Having(conds ...gen.Condition) IPmsProductSkuDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductSkuDo) Limit(limit int) IPmsProductSkuDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductSkuDo) Offset(offset int) IPmsProductSkuDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductSkuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductSkuDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductSkuDo) Unscoped() IPmsProductSkuDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductSkuDo) Create(values ...*model.PmsProductSku) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductSkuDo) CreateInBatches(values []*model.PmsProductSku, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductSkuDo) Save(values ...*model.PmsProductSku) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductSkuDo) First() (*model.PmsProductSku, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductSku), nil
	}
}

func (p pmsProductSkuDo) Take() (*model.PmsProductSku, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductSku), nil
	}
}

func (p pmsProductSkuDo) Last() (*model.PmsProductSku, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductSku), nil
	}
}

func (p pmsProductSkuDo) Find() ([]*model.PmsProductSku, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductSku), err
}

func (p pmsProductSkuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductSku, err error) {
	buf := make([]*model.PmsProductSku, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductSkuDo) FindInBatches(result *[]*model.PmsProductSku, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductSkuDo) Attrs(attrs ...field.AssignExpr) IPmsProductSkuDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductSkuDo) Assign(attrs ...field.AssignExpr) IPmsProductSkuDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductSkuDo) Joins(fields ...field.RelationField) IPmsProductSkuDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductSkuDo) Preload(fields ...field.RelationField) IPmsProductSkuDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductSkuDo) FirstOrInit() (*model.PmsProductSku, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductSku), nil
	}
}

func (p pmsProductSkuDo) FirstOrCreate() (*model.PmsProductSku, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductSku), nil
	}
}

func (p pmsProductSkuDo) FindByPage(offset int, limit int) (result []*model.PmsProductSku, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsProductSkuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductSkuDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductSkuDo) Delete(models ...*model.PmsProductSku) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductSkuDo) withDO(do gen.Dao) *pmsProductSkuDo {
	p.DO = *do.(*gen.DO)
	return p
}
