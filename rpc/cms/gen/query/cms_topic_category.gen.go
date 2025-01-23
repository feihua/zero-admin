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

func newCmsTopicCategory(db *gorm.DB, opts ...gen.DOOption) cmsTopicCategory {
	_cmsTopicCategory := cmsTopicCategory{}

	_cmsTopicCategory.cmsTopicCategoryDo.UseDB(db, opts...)
	_cmsTopicCategory.cmsTopicCategoryDo.UseModel(&model.CmsTopicCategory{})

	tableName := _cmsTopicCategory.cmsTopicCategoryDo.TableName()
	_cmsTopicCategory.ALL = field.NewAsterisk(tableName)
	_cmsTopicCategory.ID = field.NewInt64(tableName, "id")
	_cmsTopicCategory.Name = field.NewString(tableName, "name")
	_cmsTopicCategory.Icon = field.NewString(tableName, "icon")
	_cmsTopicCategory.SubjectCount = field.NewInt32(tableName, "subject_count")
	_cmsTopicCategory.ShowStatus = field.NewInt32(tableName, "show_status")
	_cmsTopicCategory.Sort = field.NewInt32(tableName, "sort")
	_cmsTopicCategory.CreateBy = field.NewString(tableName, "create_by")
	_cmsTopicCategory.CreateTime = field.NewTime(tableName, "create_time")
	_cmsTopicCategory.UpdateBy = field.NewString(tableName, "update_by")
	_cmsTopicCategory.UpdateTime = field.NewTime(tableName, "update_time")

	_cmsTopicCategory.fillFieldMap()

	return _cmsTopicCategory
}

// cmsTopicCategory 话题分类表
type cmsTopicCategory struct {
	cmsTopicCategoryDo cmsTopicCategoryDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID
	Name         field.String // 分类名称
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

func (c cmsTopicCategory) Table(newTableName string) *cmsTopicCategory {
	c.cmsTopicCategoryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsTopicCategory) As(alias string) *cmsTopicCategory {
	c.cmsTopicCategoryDo.DO = *(c.cmsTopicCategoryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsTopicCategory) updateTableName(table string) *cmsTopicCategory {
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

func (c *cmsTopicCategory) WithContext(ctx context.Context) ICmsTopicCategoryDo {
	return c.cmsTopicCategoryDo.WithContext(ctx)
}

func (c cmsTopicCategory) TableName() string { return c.cmsTopicCategoryDo.TableName() }

func (c cmsTopicCategory) Alias() string { return c.cmsTopicCategoryDo.Alias() }

func (c cmsTopicCategory) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsTopicCategoryDo.Columns(cols...)
}

func (c *cmsTopicCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsTopicCategory) fillFieldMap() {
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

func (c cmsTopicCategory) clone(db *gorm.DB) cmsTopicCategory {
	c.cmsTopicCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsTopicCategory) replaceDB(db *gorm.DB) cmsTopicCategory {
	c.cmsTopicCategoryDo.ReplaceDB(db)
	return c
}

type cmsTopicCategoryDo struct{ gen.DO }

type ICmsTopicCategoryDo interface {
	gen.SubQuery
	Debug() ICmsTopicCategoryDo
	WithContext(ctx context.Context) ICmsTopicCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsTopicCategoryDo
	WriteDB() ICmsTopicCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsTopicCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsTopicCategoryDo
	Not(conds ...gen.Condition) ICmsTopicCategoryDo
	Or(conds ...gen.Condition) ICmsTopicCategoryDo
	Select(conds ...field.Expr) ICmsTopicCategoryDo
	Where(conds ...gen.Condition) ICmsTopicCategoryDo
	Order(conds ...field.Expr) ICmsTopicCategoryDo
	Distinct(cols ...field.Expr) ICmsTopicCategoryDo
	Omit(cols ...field.Expr) ICmsTopicCategoryDo
	Join(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo
	Group(cols ...field.Expr) ICmsTopicCategoryDo
	Having(conds ...gen.Condition) ICmsTopicCategoryDo
	Limit(limit int) ICmsTopicCategoryDo
	Offset(offset int) ICmsTopicCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsTopicCategoryDo
	Unscoped() ICmsTopicCategoryDo
	Create(values ...*model.CmsTopicCategory) error
	CreateInBatches(values []*model.CmsTopicCategory, batchSize int) error
	Save(values ...*model.CmsTopicCategory) error
	First() (*model.CmsTopicCategory, error)
	Take() (*model.CmsTopicCategory, error)
	Last() (*model.CmsTopicCategory, error)
	Find() ([]*model.CmsTopicCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsTopicCategory, err error)
	FindInBatches(result *[]*model.CmsTopicCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsTopicCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsTopicCategoryDo
	Assign(attrs ...field.AssignExpr) ICmsTopicCategoryDo
	Joins(fields ...field.RelationField) ICmsTopicCategoryDo
	Preload(fields ...field.RelationField) ICmsTopicCategoryDo
	FirstOrInit() (*model.CmsTopicCategory, error)
	FirstOrCreate() (*model.CmsTopicCategory, error)
	FindByPage(offset int, limit int) (result []*model.CmsTopicCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsTopicCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsTopicCategoryDo) Debug() ICmsTopicCategoryDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsTopicCategoryDo) WithContext(ctx context.Context) ICmsTopicCategoryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsTopicCategoryDo) ReadDB() ICmsTopicCategoryDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsTopicCategoryDo) WriteDB() ICmsTopicCategoryDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsTopicCategoryDo) Session(config *gorm.Session) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsTopicCategoryDo) Clauses(conds ...clause.Expression) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsTopicCategoryDo) Returning(value interface{}, columns ...string) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsTopicCategoryDo) Not(conds ...gen.Condition) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsTopicCategoryDo) Or(conds ...gen.Condition) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsTopicCategoryDo) Select(conds ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsTopicCategoryDo) Where(conds ...gen.Condition) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsTopicCategoryDo) Order(conds ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsTopicCategoryDo) Distinct(cols ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsTopicCategoryDo) Omit(cols ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsTopicCategoryDo) Join(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsTopicCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsTopicCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsTopicCategoryDo) Group(cols ...field.Expr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsTopicCategoryDo) Having(conds ...gen.Condition) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsTopicCategoryDo) Limit(limit int) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsTopicCategoryDo) Offset(offset int) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsTopicCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsTopicCategoryDo) Unscoped() ICmsTopicCategoryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsTopicCategoryDo) Create(values ...*model.CmsTopicCategory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsTopicCategoryDo) CreateInBatches(values []*model.CmsTopicCategory, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsTopicCategoryDo) Save(values ...*model.CmsTopicCategory) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsTopicCategoryDo) First() (*model.CmsTopicCategory, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsTopicCategory), nil
	}
}

