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

func newPmsSkuStock(db *gorm.DB, opts ...gen.DOOption) pmsSkuStock {
	_pmsSkuStock := pmsSkuStock{}

	_pmsSkuStock.pmsSkuStockDo.UseDB(db, opts...)
	_pmsSkuStock.pmsSkuStockDo.UseModel(&model.PmsSkuStock{})

	tableName := _pmsSkuStock.pmsSkuStockDo.TableName()
	_pmsSkuStock.ALL = field.NewAsterisk(tableName)
	_pmsSkuStock.ID = field.NewInt64(tableName, "id")
	_pmsSkuStock.ProductID = field.NewInt64(tableName, "product_id")
	_pmsSkuStock.SkuCode = field.NewString(tableName, "sku_code")
	_pmsSkuStock.Price = field.NewInt64(tableName, "price")
	_pmsSkuStock.Stock = field.NewInt32(tableName, "stock")
	_pmsSkuStock.LowStock = field.NewInt32(tableName, "low_stock")
	_pmsSkuStock.Pic = field.NewString(tableName, "pic")
	_pmsSkuStock.Sale = field.NewInt32(tableName, "sale")
	_pmsSkuStock.PromotionPrice = field.NewInt64(tableName, "promotion_price")
	_pmsSkuStock.LockStock = field.NewInt32(tableName, "lock_stock")
	_pmsSkuStock.SpData = field.NewString(tableName, "sp_data")
	_pmsSkuStock.CreateTime = field.NewTime(tableName, "create_time")
	_pmsSkuStock.UpdateTime = field.NewTime(tableName, "update_time")

	_pmsSkuStock.fillFieldMap()

	return _pmsSkuStock
}

