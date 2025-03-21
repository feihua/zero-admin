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

	"github.com/feihua/zero-admin/rpc/cms/gen/model"
)

func newCmsSubjectCategory(db *gorm.DB, opts ...gen.DOOption) cmsSubjectCategory {
	_cmsSubjectCategory := cmsSubjectCategory{}

	_cmsSubjectCategory.cmsSubjectCategoryDo.UseDB(db, opts...)
	_cmsSubjectCategory.cmsSubjectCategoryDo.UseModel(&model.CmsSubjectCategory{})

	tableName := _cmsSubjectCategory.cmsSubjectCategoryDo.TableName()
	_cmsSubjectCategory.ALL = field.NewAsterisk(tableName)
	_cmsSubjectCategory.ID = field.NewInt64(tableName, "id")
	_cmsSubjectCategory.Name = field.NewString(tableName, "name")
	_cmsSubjectCategory.Icon = field.NewString(tableName, "icon")
	_cmsSubjectCategory.SubjectCount = field.NewInt32(tableName, "subject_count")
	_cmsSubjectCategory.ShowStatus = field.NewInt32(tableName, "show_status")
	_cmsSubjectCategory.Sort = field.NewInt32(tableName, "sort")
	_cmsSubjectCategory.CreateBy = field.NewString(tableName, "create_by")
	_cmsSubjectCategory.CreateTime = field.NewTime(tableName, "create_time")
	_cmsSubjectCategory.UpdateBy = field.NewString(tableName, "update_by")
	_cmsSubjectCategory.UpdateTime = field.NewTime(tableName, "update_time")

	_cmsSubjectCategory.fillFieldMap()

	return _cmsSubjectCategory
}

// cmsSubjectCategory 专题分类表
type cmsSubjectCategory struct {
	cmsSubjectCategoryDo cmsSubjectCategoryDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID
	Name         field.String // 专题分类名称
	Icon         field.String // 分类图标
	SubjectCount field.Int32  // 专题数量
	ShowStatus   field.Int32  // 显示状态：0->不显示；1->显示
	Sort         field.Int32  // 排序
	CreateBy     field.String // 创建者
	CreateTime   field.Time   // 创建时间
	UpdateBy     field.String // 更新者
	UpdateTime   field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (c cmsSubjectCategory) Table(newTableName string) *cmsSubjectCategory {
	c.cmsSubjectCategoryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsSubjectCategory) As(alias string) *cmsSubjectCategory {
	c.cmsSubjectCategoryDo.DO = *(c.cmsSubjectCategoryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsSubjectCategory) updateTableName(table string) *cmsSubjectCategory {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Icon = field.NewString(table, "icon")
	c.SubjectCount = field.NewInt32(table, "subject_count")
	c.ShowStatus = field.NewInt32(table, "show_status")
	c.Sort = field.NewInt32(table, "sort")
	c.CreateBy = field.NewString(table, "create_by")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateBy = field.NewString(table, "update_by")
	c.UpdateTime = field.NewTime(table, "update_time")

	c.fillFieldMap()

	return c
}

func (c *cmsSubjectCategory) WithContext(ctx context.Context) ICmsSubjectCategoryDo {
	return c.cmsSubjectCategoryDo.WithContext(ctx)
}

func (c cmsSubjectCategory) TableName() string { return c.cmsSubjectCategoryDo.TableName() }

func (c cmsSubjectCategory) Alias() string { return c.cmsSubjectCategoryDo.Alias() }

func (c cmsSubjectCategory) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsSubjectCategoryDo.Columns(cols...)
}

func (c *cmsSubjectCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsSubjectCategory) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["icon"] = c.Icon
	c.fieldMap["subject_count"] = c.SubjectCount
	c.fieldMap["show_status"] = c.ShowStatus
	c.fieldMap["sort"] = c.Sort
	c.fieldMap["create_by"] = c.CreateBy
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_by"] = c.UpdateBy
	c.fieldMap["update_time"] = c.UpdateTime
}

func (c cmsSubjectCategory) clone(db *gorm.DB) cmsSubjectCategory {
	c.cmsSubjectCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsSubjectCategory) replaceDB(db *gorm.DB) cmsSubjectCategory {
	c.cmsSubjectCategoryDo.ReplaceDB(db)
	return c
}

type cmsSubjectCategoryDo struct{ gen.DO }

