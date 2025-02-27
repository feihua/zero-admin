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

func newCmsSubjectComment(db *gorm.DB, opts ...gen.DOOption) cmsSubjectComment {
	_cmsSubjectComment := cmsSubjectComment{}

	_cmsSubjectComment.cmsSubjectCommentDo.UseDB(db, opts...)
	_cmsSubjectComment.cmsSubjectCommentDo.UseModel(&model.CmsSubjectComment{})

	tableName := _cmsSubjectComment.cmsSubjectCommentDo.TableName()
	_cmsSubjectComment.ALL = field.NewAsterisk(tableName)
	_cmsSubjectComment.ID = field.NewInt64(tableName, "id")
	_cmsSubjectComment.SubjectID = field.NewInt64(tableName, "subject_id")
	_cmsSubjectComment.MemberNickName = field.NewString(tableName, "member_nick_name")
	_cmsSubjectComment.MemberIcon = field.NewString(tableName, "member_icon")
	_cmsSubjectComment.Content = field.NewString(tableName, "content")
	_cmsSubjectComment.CreateTime = field.NewTime(tableName, "create_time")
	_cmsSubjectComment.ShowStatus = field.NewInt32(tableName, "show_status")

	_cmsSubjectComment.fillFieldMap()

	return _cmsSubjectComment
}

// cmsSubjectComment 专题评论表
type cmsSubjectComment struct {
	cmsSubjectCommentDo cmsSubjectCommentDo

	ALL            field.Asterisk
	ID             field.Int64  // 编号
	SubjectID      field.Int64  // 关联专题id
	MemberNickName field.String // 关联会员昵称
	MemberIcon     field.String // 会员头像
	Content        field.String // 评论内容
	CreateTime     field.Time   // 创建时间
	ShowStatus     field.Int32  // 是否显示，0->不显示；1->显示

	fieldMap map[string]field.Expr
}

