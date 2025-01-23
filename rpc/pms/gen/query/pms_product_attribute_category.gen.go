// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/pms/gen/model"
)

func newPmsProductAttributeCategory(db *gorm.DB, opts ...gen.DOOption) pmsProductAttributeCategory {
	_pmsProductAttributeCategory := pmsProductAttributeCategory{}

	_pmsProductAttributeCategory.pmsProductAttributeCategoryDo.UseDB(db, opts...)
	_pmsProductAttributeCategory.pmsProductAttributeCategoryDo.UseModel(&model.PmsProductAttributeCategory{})

	tableName := _pmsProductAttributeCategory.pmsProductAttributeCategoryDo.TableName()
	_pmsProductAttributeCategory.ALL = field.NewAsterisk(tableName)
	_pmsProductAttributeCategory.ID = field.NewInt64(tableName, "id")
	_pmsProductAttributeCategory.Name = field.NewString(tableName, "name")
	_pmsProductAttributeCategory.AttributeCount = field.NewInt32(tableName, "attribute_count")
	_pmsProductAttributeCategory.ParamCount = field.NewInt32(tableName, "param_count")

	_pmsProductAttributeCategory.fillFieldMap()

	return _pmsProductAttributeCategory
}

// pmsProductAttributeCategory 产品属性分类表
type pmsProductAttributeCategory struct {
	pmsProductAttributeCategoryDo pmsProductAttributeCategoryDo

	ALL            field.Asterisk
	ID             field.Int64
	Name           field.String // 商品属性分类名称
	AttributeCount field.Int32  // 属性数量
	ParamCount     field.Int32  // 参数数量

	fieldMap map[string]field.Expr
}

func (p pmsProductAttributeCategory) Table(newTableName string) *pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductAttributeCategory) As(alias string) *pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.DO = *(p.pmsProductAttributeCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductAttributeCategory) updateTableName(table string) *pmsProductAttributeCategory {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.AttributeCount = field.NewInt32(table, "attribute_count")
	p.ParamCount = field.NewInt32(table, "param_count")

	p.fillFieldMap()

	return p
}

func (p *pmsProductAttributeCategory) WithContext(ctx context.Context) IPmsProductAttributeCategoryDo {
	return p.pmsProductAttributeCategoryDo.WithContext(ctx)
}

func (p pmsProductAttributeCategory) TableName() string {
	return p.pmsProductAttributeCategoryDo.TableName()
}

func (p pmsProductAttributeCategory) Alias() string { return p.pmsProductAttributeCategoryDo.Alias() }

func (p pmsProductAttributeCategory) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductAttributeCategoryDo.Columns(cols...)
}

func (p *pmsProductAttributeCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductAttributeCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["attribute_count"] = p.AttributeCount
	p.fieldMap["param_count"] = p.ParamCount
}

func (p pmsProductAttributeCategory) clone(db *gorm.DB) pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductAttributeCategory) replaceDB(db *gorm.DB) pmsProductAttributeCategory {
	p.pmsProductAttributeCategoryDo.ReplaceDB(db)
	return p
}

type pmsProductAttributeCategoryDo struct{ gen.DO }

