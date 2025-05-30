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

func newPmsProductBrand(db *gorm.DB, opts ...gen.DOOption) pmsProductBrand {
	_pmsProductBrand := pmsProductBrand{}

	_pmsProductBrand.pmsProductBrandDo.UseDB(db, opts...)
	_pmsProductBrand.pmsProductBrandDo.UseModel(&model.PmsProductBrand{})

	tableName := _pmsProductBrand.pmsProductBrandDo.TableName()
	_pmsProductBrand.ALL = field.NewAsterisk(tableName)
	_pmsProductBrand.ID = field.NewInt64(tableName, "id")
	_pmsProductBrand.Name = field.NewString(tableName, "name")
	_pmsProductBrand.Logo = field.NewString(tableName, "logo")
	_pmsProductBrand.BigPic = field.NewString(tableName, "big_pic")
	_pmsProductBrand.Description = field.NewString(tableName, "description")
	_pmsProductBrand.FirstLetter = field.NewString(tableName, "first_letter")
	_pmsProductBrand.Sort = field.NewInt32(tableName, "sort")
	_pmsProductBrand.RecommendStatus = field.NewInt32(tableName, "recommend_status")
	_pmsProductBrand.ProductCount = field.NewInt32(tableName, "product_count")
	_pmsProductBrand.ProductCommentCount = field.NewInt32(tableName, "product_comment_count")
	_pmsProductBrand.IsEnabled = field.NewInt32(tableName, "is_enabled")
	_pmsProductBrand.CreateBy = field.NewInt64(tableName, "create_by")
	_pmsProductBrand.CreateTime = field.NewTime(tableName, "create_time")
	_pmsProductBrand.UpdateBy = field.NewInt64(tableName, "update_by")
	_pmsProductBrand.UpdateTime = field.NewTime(tableName, "update_time")
	_pmsProductBrand.IsDeleted = field.NewInt32(tableName, "is_deleted")

	_pmsProductBrand.fillFieldMap()

	return _pmsProductBrand
}

// pmsProductBrand 商品品牌
type pmsProductBrand struct {
	pmsProductBrandDo pmsProductBrandDo

	ALL                 field.Asterisk
	ID                  field.Int64
	Name                field.String // 品牌名称
	Logo                field.String // 品牌logo
	BigPic              field.String // 专区大图
	Description         field.String // 描述
	FirstLetter         field.String // 首字母
	Sort                field.Int32  // 排序
	RecommendStatus     field.Int32  // 推荐状态
	ProductCount        field.Int32  // 产品数量
	ProductCommentCount field.Int32  // 产品评论数量
	IsEnabled           field.Int32  // 是否启用
	CreateBy            field.Int64  // 创建人ID
	CreateTime          field.Time   // 创建时间
	UpdateBy            field.Int64  // 更新人ID
	UpdateTime          field.Time   // 更新时间
	IsDeleted           field.Int32  // 是否删除

	fieldMap map[string]field.Expr
}

func (p pmsProductBrand) Table(newTableName string) *pmsProductBrand {
	p.pmsProductBrandDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductBrand) As(alias string) *pmsProductBrand {
	p.pmsProductBrandDo.DO = *(p.pmsProductBrandDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductBrand) updateTableName(table string) *pmsProductBrand {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Name = field.NewString(table, "name")
	p.Logo = field.NewString(table, "logo")
	p.BigPic = field.NewString(table, "big_pic")
	p.Description = field.NewString(table, "description")
	p.FirstLetter = field.NewString(table, "first_letter")
	p.Sort = field.NewInt32(table, "sort")
	p.RecommendStatus = field.NewInt32(table, "recommend_status")
	p.ProductCount = field.NewInt32(table, "product_count")
	p.ProductCommentCount = field.NewInt32(table, "product_comment_count")
	p.IsEnabled = field.NewInt32(table, "is_enabled")
	p.CreateBy = field.NewInt64(table, "create_by")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateBy = field.NewInt64(table, "update_by")
	p.UpdateTime = field.NewTime(table, "update_time")
	p.IsDeleted = field.NewInt32(table, "is_deleted")

	p.fillFieldMap()

	return p
}

func (p *pmsProductBrand) WithContext(ctx context.Context) IPmsProductBrandDo {
	return p.pmsProductBrandDo.WithContext(ctx)
}

func (p pmsProductBrand) TableName() string { return p.pmsProductBrandDo.TableName() }

func (p pmsProductBrand) Alias() string { return p.pmsProductBrandDo.Alias() }

func (p pmsProductBrand) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductBrandDo.Columns(cols...)
}

func (p *pmsProductBrand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductBrand) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["logo"] = p.Logo
	p.fieldMap["big_pic"] = p.BigPic
	p.fieldMap["description"] = p.Description
	p.fieldMap["first_letter"] = p.FirstLetter
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["recommend_status"] = p.RecommendStatus
	p.fieldMap["product_count"] = p.ProductCount
	p.fieldMap["product_comment_count"] = p.ProductCommentCount
	p.fieldMap["is_enabled"] = p.IsEnabled
	p.fieldMap["create_by"] = p.CreateBy
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_by"] = p.UpdateBy
	p.fieldMap["update_time"] = p.UpdateTime
	p.fieldMap["is_deleted"] = p.IsDeleted
}

