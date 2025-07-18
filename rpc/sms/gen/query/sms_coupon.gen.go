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

	"github.com/feihua/zero-admin/rpc/sms/gen/model"
)

func newSmsCoupon(db *gorm.DB, opts ...gen.DOOption) smsCoupon {
	_smsCoupon := smsCoupon{}

	_smsCoupon.smsCouponDo.UseDB(db, opts...)
	_smsCoupon.smsCouponDo.UseModel(&model.SmsCoupon{})

	tableName := _smsCoupon.smsCouponDo.TableName()
	_smsCoupon.ALL = field.NewAsterisk(tableName)
	_smsCoupon.ID = field.NewInt64(tableName, "id")
	_smsCoupon.TypeID = field.NewInt64(tableName, "type_id")
	_smsCoupon.Name = field.NewString(tableName, "name")
	_smsCoupon.Code = field.NewString(tableName, "code")
	_smsCoupon.Amount = field.NewFloat64(tableName, "amount")
	_smsCoupon.MinAmount = field.NewFloat64(tableName, "min_amount")
	_smsCoupon.StartTime = field.NewTime(tableName, "start_time")
	_smsCoupon.EndTime = field.NewTime(tableName, "end_time")
	_smsCoupon.TotalCount = field.NewInt32(tableName, "total_count")
	_smsCoupon.ReceivedCount = field.NewInt32(tableName, "received_count")
	_smsCoupon.UsedCount = field.NewInt32(tableName, "used_count")
	_smsCoupon.PerLimit = field.NewInt32(tableName, "per_limit")
	_smsCoupon.Status = field.NewInt32(tableName, "status")
	_smsCoupon.IsEnabled = field.NewInt32(tableName, "is_enabled")
	_smsCoupon.Description = field.NewString(tableName, "description")
	_smsCoupon.CreateBy = field.NewInt64(tableName, "create_by")
	_smsCoupon.CreateTime = field.NewTime(tableName, "create_time")
	_smsCoupon.UpdateBy = field.NewInt64(tableName, "update_by")
	_smsCoupon.UpdateTime = field.NewTime(tableName, "update_time")
	_smsCoupon.IsDeleted = field.NewInt32(tableName, "is_deleted")

	_smsCoupon.fillFieldMap()

	return _smsCoupon
}

// smsCoupon 优惠券表
type smsCoupon struct {
	smsCouponDo smsCouponDo

	ALL           field.Asterisk
	ID            field.Int64   // 优惠券ID
	TypeID        field.Int64   // 优惠券类型ID
	Name          field.String  // 优惠券名称
	Code          field.String  // 优惠券码
	Amount        field.Float64 // 优惠金额/折扣率
	MinAmount     field.Float64 // 最低使用金额
	StartTime     field.Time    // 生效时间
	EndTime       field.Time    // 失效时间
	TotalCount    field.Int32   // 发放总量
	ReceivedCount field.Int32   // 已领取数量
	UsedCount     field.Int32   // 已使用数量
	PerLimit      field.Int32   // 每人限领数量
	Status        field.Int32   // 状态：0-未开始，1-进行中，2-已结束，3-已取消
	IsEnabled     field.Int32   // 是否启用
	Description   field.String  // 使用说明
	CreateBy      field.Int64   // 创建人ID
	CreateTime    field.Time    // 创建时间
	UpdateBy      field.Int64   // 更新人ID
	UpdateTime    field.Time    // 更新时间
	IsDeleted     field.Int32   // 是否删除

	fieldMap map[string]field.Expr
}