type IPmsProductAttributeCategoryDo interface {
	gen.SubQuery
	Debug() IPmsProductAttributeCategoryDo
	WithContext(ctx context.Context) IPmsProductAttributeCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductAttributeCategoryDo
	WriteDB() IPmsProductAttributeCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductAttributeCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductAttributeCategoryDo
	Not(conds ...gen.Condition) IPmsProductAttributeCategoryDo
	Or(conds ...gen.Condition) IPmsProductAttributeCategoryDo
	Select(conds ...field.Expr) IPmsProductAttributeCategoryDo
	Where(conds ...gen.Condition) IPmsProductAttributeCategoryDo
	Order(conds ...field.Expr) IPmsProductAttributeCategoryDo
	Distinct(cols ...field.Expr) IPmsProductAttributeCategoryDo
	Omit(cols ...field.Expr) IPmsProductAttributeCategoryDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo
	Group(cols ...field.Expr) IPmsProductAttributeCategoryDo
	Having(conds ...gen.Condition) IPmsProductAttributeCategoryDo
	Limit(limit int) IPmsProductAttributeCategoryDo
	Offset(offset int) IPmsProductAttributeCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductAttributeCategoryDo
	Unscoped() IPmsProductAttributeCategoryDo
	Create(values ...*model.PmsProductAttributeCategory) error
	CreateInBatches(values []*model.PmsProductAttributeCategory, batchSize int) error
	Save(values ...*model.PmsProductAttributeCategory) error
	First() (*model.PmsProductAttributeCategory, error)
	Take() (*model.PmsProductAttributeCategory, error)
	Last() (*model.PmsProductAttributeCategory, error)
	Find() ([]*model.PmsProductAttributeCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductAttributeCategory, err error)
	FindInBatches(result *[]*model.PmsProductAttributeCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductAttributeCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductAttributeCategoryDo
	Assign(attrs ...field.AssignExpr) IPmsProductAttributeCategoryDo
	Joins(fields ...field.RelationField) IPmsProductAttributeCategoryDo
	Preload(fields ...field.RelationField) IPmsProductAttributeCategoryDo
	FirstOrInit() (*model.PmsProductAttributeCategory, error)
	FirstOrCreate() (*model.PmsProductAttributeCategory, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductAttributeCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductAttributeCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductAttributeCategoryDo) Debug() IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductAttributeCategoryDo) WithContext(ctx context.Context) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductAttributeCategoryDo) ReadDB() IPmsProductAttributeCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductAttributeCategoryDo) WriteDB() IPmsProductAttributeCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductAttributeCategoryDo) Session(config *gorm.Session) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductAttributeCategoryDo) Clauses(conds ...clause.Expression) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductAttributeCategoryDo) Returning(value interface{}, columns ...string) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductAttributeCategoryDo) Not(conds ...gen.Condition) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductAttributeCategoryDo) Or(conds ...gen.Condition) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductAttributeCategoryDo) Select(conds ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductAttributeCategoryDo) Where(conds ...gen.Condition) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductAttributeCategoryDo) Order(conds ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductAttributeCategoryDo) Distinct(cols ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductAttributeCategoryDo) Omit(cols ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductAttributeCategoryDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductAttributeCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductAttributeCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductAttributeCategoryDo) Group(cols ...field.Expr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductAttributeCategoryDo) Having(conds ...gen.Condition) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductAttributeCategoryDo) Limit(limit int) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductAttributeCategoryDo) Offset(offset int) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductAttributeCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductAttributeCategoryDo) Unscoped() IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductAttributeCategoryDo) Create(values ...*model.PmsProductAttributeCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductAttributeCategoryDo) CreateInBatches(values []*model.PmsProductAttributeCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductAttributeCategoryDo) Save(values ...*model.PmsProductAttributeCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductAttributeCategoryDo) First() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Take() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Last() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) Find() ([]*model.PmsProductAttributeCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductAttributeCategory), err
}

func (p pmsProductAttributeCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductAttributeCategory, err error) {
	buf := make([]*model.PmsProductAttributeCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductAttributeCategoryDo) FindInBatches(result *[]*model.PmsProductAttributeCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductAttributeCategoryDo) Attrs(attrs ...field.AssignExpr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductAttributeCategoryDo) Assign(attrs ...field.AssignExpr) IPmsProductAttributeCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductAttributeCategoryDo) Joins(fields ...field.RelationField) IPmsProductAttributeCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductAttributeCategoryDo) Preload(fields ...field.RelationField) IPmsProductAttributeCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductAttributeCategoryDo) FirstOrInit() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) FirstOrCreate() (*model.PmsProductAttributeCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductAttributeCategory), nil
	}
}

func (p pmsProductAttributeCategoryDo) FindByPage(offset int, limit int) (result []*model.PmsProductAttributeCategory, count int64, err error) {
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

func (p pmsProductAttributeCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductAttributeCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductAttributeCategoryDo) Delete(models ...*model.PmsProductAttributeCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductAttributeCategoryDo) withDO(do gen.Dao) *pmsProductAttributeCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}
