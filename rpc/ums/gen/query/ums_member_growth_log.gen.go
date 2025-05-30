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

	"github.com/feihua/zero-admin/rpc/ums/gen/model"
)

func newUmsMemberGrowthLog(db *gorm.DB, opts ...gen.DOOption) umsMemberGrowthLog {
	_umsMemberGrowthLog := umsMemberGrowthLog{}

	_umsMemberGrowthLog.umsMemberGrowthLogDo.UseDB(db, opts...)
	_umsMemberGrowthLog.umsMemberGrowthLogDo.UseModel(&model.UmsMemberGrowthLog{})

	tableName := _umsMemberGrowthLog.umsMemberGrowthLogDo.TableName()
	_umsMemberGrowthLog.ALL = field.NewAsterisk(tableName)
	_umsMemberGrowthLog.ID = field.NewInt64(tableName, "id")
	_umsMemberGrowthLog.MemberID = field.NewInt64(tableName, "member_id")
	_umsMemberGrowthLog.ChangeType = field.NewInt32(tableName, "change_type")
	_umsMemberGrowthLog.ChangeGrowth = field.NewInt32(tableName, "change_growth")
	_umsMemberGrowthLog.SourceType = field.NewInt32(tableName, "source_type")
	_umsMemberGrowthLog.Description = field.NewString(tableName, "description")
	_umsMemberGrowthLog.OperateMan = field.NewString(tableName, "operate_man")
	_umsMemberGrowthLog.OperateNote = field.NewString(tableName, "operate_note")
	_umsMemberGrowthLog.CreateTime = field.NewTime(tableName, "create_time")

	_umsMemberGrowthLog.fillFieldMap()

	return _umsMemberGrowthLog
}