func (p pmsProductBrand) clone(db *gorm.DB) pmsProductBrand {
	p.pmsProductBrandDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductBrand) replaceDB(db *gorm.DB) pmsProductBrand {
	p.pmsProductBrandDo.ReplaceDB(db)
	return p
}

type pmsProductBrandDo struct{ gen.DO }

type IPmsProductBrandDo interface {
	gen.SubQuery
	Debug() IPmsProductBrandDo
	WithContext(ctx context.Context) IPmsProductBrandDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductBrandDo
	WriteDB() IPmsProductBrandDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductBrandDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductBrandDo
	Not(conds ...gen.Condition) IPmsProductBrandDo
	Or(conds ...gen.Condition) IPmsProductBrandDo
	Select(conds ...field.Expr) IPmsProductBrandDo
	Where(conds ...gen.Condition) IPmsProductBrandDo
	Order(conds ...field.Expr) IPmsProductBrandDo
	Distinct(cols ...field.Expr) IPmsProductBrandDo
	Omit(cols ...field.Expr) IPmsProductBrandDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo
	Group(cols ...field.Expr) IPmsProductBrandDo
	Having(conds ...gen.Condition) IPmsProductBrandDo
	Limit(limit int) IPmsProductBrandDo
	Offset(offset int) IPmsProductBrandDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductBrandDo
	Unscoped() IPmsProductBrandDo
	Create(values ...*model.PmsProductBrand) error
	CreateInBatches(values []*model.PmsProductBrand, batchSize int) error
	Save(values ...*model.PmsProductBrand) error
	First() (*model.PmsProductBrand, error)
	Take() (*model.PmsProductBrand, error)
	Last() (*model.PmsProductBrand, error)
	Find() ([]*model.PmsProductBrand, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductBrand, err error)
	FindInBatches(result *[]*model.PmsProductBrand, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductBrand) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductBrandDo
	Assign(attrs ...field.AssignExpr) IPmsProductBrandDo
	Joins(fields ...field.RelationField) IPmsProductBrandDo
	Preload(fields ...field.RelationField) IPmsProductBrandDo
	FirstOrInit() (*model.PmsProductBrand, error)
	FirstOrCreate() (*model.PmsProductBrand, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductBrand, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductBrandDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductBrandDo) Debug() IPmsProductBrandDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductBrandDo) WithContext(ctx context.Context) IPmsProductBrandDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductBrandDo) ReadDB() IPmsProductBrandDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductBrandDo) WriteDB() IPmsProductBrandDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductBrandDo) Session(config *gorm.Session) IPmsProductBrandDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductBrandDo) Clauses(conds ...clause.Expression) IPmsProductBrandDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductBrandDo) Returning(value interface{}, columns ...string) IPmsProductBrandDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductBrandDo) Not(conds ...gen.Condition) IPmsProductBrandDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductBrandDo) Or(conds ...gen.Condition) IPmsProductBrandDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductBrandDo) Select(conds ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductBrandDo) Where(conds ...gen.Condition) IPmsProductBrandDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductBrandDo) Order(conds ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductBrandDo) Distinct(cols ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductBrandDo) Omit(cols ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductBrandDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductBrandDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductBrandDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductBrandDo) Group(cols ...field.Expr) IPmsProductBrandDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductBrandDo) Having(conds ...gen.Condition) IPmsProductBrandDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductBrandDo) Limit(limit int) IPmsProductBrandDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductBrandDo) Offset(offset int) IPmsProductBrandDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductBrandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductBrandDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductBrandDo) Unscoped() IPmsProductBrandDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductBrandDo) Create(values ...*model.PmsProductBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductBrandDo) CreateInBatches(values []*model.PmsProductBrand, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductBrandDo) Save(values ...*model.PmsProductBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductBrandDo) First() (*model.PmsProductBrand, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductBrand), nil
	}
}

func (p pmsProductBrandDo) Take() (*model.PmsProductBrand, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductBrand), nil
	}
}

func (p pmsProductBrandDo) Last() (*model.PmsProductBrand, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductBrand), nil
	}
}

func (p pmsProductBrandDo) Find() ([]*model.PmsProductBrand, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductBrand), err
}

func (p pmsProductBrandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductBrand, err error) {
	buf := make([]*model.PmsProductBrand, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductBrandDo) FindInBatches(result *[]*model.PmsProductBrand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductBrandDo) Attrs(attrs ...field.AssignExpr) IPmsProductBrandDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductBrandDo) Assign(attrs ...field.AssignExpr) IPmsProductBrandDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductBrandDo) Joins(fields ...field.RelationField) IPmsProductBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductBrandDo) Preload(fields ...field.RelationField) IPmsProductBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductBrandDo) FirstOrInit() (*model.PmsProductBrand, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductBrand), nil
	}
}

func (p pmsProductBrandDo) FirstOrCreate() (*model.PmsProductBrand, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductBrand), nil
	}
}

func (p pmsProductBrandDo) FindByPage(offset int, limit int) (result []*model.PmsProductBrand, count int64, err error) {
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

func (p pmsProductBrandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductBrandDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductBrandDo) Delete(models ...*model.PmsProductBrand) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductBrandDo) withDO(do gen.Dao) *pmsProductBrandDo {
	p.DO = *do.(*gen.DO)
	return p
}
