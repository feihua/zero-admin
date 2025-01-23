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

func newPmsFeightTemplate(db *gorm.DB, opts ...gen.DOOption) pmsFeightTemplate {
	_pmsFeightTemplate := pmsFeightTemplate{}

	_pmsFeightTemplate.pmsFeightTemplateDo.UseDB(db, opts...)
	_pmsFeightTemplate.pmsFeightTemplateDo.UseModel(&model.PmsFeightTemplate{})

	tableName := _pmsFeightTemplate.pmsFeightTemplateDo.TableName()
	_pmsFeightTemplate.ALL = field.NewAsterisk(tableName)
	_pmsFeightTemplate.ID = field.NewInt64(tableName, "id")
	_pmsFeightTemplate.Name = field.NewString(tableName, "name")
	_pmsFeightTemplate.ChargeType = field.NewInt32(tableName, "charge_type")
	_pmsFeightTemplate.FirstWeight = field.NewInt64(tableName, "first_weight")
	_pmsFeightTemplate.FirstFee = field.NewInt64(tableName, "first_fee")
	_pmsFeightTemplate.ContinueWeight = field.NewInt64(tableName, "continue_weight")
	_pmsFeightTemplate.ContinueFee = field.NewInt64(tableName, "continue_fee")
	_pmsFeightTemplate.Dest = field.NewString(tableName, "dest")

	_pmsFeightTemplate.fillFieldMap()

	return _pmsFeightTemplate
}

// pmsFeightTemplate 运费模版
type pmsFeightTemplate struct {
	pmsFeightTemplateDo pmsFeightTemplateDo

	ALL            field.Asterisk
	ID             field.Int64
	Name           field.String // 运费模版名称
	ChargeType     field.Int32  // 计费类型:0->按重量；1->按件数
	FirstWeight    field.Int64  // 首重kg
	FirstFee       field.Int64  // 首费（元）
	ContinueWeight field.Int64  // 续重kg
	ContinueFee    field.Int64  // 续费（元）
	Dest           field.String // 目的地（省、市）

	fieldMap map[string]field.Expr
}

func (p pmsFeightTemplate) Table(newTableName string) *pmsFeightTemplate {
	p.pmsFeightTemplateDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsFeightTemplate) As(alias string) *pmsFeightTemplate {
	p.pmsFeightTemplateDo.DO = *(p.pmsFeightTemplateDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsFeightTemplate) updateTableName(table string) *pmsFeightTemplate {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.ChargeType = field.NewInt32(table, "charge_type")
	p.FirstWeight = field.NewInt64(table, "first_weight")
	p.FirstFee = field.NewInt64(table, "first_fee")
	p.ContinueWeight = field.NewInt64(table, "continue_weight")
	p.ContinueFee = field.NewInt64(table, "continue_fee")
	p.Dest = field.NewString(table, "dest")

	p.fillFieldMap()

	return p
}

func (p *pmsFeightTemplate) WithContext(ctx context.Context) IPmsFeightTemplateDo {
	return p.pmsFeightTemplateDo.WithContext(ctx)
}

func (p pmsFeightTemplate) TableName() string { return p.pmsFeightTemplateDo.TableName() }

func (p pmsFeightTemplate) Alias() string { return p.pmsFeightTemplateDo.Alias() }

func (p pmsFeightTemplate) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsFeightTemplateDo.Columns(cols...)
}

func (p *pmsFeightTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsFeightTemplate) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["charge_type"] = p.ChargeType
	p.fieldMap["first_weight"] = p.FirstWeight
	p.fieldMap["first_fee"] = p.FirstFee
	p.fieldMap["continue_weight"] = p.ContinueWeight
	p.fieldMap["continue_fee"] = p.ContinueFee
	p.fieldMap["dest"] = p.Dest
}

func (p pmsFeightTemplate) clone(db *gorm.DB) pmsFeightTemplate {
	p.pmsFeightTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsFeightTemplate) replaceDB(db *gorm.DB) pmsFeightTemplate {
	p.pmsFeightTemplateDo.ReplaceDB(db)
	return p
}

type pmsFeightTemplateDo struct{ gen.DO }