// pmsSkuStock sku的库存
type pmsSkuStock struct {
	pmsSkuStockDo pmsSkuStockDo

	ALL            field.Asterisk
	ID             field.Int64
	ProductID      field.Int64  // 商品id
	SkuCode        field.String // sku编码
	Price          field.Int64  // 价格
	Stock          field.Int32  // 库存
	LowStock       field.Int32  // 预警库存
	Pic            field.String // 展示图片
	Sale           field.Int32  // 销量
	PromotionPrice field.Int64  // 单品促销价格
	LockStock      field.Int32  // 锁定库存
	SpData         field.String // 商品销售属性，json格式
	CreateTime     field.Time   // 创建时间
	UpdateTime     field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (p pmsSkuStock) Table(newTableName string) *pmsSkuStock {
	p.pmsSkuStockDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsSkuStock) As(alias string) *pmsSkuStock {
	p.pmsSkuStockDo.DO = *(p.pmsSkuStockDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsSkuStock) updateTableName(table string) *pmsSkuStock {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProductID = field.NewInt64(table, "product_id")
	p.SkuCode = field.NewString(table, "sku_code")
	p.Price = field.NewInt64(table, "price")
	p.Stock = field.NewInt32(table, "stock")
	p.LowStock = field.NewInt32(table, "low_stock")
	p.Pic = field.NewString(table, "pic")
	p.Sale = field.NewInt32(table, "sale")
	p.PromotionPrice = field.NewInt64(table, "promotion_price")
	p.LockStock = field.NewInt32(table, "lock_stock")
	p.SpData = field.NewString(table, "sp_data")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateTime = field.NewTime(table, "update_time")

	p.fillFieldMap()

	return p
}

func (p *pmsSkuStock) WithContext(ctx context.Context) IPmsSkuStockDo {
	return p.pmsSkuStockDo.WithContext(ctx)
}

func (p pmsSkuStock) TableName() string { return p.pmsSkuStockDo.TableName() }

func (p pmsSkuStock) Alias() string { return p.pmsSkuStockDo.Alias() }

func (p pmsSkuStock) Columns(cols ...field.Expr) gen.Columns { return p.pmsSkuStockDo.Columns(cols...) }

func (p *pmsSkuStock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsSkuStock) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["id"] = p.ID
	p.fieldMap["product_id"] = p.ProductID
	p.fieldMap["sku_code"] = p.SkuCode
	p.fieldMap["price"] = p.Price
	p.fieldMap["stock"] = p.Stock
	p.fieldMap["low_stock"] = p.LowStock
	p.fieldMap["pic"] = p.Pic
	p.fieldMap["sale"] = p.Sale
	p.fieldMap["promotion_price"] = p.PromotionPrice
	p.fieldMap["lock_stock"] = p.LockStock
	p.fieldMap["sp_data"] = p.SpData
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_time"] = p.UpdateTime
}

func (p pmsSkuStock) clone(db *gorm.DB) pmsSkuStock {
	p.pmsSkuStockDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsSkuStock) replaceDB(db *gorm.DB) pmsSkuStock {
	p.pmsSkuStockDo.ReplaceDB(db)
	return p
}

type pmsSkuStockDo struct{ gen.DO }

type IPmsSkuStockDo interface {
	gen.SubQuery
	Debug() IPmsSkuStockDo
	WithContext(ctx context.Context) IPmsSkuStockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsSkuStockDo
	WriteDB() IPmsSkuStockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsSkuStockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsSkuStockDo
	Not(conds ...gen.Condition) IPmsSkuStockDo
	Or(conds ...gen.Condition) IPmsSkuStockDo
	Select(conds ...field.Expr) IPmsSkuStockDo
	Where(conds ...gen.Condition) IPmsSkuStockDo
	Order(conds ...field.Expr) IPmsSkuStockDo
	Distinct(cols ...field.Expr) IPmsSkuStockDo
	Omit(cols ...field.Expr) IPmsSkuStockDo
	Join(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo
	Group(cols ...field.Expr) IPmsSkuStockDo
	Having(conds ...gen.Condition) IPmsSkuStockDo
	Limit(limit int) IPmsSkuStockDo
	Offset(offset int) IPmsSkuStockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSkuStockDo
	Unscoped() IPmsSkuStockDo
	Create(values ...*model.PmsSkuStock) error
	CreateInBatches(values []*model.PmsSkuStock, batchSize int) error
	Save(values ...*model.PmsSkuStock) error
	First() (*model.PmsSkuStock, error)
	Take() (*model.PmsSkuStock, error)
	Last() (*model.PmsSkuStock, error)
	Find() ([]*model.PmsSkuStock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuStock, err error)
	FindInBatches(result *[]*model.PmsSkuStock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsSkuStock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsSkuStockDo
	Assign(attrs ...field.AssignExpr) IPmsSkuStockDo
	Joins(fields ...field.RelationField) IPmsSkuStockDo
	Preload(fields ...field.RelationField) IPmsSkuStockDo
	FirstOrInit() (*model.PmsSkuStock, error)
	FirstOrCreate() (*model.PmsSkuStock, error)
	FindByPage(offset int, limit int) (result []*model.PmsSkuStock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsSkuStockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsSkuStockDo) Debug() IPmsSkuStockDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsSkuStockDo) WithContext(ctx context.Context) IPmsSkuStockDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsSkuStockDo) ReadDB() IPmsSkuStockDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsSkuStockDo) WriteDB() IPmsSkuStockDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsSkuStockDo) Session(config *gorm.Session) IPmsSkuStockDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsSkuStockDo) Clauses(conds ...clause.Expression) IPmsSkuStockDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsSkuStockDo) Returning(value interface{}, columns ...string) IPmsSkuStockDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsSkuStockDo) Not(conds ...gen.Condition) IPmsSkuStockDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsSkuStockDo) Or(conds ...gen.Condition) IPmsSkuStockDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsSkuStockDo) Select(conds ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsSkuStockDo) Where(conds ...gen.Condition) IPmsSkuStockDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsSkuStockDo) Order(conds ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsSkuStockDo) Distinct(cols ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsSkuStockDo) Omit(cols ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsSkuStockDo) Join(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsSkuStockDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsSkuStockDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsSkuStockDo) Group(cols ...field.Expr) IPmsSkuStockDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsSkuStockDo) Having(conds ...gen.Condition) IPmsSkuStockDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsSkuStockDo) Limit(limit int) IPmsSkuStockDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsSkuStockDo) Offset(offset int) IPmsSkuStockDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsSkuStockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsSkuStockDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsSkuStockDo) Unscoped() IPmsSkuStockDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsSkuStockDo) Create(values ...*model.PmsSkuStock) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsSkuStockDo) CreateInBatches(values []*model.PmsSkuStock, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsSkuStockDo) Save(values ...*model.PmsSkuStock) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsSkuStockDo) First() (*model.PmsSkuStock, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuStock), nil
	}
}

func (p pmsSkuStockDo) Take() (*model.PmsSkuStock, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuStock), nil
	}
}

func (p pmsSkuStockDo) Last() (*model.PmsSkuStock, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuStock), nil
	}
}

func (p pmsSkuStockDo) Find() ([]*model.PmsSkuStock, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsSkuStock), err
}

func (p pmsSkuStockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsSkuStock, err error) {
	buf := make([]*model.PmsSkuStock, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsSkuStockDo) FindInBatches(result *[]*model.PmsSkuStock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsSkuStockDo) Attrs(attrs ...field.AssignExpr) IPmsSkuStockDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsSkuStockDo) Assign(attrs ...field.AssignExpr) IPmsSkuStockDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsSkuStockDo) Joins(fields ...field.RelationField) IPmsSkuStockDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsSkuStockDo) Preload(fields ...field.RelationField) IPmsSkuStockDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsSkuStockDo) FirstOrInit() (*model.PmsSkuStock, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuStock), nil
	}
}

func (p pmsSkuStockDo) FirstOrCreate() (*model.PmsSkuStock, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsSkuStock), nil
	}
}

func (p pmsSkuStockDo) FindByPage(offset int, limit int) (result []*model.PmsSkuStock, count int64, err error) {
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

func (p pmsSkuStockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsSkuStockDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsSkuStockDo) Delete(models ...*model.PmsSkuStock) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsSkuStockDo) withDO(do gen.Dao) *pmsSkuStockDo {
	p.DO = *do.(*gen.DO)
	return p
}