func (c cmsSubjectComment) Table(newTableName string) *cmsSubjectComment {
	c.cmsSubjectCommentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cmsSubjectComment) As(alias string) *cmsSubjectComment {
	c.cmsSubjectCommentDo.DO = *(c.cmsSubjectCommentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cmsSubjectComment) updateTableName(table string) *cmsSubjectComment {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.SubjectID = field.NewInt64(table, "subject_id")
	c.MemberNickName = field.NewString(table, "member_nick_name")
	c.MemberIcon = field.NewString(table, "member_icon")
	c.Content = field.NewString(table, "content")
	c.CreateTime = field.NewTime(table, "create_time")
	c.ShowStatus = field.NewInt32(table, "show_status")

	c.fillFieldMap()

	return c
}

func (c *cmsSubjectComment) WithContext(ctx context.Context) ICmsSubjectCommentDo {
	return c.cmsSubjectCommentDo.WithContext(ctx)
}

func (c cmsSubjectComment) TableName() string { return c.cmsSubjectCommentDo.TableName() }

func (c cmsSubjectComment) Alias() string { return c.cmsSubjectCommentDo.Alias() }

func (c cmsSubjectComment) Columns(cols ...field.Expr) gen.Columns {
	return c.cmsSubjectCommentDo.Columns(cols...)
}

func (c *cmsSubjectComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cmsSubjectComment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["subject_id"] = c.SubjectID
	c.fieldMap["member_nick_name"] = c.MemberNickName
	c.fieldMap["member_icon"] = c.MemberIcon
	c.fieldMap["content"] = c.Content
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["show_status"] = c.ShowStatus
}

func (c cmsSubjectComment) clone(db *gorm.DB) cmsSubjectComment {
	c.cmsSubjectCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cmsSubjectComment) replaceDB(db *gorm.DB) cmsSubjectComment {
	c.cmsSubjectCommentDo.ReplaceDB(db)
	return c
}

type cmsSubjectCommentDo struct{ gen.DO }

type ICmsSubjectCommentDo interface {
	gen.SubQuery
	Debug() ICmsSubjectCommentDo
	WithContext(ctx context.Context) ICmsSubjectCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICmsSubjectCommentDo
	WriteDB() ICmsSubjectCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICmsSubjectCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICmsSubjectCommentDo
	Not(conds ...gen.Condition) ICmsSubjectCommentDo
	Or(conds ...gen.Condition) ICmsSubjectCommentDo
	Select(conds ...field.Expr) ICmsSubjectCommentDo
	Where(conds ...gen.Condition) ICmsSubjectCommentDo
	Order(conds ...field.Expr) ICmsSubjectCommentDo
	Distinct(cols ...field.Expr) ICmsSubjectCommentDo
	Omit(cols ...field.Expr) ICmsSubjectCommentDo
	Join(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo
	Group(cols ...field.Expr) ICmsSubjectCommentDo
	Having(conds ...gen.Condition) ICmsSubjectCommentDo
	Limit(limit int) ICmsSubjectCommentDo
	Offset(offset int) ICmsSubjectCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsSubjectCommentDo
	Unscoped() ICmsSubjectCommentDo
	Create(values ...*model.CmsSubjectComment) error
	CreateInBatches(values []*model.CmsSubjectComment, batchSize int) error
	Save(values ...*model.CmsSubjectComment) error
	First() (*model.CmsSubjectComment, error)
	Take() (*model.CmsSubjectComment, error)
	Last() (*model.CmsSubjectComment, error)
	Find() ([]*model.CmsSubjectComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsSubjectComment, err error)
	FindInBatches(result *[]*model.CmsSubjectComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CmsSubjectComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICmsSubjectCommentDo
	Assign(attrs ...field.AssignExpr) ICmsSubjectCommentDo
	Joins(fields ...field.RelationField) ICmsSubjectCommentDo
	Preload(fields ...field.RelationField) ICmsSubjectCommentDo
	FirstOrInit() (*model.CmsSubjectComment, error)
	FirstOrCreate() (*model.CmsSubjectComment, error)
	FindByPage(offset int, limit int) (result []*model.CmsSubjectComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICmsSubjectCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cmsSubjectCommentDo) Debug() ICmsSubjectCommentDo {
	return c.withDO(c.DO.Debug())
}

func (c cmsSubjectCommentDo) WithContext(ctx context.Context) ICmsSubjectCommentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cmsSubjectCommentDo) ReadDB() ICmsSubjectCommentDo {
	return c.Clauses(dbresolver.Read)
}

func (c cmsSubjectCommentDo) WriteDB() ICmsSubjectCommentDo {
	return c.Clauses(dbresolver.Write)
}

func (c cmsSubjectCommentDo) Session(config *gorm.Session) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Session(config))
}

func (c cmsSubjectCommentDo) Clauses(conds ...clause.Expression) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cmsSubjectCommentDo) Returning(value interface{}, columns ...string) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cmsSubjectCommentDo) Not(conds ...gen.Condition) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cmsSubjectCommentDo) Or(conds ...gen.Condition) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cmsSubjectCommentDo) Select(conds ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cmsSubjectCommentDo) Where(conds ...gen.Condition) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cmsSubjectCommentDo) Order(conds ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cmsSubjectCommentDo) Distinct(cols ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cmsSubjectCommentDo) Omit(cols ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cmsSubjectCommentDo) Join(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cmsSubjectCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cmsSubjectCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cmsSubjectCommentDo) Group(cols ...field.Expr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cmsSubjectCommentDo) Having(conds ...gen.Condition) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cmsSubjectCommentDo) Limit(limit int) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cmsSubjectCommentDo) Offset(offset int) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cmsSubjectCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cmsSubjectCommentDo) Unscoped() ICmsSubjectCommentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cmsSubjectCommentDo) Create(values ...*model.CmsSubjectComment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cmsSubjectCommentDo) CreateInBatches(values []*model.CmsSubjectComment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cmsSubjectCommentDo) Save(values ...*model.CmsSubjectComment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cmsSubjectCommentDo) First() (*model.CmsSubjectComment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectComment), nil
	}
}

func (c cmsSubjectCommentDo) Take() (*model.CmsSubjectComment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectComment), nil
	}
}

func (c cmsSubjectCommentDo) Last() (*model.CmsSubjectComment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectComment), nil
	}
}

func (c cmsSubjectCommentDo) Find() ([]*model.CmsSubjectComment, error) {
	result, err := c.DO.Find()
	return result.([]*model.CmsSubjectComment), err
}

func (c cmsSubjectCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CmsSubjectComment, err error) {
	buf := make([]*model.CmsSubjectComment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cmsSubjectCommentDo) FindInBatches(result *[]*model.CmsSubjectComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cmsSubjectCommentDo) Attrs(attrs ...field.AssignExpr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cmsSubjectCommentDo) Assign(attrs ...field.AssignExpr) ICmsSubjectCommentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cmsSubjectCommentDo) Joins(fields ...field.RelationField) ICmsSubjectCommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cmsSubjectCommentDo) Preload(fields ...field.RelationField) ICmsSubjectCommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cmsSubjectCommentDo) FirstOrInit() (*model.CmsSubjectComment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectComment), nil
	}
}

func (c cmsSubjectCommentDo) FirstOrCreate() (*model.CmsSubjectComment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CmsSubjectComment), nil
	}
}

func (c cmsSubjectCommentDo) FindByPage(offset int, limit int) (result []*model.CmsSubjectComment, count int64, err error) {
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

func (c cmsSubjectCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cmsSubjectCommentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cmsSubjectCommentDo) Delete(models ...*model.CmsSubjectComment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cmsSubjectCommentDo) withDO(do gen.Dao) *cmsSubjectCommentDo {
	c.DO = *do.(*gen.DO)
	return c
}