type ICmsSubjectCategoryDo interface {
	gen.SubQuery
	Debug() ICmsSubjectCategoryDo
	WithContext(ctx context.Context) ICmsSubjectCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsSubjectCategoryDo
	WriteDB() ICmsSubjectCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsSubjectCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsSubjectCategoryDo
	Not(conds ...gen.Condition) ICmsSubjectCategoryDo
	Or(conds ...gen.Condition) ICmsSubjectCategoryDo
	Select(conds ...field.Expr) ICmsSubjectCategoryDo
	Where(conds ...gen.Condition) ICmsSubjectCategoryDo
	Order(conds ...field.Expr) ICmsSubjectCategoryDo
	Distinct(cols ...field.Expr) ICmsSubjectCategoryDo
	Omit(cols ...field.Expr) ICmsSubjectCategoryDo
	Join(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo
	Group(cols ...field.Expr) ICmsSubjectCategoryDo
	Having(conds ...gen.Condition) ICmsSubjectCategoryDo
	Limit(limit int) ICmsSubjectCategoryDo
	Offset(offset int) ICmsSubjectCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsSubjectCategoryDo
	Unscoped() ICmsSubjectCategoryDo
	Create(values ...*model.CmsSubjectCategory) error
	CreateInBatches(values []*model.CmsSubjectCategory, batchSize int) error
	Save(values ...*model.CmsSubjectCategory) error
	First() (*model.CmsSubjectCategory, error)
	Take() (*model.CmsSubjectCategory, error)
	Last() (*model.CmsSubjectCategory, error)
	Find() ([]*model.CmsSubjectCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsSubjectCategory, err error)
	FindInBatches(result *[]*model.CmsSubjectCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsSubjectCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsSubjectCategoryDo
	Assign(attrs ...field.AssignExpr) ICmsSubjectCategoryDo
	Joins(fields ...field.RelationField) ICmsSubjectCategoryDo
	Preload(fields ...field.RelationField) ICmsSubjectCategoryDo
	FirstOrInit() (*model.CmsSubjectCategory, error)
	FirstOrCreate() (*model.CmsSubjectCategory, error)
	FindByPage(offset int, limit int) (result []*model.CmsSubjectCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsSubjectCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsSubjectCategoryDo) Debug() ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsSubjectCategoryDo) WithContext(ctx context.Context) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsSubjectCategoryDo) ReadDB() ICmsSubjectCategoryDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsSubjectCategoryDo) WriteDB() ICmsSubjectCategoryDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsSubjectCategoryDo) Session(config *gorm.Session) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsSubjectCategoryDo) Clauses(conds ...clause.Expression) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsSubjectCategoryDo) Returning(value interface{}, columns ...string) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsSubjectCategoryDo) Not(conds ...gen.Condition) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsSubjectCategoryDo) Or(conds ...gen.Condition) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsSubjectCategoryDo) Select(conds ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsSubjectCategoryDo) Where(conds ...gen.Condition) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsSubjectCategoryDo) Order(conds ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsSubjectCategoryDo) Distinct(cols ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsSubjectCategoryDo) Omit(cols ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsSubjectCategoryDo) Join(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsSubjectCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsSubjectCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsSubjectCategoryDo) Group(cols ...field.Expr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsSubjectCategoryDo) Having(conds ...gen.Condition) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsSubjectCategoryDo) Limit(limit int) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsSubjectCategoryDo) Offset(offset int) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsSubjectCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsSubjectCategoryDo) Unscoped() ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsSubjectCategoryDo) Create(values ...*model.CmsSubjectCategory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsSubjectCategoryDo) CreateInBatches(values []*model.CmsSubjectCategory, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsSubjectCategoryDo) Save(values ...*model.CmsSubjectCategory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsSubjectCategoryDo) First() (*model.CmsSubjectCategory, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectCategory), nil
	}
}

func (c cmsSubjectCategoryDo) Take() (*model.CmsSubjectCategory, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectCategory), nil
	}
}

func (c cmsSubjectCategoryDo) Last() (*model.CmsSubjectCategory, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectCategory), nil
	}
}

func (c cmsSubjectCategoryDo) Find() ([]*model.CmsSubjectCategory, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsSubjectCategory), err
}

func (c cmsSubjectCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsSubjectCategory, err error) {
	buf := make([]*model.CmsSubjectCategory, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsSubjectCategoryDo) FindInBatches(result *[]*model.CmsSubjectCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsSubjectCategoryDo) Attrs(attrs ...field.AssignExpr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsSubjectCategoryDo) Assign(attrs ...field.AssignExpr) ICmsSubjectCategoryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsSubjectCategoryDo) Joins(fields ...field.RelationField) ICmsSubjectCategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsSubjectCategoryDo) Preload(fields ...field.RelationField) ICmsSubjectCategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsSubjectCategoryDo) FirstOrInit() (*model.CmsSubjectCategory, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectCategory), nil
	}
}

func (c cmsSubjectCategoryDo) FirstOrCreate() (*model.CmsSubjectCategory, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectCategory), nil
	}
}

func (c cmsSubjectCategoryDo) FindByPage(offset int, limit int) (result []*model.CmsSubjectCategory, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cmsSubjectCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsSubjectCategoryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsSubjectCategoryDo) Delete(models ...*model.CmsSubjectCategory) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsSubjectCategoryDo) withDO(do gen.Dao) *cmsSubjectCategoryDo {
	c.DO = *do.(*gen.DO)
	return c
}