type IPmsFeightTemplateDo interface {
	gen.SubQuery
	Debug() IPmsFeightTemplateDo
	WithContext(ctx context.Context) IPmsFeightTemplateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsFeightTemplateDo
	WriteDB() IPmsFeightTemplateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsFeightTemplateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsFeightTemplateDo
	Not(conds ...gen.Condition) IPmsFeightTemplateDo
	Or(conds ...gen.Condition) IPmsFeightTemplateDo
	Select(conds ...field.Expr) IPmsFeightTemplateDo
	Where(conds ...gen.Condition) IPmsFeightTemplateDo
	Order(conds ...field.Expr) IPmsFeightTemplateDo
	Distinct(cols ...field.Expr) IPmsFeightTemplateDo
	Omit(cols ...field.Expr) IPmsFeightTemplateDo
	Join(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo
	Group(cols ...field.Expr) IPmsFeightTemplateDo
	Having(conds ...gen.Condition) IPmsFeightTemplateDo
	Limit(limit int) IPmsFeightTemplateDo
	Offset(offset int) IPmsFeightTemplateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsFeightTemplateDo
	Unscoped() IPmsFeightTemplateDo
	Create(values ...*model.PmsFeightTemplate) error
	CreateInBatches(values []*model.PmsFeightTemplate, batchSize int) error
	Save(values ...*model.PmsFeightTemplate) error
	First() (*model.PmsFeightTemplate, error)
	Take() (*model.PmsFeightTemplate, error)
	Last() (*model.PmsFeightTemplate, error)
	Find() ([]*model.PmsFeightTemplate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsFeightTemplate, err error)
	FindInBatches(result *[]*model.PmsFeightTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsFeightTemplate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsFeightTemplateDo
	Assign(attrs ...field.AssignExpr) IPmsFeightTemplateDo
	Joins(fields ...field.RelationField) IPmsFeightTemplateDo
	Preload(fields ...field.RelationField) IPmsFeightTemplateDo
	FirstOrInit() (*model.PmsFeightTemplate, error)
	FirstOrCreate() (*model.PmsFeightTemplate, error)
	FindByPage(offset int, limit int) (result []*model.PmsFeightTemplate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsFeightTemplateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsFeightTemplateDo) Debug() IPmsFeightTemplateDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsFeightTemplateDo) WithContext(ctx context.Context) IPmsFeightTemplateDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsFeightTemplateDo) ReadDB() IPmsFeightTemplateDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsFeightTemplateDo) WriteDB() IPmsFeightTemplateDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsFeightTemplateDo) Session(config *gorm.Session) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsFeightTemplateDo) Clauses(conds ...clause.Expression) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsFeightTemplateDo) Returning(value interface{}, columns ...string) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsFeightTemplateDo) Not(conds ...gen.Condition) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsFeightTemplateDo) Or(conds ...gen.Condition) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsFeightTemplateDo) Select(conds ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsFeightTemplateDo) Where(conds ...gen.Condition) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsFeightTemplateDo) Order(conds ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsFeightTemplateDo) Distinct(cols ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsFeightTemplateDo) Omit(cols ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsFeightTemplateDo) Join(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsFeightTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsFeightTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsFeightTemplateDo) Group(cols ...field.Expr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsFeightTemplateDo) Having(conds ...gen.Condition) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsFeightTemplateDo) Limit(limit int) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsFeightTemplateDo) Offset(offset int) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsFeightTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsFeightTemplateDo) Unscoped() IPmsFeightTemplateDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsFeightTemplateDo) Create(values ...*model.PmsFeightTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsFeightTemplateDo) CreateInBatches(values []*model.PmsFeightTemplate, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsFeightTemplateDo) Save(values ...*model.PmsFeightTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsFeightTemplateDo) First() (*model.PmsFeightTemplate, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsFeightTemplate), nil
	}
}

func (p pmsFeightTemplateDo) Take() (*model.PmsFeightTemplate, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsFeightTemplate), nil
	}
}

func (p pmsFeightTemplateDo) Last() (*model.PmsFeightTemplate, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsFeightTemplate), nil
	}
}

func (p pmsFeightTemplateDo) Find() ([]*model.PmsFeightTemplate, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsFeightTemplate), err
}

func (p pmsFeightTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsFeightTemplate, err error) {
	buf := make([]*model.PmsFeightTemplate, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsFeightTemplateDo) FindInBatches(result *[]*model.PmsFeightTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsFeightTemplateDo) Attrs(attrs ...field.AssignExpr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsFeightTemplateDo) Assign(attrs ...field.AssignExpr) IPmsFeightTemplateDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsFeightTemplateDo) Joins(fields ...field.RelationField) IPmsFeightTemplateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsFeightTemplateDo) Preload(fields ...field.RelationField) IPmsFeightTemplateDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsFeightTemplateDo) FirstOrInit() (*model.PmsFeightTemplate, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsFeightTemplate), nil
	}
}

func (p pmsFeightTemplateDo) FirstOrCreate() (*model.PmsFeightTemplate, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsFeightTemplate), nil
	}
}

func (p pmsFeightTemplateDo) FindByPage(offset int, limit int) (result []*model.PmsFeightTemplate, count int64, err error) {
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

func (p pmsFeightTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsFeightTemplateDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsFeightTemplateDo) Delete(models ...*model.PmsFeightTemplate) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsFeightTemplateDo) withDO(do gen.Dao) *pmsFeightTemplateDo {
	p.DO = *do.(*gen.DO)
	return p
}