// umsMemberGrowthLog 会员成长值记录表
type umsMemberGrowthLog struct {
	umsMemberGrowthLogDo umsMemberGrowthLogDo

	ALL          field.Asterisk
	ID           field.Int64
	MemberID     field.Int64  // 会员ID
	ChangeType   field.Int32  // 变更类型：1-添加成长值，2-减少成长值
	ChangeGrowth field.Int32  // 变更成长值
	SourceType   field.Int32  // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
	Description  field.String // 描述
	OperateMan   field.String // 操作人员
	OperateNote  field.String // 操作备注
	CreateTime   field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (u umsMemberGrowthLog) Table(newTableName string) *umsMemberGrowthLog {
	u.umsMemberGrowthLogDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberGrowthLog) As(alias string) *umsMemberGrowthLog {
	u.umsMemberGrowthLogDo.DO = *(u.umsMemberGrowthLogDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberGrowthLog) updateTableName(table string) *umsMemberGrowthLog {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberID = field.NewInt64(table, "member_id")
	u.ChangeType = field.NewInt32(table, "change_type")
	u.ChangeGrowth = field.NewInt32(table, "change_growth")
	u.SourceType = field.NewInt32(table, "source_type")
	u.Description = field.NewString(table, "description")
	u.OperateMan = field.NewString(table, "operate_man")
	u.OperateNote = field.NewString(table, "operate_note")
	u.CreateTime = field.NewTime(table, "create_time")

	u.fillFieldMap()

	return u
}

func (u *umsMemberGrowthLog) WithContext(ctx context.Context) IUmsMemberGrowthLogDo {
	return u.umsMemberGrowthLogDo.WithContext(ctx)
}

func (u umsMemberGrowthLog) TableName() string { return u.umsMemberGrowthLogDo.TableName() }

func (u umsMemberGrowthLog) Alias() string { return u.umsMemberGrowthLogDo.Alias() }

func (u umsMemberGrowthLog) Columns(cols ...field.Expr) gen.Columns {
	return u.umsMemberGrowthLogDo.Columns(cols...)
}

func (u *umsMemberGrowthLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberGrowthLog) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["change_type"] = u.ChangeType
	u.fieldMap["change_growth"] = u.ChangeGrowth
	u.fieldMap["source_type"] = u.SourceType
	u.fieldMap["description"] = u.Description
	u.fieldMap["operate_man"] = u.OperateMan
	u.fieldMap["operate_note"] = u.OperateNote
	u.fieldMap["create_time"] = u.CreateTime
}

func (u umsMemberGrowthLog) clone(db *gorm.DB) umsMemberGrowthLog {
	u.umsMemberGrowthLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberGrowthLog) replaceDB(db *gorm.DB) umsMemberGrowthLog {
	u.umsMemberGrowthLogDo.ReplaceDB(db)
	return u
}

type umsMemberGrowthLogDo struct{ gen.DO }

type IUmsMemberGrowthLogDo interface {
	gen.SubQuery
	Debug() IUmsMemberGrowthLogDo
	WithContext(ctx context.Context) IUmsMemberGrowthLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberGrowthLogDo
	WriteDB() IUmsMemberGrowthLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberGrowthLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberGrowthLogDo
	Not(conds ...gen.Condition) IUmsMemberGrowthLogDo
	Or(conds ...gen.Condition) IUmsMemberGrowthLogDo
	Select(conds ...field.Expr) IUmsMemberGrowthLogDo
	Where(conds ...gen.Condition) IUmsMemberGrowthLogDo
	Order(conds ...field.Expr) IUmsMemberGrowthLogDo
	Distinct(cols ...field.Expr) IUmsMemberGrowthLogDo
	Omit(cols ...field.Expr) IUmsMemberGrowthLogDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo
	Group(cols ...field.Expr) IUmsMemberGrowthLogDo
	Having(conds ...gen.Condition) IUmsMemberGrowthLogDo
	Limit(limit int) IUmsMemberGrowthLogDo
	Offset(offset int) IUmsMemberGrowthLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberGrowthLogDo
	Unscoped() IUmsMemberGrowthLogDo
	Create(values ...*model.UmsMemberGrowthLog) error
	CreateInBatches(values []*model.UmsMemberGrowthLog, batchSize int) error
	Save(values ...*model.UmsMemberGrowthLog) error
	First() (*model.UmsMemberGrowthLog, error)
	Take() (*model.UmsMemberGrowthLog, error)
	Last() (*model.UmsMemberGrowthLog, error)
	Find() ([]*model.UmsMemberGrowthLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberGrowthLog, err error)
	FindInBatches(result *[]*model.UmsMemberGrowthLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberGrowthLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberGrowthLogDo
	Assign(attrs ...field.AssignExpr) IUmsMemberGrowthLogDo
	Joins(fields ...field.RelationField) IUmsMemberGrowthLogDo
	Preload(fields ...field.RelationField) IUmsMemberGrowthLogDo
	FirstOrInit() (*model.UmsMemberGrowthLog, error)
	FirstOrCreate() (*model.UmsMemberGrowthLog, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberGrowthLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberGrowthLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberGrowthLogDo) Debug() IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberGrowthLogDo) WithContext(ctx context.Context) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberGrowthLogDo) ReadDB() IUmsMemberGrowthLogDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberGrowthLogDo) WriteDB() IUmsMemberGrowthLogDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberGrowthLogDo) Session(config *gorm.Session) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberGrowthLogDo) Clauses(conds ...clause.Expression) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberGrowthLogDo) Returning(value interface{}, columns ...string) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberGrowthLogDo) Not(conds ...gen.Condition) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberGrowthLogDo) Or(conds ...gen.Condition) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberGrowthLogDo) Select(conds ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberGrowthLogDo) Where(conds ...gen.Condition) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberGrowthLogDo) Order(conds ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberGrowthLogDo) Distinct(cols ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberGrowthLogDo) Omit(cols ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberGrowthLogDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberGrowthLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberGrowthLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberGrowthLogDo) Group(cols ...field.Expr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberGrowthLogDo) Having(conds ...gen.Condition) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberGrowthLogDo) Limit(limit int) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberGrowthLogDo) Offset(offset int) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberGrowthLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberGrowthLogDo) Unscoped() IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberGrowthLogDo) Create(values ...*model.UmsMemberGrowthLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberGrowthLogDo) CreateInBatches(values []*model.UmsMemberGrowthLog, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberGrowthLogDo) Save(values ...*model.UmsMemberGrowthLog) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberGrowthLogDo) First() (*model.UmsMemberGrowthLog, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberGrowthLog), nil
	}
}

func (u umsMemberGrowthLogDo) Take() (*model.UmsMemberGrowthLog, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberGrowthLog), nil
	}
}

func (u umsMemberGrowthLogDo) Last() (*model.UmsMemberGrowthLog, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberGrowthLog), nil
	}
}

func (u umsMemberGrowthLogDo) Find() ([]*model.UmsMemberGrowthLog, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberGrowthLog), err
}

func (u umsMemberGrowthLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberGrowthLog, err error) {
	buf := make([]*model.UmsMemberGrowthLog, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberGrowthLogDo) FindInBatches(result *[]*model.UmsMemberGrowthLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberGrowthLogDo) Attrs(attrs ...field.AssignExpr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberGrowthLogDo) Assign(attrs ...field.AssignExpr) IUmsMemberGrowthLogDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberGrowthLogDo) Joins(fields ...field.RelationField) IUmsMemberGrowthLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberGrowthLogDo) Preload(fields ...field.RelationField) IUmsMemberGrowthLogDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberGrowthLogDo) FirstOrInit() (*model.UmsMemberGrowthLog, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberGrowthLog), nil
	}
}

func (u umsMemberGrowthLogDo) FirstOrCreate() (*model.UmsMemberGrowthLog, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberGrowthLog), nil
	}
}

func (u umsMemberGrowthLogDo) FindByPage(offset int, limit int) (result []*model.UmsMemberGrowthLog, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsMemberGrowthLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberGrowthLogDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberGrowthLogDo) Delete(models ...*model.UmsMemberGrowthLog) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberGrowthLogDo) withDO(do gen.Dao) *umsMemberGrowthLogDo {
	u.DO = *do.(*gen.DO)
	return u
}