func (c cmsTopicCategoryDo) Take() (*model.CmsTopicCategory, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsTopicCategory), nil
	}
}

func (c cmsTopicCategoryDo) Last() (*model.CmsTopicCategory, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsTopicCategory), nil
	}
}

func (c cmsTopicCategoryDo) Find() ([]*model.CmsTopicCategory, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsTopicCategory), err
}

func (c cmsTopicCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsTopicCategory, err error) {
	buf := make([]*model.CmsTopicCategory, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsTopicCategoryDo) FindInBatches(result *[]*model.CmsTopicCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsTopicCategoryDo) Attrs(attrs ...field.AssignExpr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsTopicCategoryDo) Assign(attrs ...field.AssignExpr) ICmsTopicCategoryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsTopicCategoryDo) Joins(fields ...field.RelationField) ICmsTopicCategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsTopicCategoryDo) Preload(fields ...field.RelationField) ICmsTopicCategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsTopicCategoryDo) FirstOrInit() (*model.CmsTopicCategory, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsTopicCategory), nil
	}
}

func (c cmsTopicCategoryDo) FirstOrCreate() (*model.CmsTopicCategory, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsTopicCategory), nil
	}
}

func (c cmsTopicCategoryDo) FindByPage(offset int, limit int) (result []*model.CmsTopicCategory, count int64, err error) {
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

func (c cmsTopicCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsTopicCategoryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsTopicCategoryDo) Delete(models ...*model.CmsTopicCategory) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsTopicCategoryDo) withDO(do gen.Dao) *cmsTopicCategoryDo {
	c.DO = *do.(*gen.DO)
	return c
}