func (s smsCoupon) Table(newTableName string) *smsCoupon {
	s.smsCouponDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsCoupon) As(alias string) *smsCoupon {
	s.smsCouponDo.DO = *(s.smsCouponDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsCoupon) updateTableName(table string) *smsCoupon {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.TypeID = field.NewInt64(table, "type_id")
	s.Name = field.NewString(table, "name")
	s.Code = field.NewString(table, "code")
	s.Amount = field.NewFloat64(table, "amount")
	s.MinAmount = field.NewFloat64(table, "min_amount")
	s.StartTime = field.NewTime(table, "start_time")
	s.EndTime = field.NewTime(table, "end_time")
	s.TotalCount = field.NewInt32(table, "total_count")
	s.ReceivedCount = field.NewInt32(table, "received_count")
	s.UsedCount = field.NewInt32(table, "used_count")
	s.PerLimit = field.NewInt32(table, "per_limit")
	s.Status = field.NewInt32(table, "status")
	s.IsEnabled = field.NewInt32(table, "is_enabled")
	s.Description = field.NewString(table, "description")
	s.CreateBy = field.NewInt64(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewInt64(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.IsDeleted = field.NewInt32(table, "is_deleted")

	s.fillFieldMap()

	return s
}

func (s *smsCoupon) WithContext(ctx context.Context) ISmsCouponDo {
	return s.smsCouponDo.WithContext(ctx)
}

func (s smsCoupon) TableName() string { return s.smsCouponDo.TableName() }

func (s smsCoupon) Alias() string { return s.smsCouponDo.Alias() }

func (s smsCoupon) Columns(cols ...field.Expr) gen.Columns { return s.smsCouponDo.Columns(cols...) }

func (s *smsCoupon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsCoupon) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["id"] = s.ID
	s.fieldMap["type_id"] = s.TypeID
	s.fieldMap["name"] = s.Name
	s.fieldMap["code"] = s.Code
	s.fieldMap["amount"] = s.Amount
	s.fieldMap["min_amount"] = s.MinAmount
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["end_time"] = s.EndTime
	s.fieldMap["total_count"] = s.TotalCount
	s.fieldMap["received_count"] = s.ReceivedCount
	s.fieldMap["used_count"] = s.UsedCount
	s.fieldMap["per_limit"] = s.PerLimit
	s.fieldMap["status"] = s.Status
	s.fieldMap["is_enabled"] = s.IsEnabled
	s.fieldMap["description"] = s.Description
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["is_deleted"] = s.IsDeleted
}

func (s smsCoupon) clone(db *gorm.DB) smsCoupon {
	s.smsCouponDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsCoupon) replaceDB(db *gorm.DB) smsCoupon {
	s.smsCouponDo.ReplaceDB(db)
	return s
}

type smsCouponDo struct{ gen.DO }

type ISmsCouponDo interface {
	gen.SubQuery
	Debug() ISmsCouponDo
	WithContext(ctx context.Context) ISmsCouponDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsCouponDo
	WriteDB() ISmsCouponDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsCouponDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsCouponDo
	Not(conds ...gen.Condition) ISmsCouponDo
	Or(conds ...gen.Condition) ISmsCouponDo
	Select(conds ...field.Expr) ISmsCouponDo
	Where(conds ...gen.Condition) ISmsCouponDo
	Order(conds ...field.Expr) ISmsCouponDo
	Distinct(cols ...field.Expr) ISmsCouponDo
	Omit(cols ...field.Expr) ISmsCouponDo
	Join(table schema.Tabler, on ...field.Expr) ISmsCouponDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsCouponDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsCouponDo
	Group(cols ...field.Expr) ISmsCouponDo
	Having(conds ...gen.Condition) ISmsCouponDo
	Limit(limit int) ISmsCouponDo
	Offset(offset int) ISmsCouponDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsCouponDo
	Unscoped() ISmsCouponDo
	Create(values ...*model.SmsCoupon) error
	CreateInBatches(values []*model.SmsCoupon, batchSize int) error
	Save(values ...*model.SmsCoupon) error
	First() (*model.SmsCoupon, error)
	Take() (*model.SmsCoupon, error)
	Last() (*model.SmsCoupon, error)
	Find() ([]*model.SmsCoupon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCoupon, err error)
	FindInBatches(result *[]*model.SmsCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsCoupon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsCouponDo
	Assign(attrs ...field.AssignExpr) ISmsCouponDo
	Joins(fields ...field.RelationField) ISmsCouponDo
	Preload(fields ...field.RelationField) ISmsCouponDo
	FirstOrInit() (*model.SmsCoupon, error)
	FirstOrCreate() (*model.SmsCoupon, error)
	FindByPage(offset int, limit int) (result []*model.SmsCoupon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsCouponDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsCouponDo) Debug() ISmsCouponDo {
	return s.withDO(s.DO.Debug())
}

func (s smsCouponDo) WithContext(ctx context.Context) ISmsCouponDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsCouponDo) ReadDB() ISmsCouponDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsCouponDo) WriteDB() ISmsCouponDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsCouponDo) Session(config *gorm.Session) ISmsCouponDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsCouponDo) Clauses(conds ...clause.Expression) ISmsCouponDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsCouponDo) Returning(value interface{}, columns ...string) ISmsCouponDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsCouponDo) Not(conds ...gen.Condition) ISmsCouponDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsCouponDo) Or(conds ...gen.Condition) ISmsCouponDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsCouponDo) Select(conds ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsCouponDo) Where(conds ...gen.Condition) ISmsCouponDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsCouponDo) Order(conds ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsCouponDo) Distinct(cols ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsCouponDo) Omit(cols ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsCouponDo) Join(table schema.Tabler, on ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsCouponDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsCouponDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsCouponDo) Group(cols ...field.Expr) ISmsCouponDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsCouponDo) Having(conds ...gen.Condition) ISmsCouponDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsCouponDo) Limit(limit int) ISmsCouponDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsCouponDo) Offset(offset int) ISmsCouponDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsCouponDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsCouponDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsCouponDo) Unscoped() ISmsCouponDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsCouponDo) Create(values ...*model.SmsCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsCouponDo) CreateInBatches(values []*model.SmsCoupon, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsCouponDo) Save(values ...*model.SmsCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsCouponDo) First() (*model.SmsCoupon, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCoupon), nil
	}
}

func (s smsCouponDo) Take() (*model.SmsCoupon, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCoupon), nil
	}
}

func (s smsCouponDo) Last() (*model.SmsCoupon, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCoupon), nil
	}
}

func (s smsCouponDo) Find() ([]*model.SmsCoupon, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsCoupon), err
}

func (s smsCouponDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCoupon, err error) {
	buf := make([]*model.SmsCoupon, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsCouponDo) FindInBatches(result *[]*model.SmsCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsCouponDo) Attrs(attrs ...field.AssignExpr) ISmsCouponDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsCouponDo) Assign(attrs ...field.AssignExpr) ISmsCouponDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsCouponDo) Joins(fields ...field.RelationField) ISmsCouponDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsCouponDo) Preload(fields ...field.RelationField) ISmsCouponDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsCouponDo) FirstOrInit() (*model.SmsCoupon, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCoupon), nil
	}
}

func (s smsCouponDo) FirstOrCreate() (*model.SmsCoupon, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCoupon), nil
	}
}

func (s smsCouponDo) FindByPage(offset int, limit int) (result []*model.SmsCoupon, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s smsCouponDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsCouponDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsCouponDo) Delete(models ...*model.SmsCoupon) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsCouponDo) withDO(do gen.Dao) *smsCouponDo {
	s.DO = *do.(*gen.DO)
	return s
}
