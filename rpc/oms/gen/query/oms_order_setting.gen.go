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

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
)

func newOmsOrderSetting(db *gorm.DB, opts ...gen.DOOption) omsOrderSetting {
	_omsOrderSetting := omsOrderSetting{}

	_omsOrderSetting.omsOrderSettingDo.UseDB(db, opts...)
	_omsOrderSetting.omsOrderSettingDo.UseModel(&model.OmsOrderSetting{})

	tableName := _omsOrderSetting.omsOrderSettingDo.TableName()
	_omsOrderSetting.ALL = field.NewAsterisk(tableName)
	_omsOrderSetting.ID = field.NewInt64(tableName, "id")
	_omsOrderSetting.FlashOrderOvertime = field.NewInt32(tableName, "flash_order_overtime")
	_omsOrderSetting.NormalOrderOvertime = field.NewInt32(tableName, "normal_order_overtime")
	_omsOrderSetting.ConfirmOvertime = field.NewInt32(tableName, "confirm_overtime")
	_omsOrderSetting.FinishOvertime = field.NewInt32(tableName, "finish_overtime")
	_omsOrderSetting.CommentOvertime = field.NewInt32(tableName, "comment_overtime")

	_omsOrderSetting.fillFieldMap()

	return _omsOrderSetting
}

// omsOrderSetting 订单设置表
type omsOrderSetting struct {
	omsOrderSettingDo omsOrderSettingDo

	ALL                 field.Asterisk
	ID                  field.Int64
	FlashOrderOvertime  field.Int32 // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime field.Int32 // 正常订单超时时间(分)
	ConfirmOvertime     field.Int32 // 发货后自动确认收货时间（天）
	FinishOvertime      field.Int32 // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     field.Int32 // 订单完成后自动好评时间（天）

	fieldMap map[string]field.Expr
}

func (o omsOrderSetting) Table(newTableName string) *omsOrderSetting {
	o.omsOrderSettingDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderSetting) As(alias string) *omsOrderSetting {
	o.omsOrderSettingDo.DO = *(o.omsOrderSettingDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderSetting) updateTableName(table string) *omsOrderSetting {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.FlashOrderOvertime = field.NewInt32(table, "flash_order_overtime")
	o.NormalOrderOvertime = field.NewInt32(table, "normal_order_overtime")
	o.ConfirmOvertime = field.NewInt32(table, "confirm_overtime")
	o.FinishOvertime = field.NewInt32(table, "finish_overtime")
	o.CommentOvertime = field.NewInt32(table, "comment_overtime")

	o.fillFieldMap()

	return o
}

func (o *omsOrderSetting) WithContext(ctx context.Context) IOmsOrderSettingDo {
	return o.omsOrderSettingDo.WithContext(ctx)
}

func (o omsOrderSetting) TableName() string { return o.omsOrderSettingDo.TableName() }

func (o omsOrderSetting) Alias() string { return o.omsOrderSettingDo.Alias() }

func (o omsOrderSetting) Columns(cols ...field.Expr) gen.Columns {
	return o.omsOrderSettingDo.Columns(cols...)
}

func (o *omsOrderSetting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderSetting) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["flash_order_overtime"] = o.FlashOrderOvertime
	o.fieldMap["normal_order_overtime"] = o.NormalOrderOvertime
	o.fieldMap["confirm_overtime"] = o.ConfirmOvertime
	o.fieldMap["finish_overtime"] = o.FinishOvertime
	o.fieldMap["comment_overtime"] = o.CommentOvertime
}

func (o omsOrderSetting) clone(db *gorm.DB) omsOrderSetting {
	o.omsOrderSettingDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderSetting) replaceDB(db *gorm.DB) omsOrderSetting {
	o.omsOrderSettingDo.ReplaceDB(db)
	return o
}

type omsOrderSettingDo struct{ gen.DO }

type IOmsOrderSettingDo interface {
	gen.SubQuery
	Debug() IOmsOrderSettingDo
	WithContext(ctx context.Context) IOmsOrderSettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderSettingDo
	WriteDB() IOmsOrderSettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderSettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderSettingDo
	Not(conds ...gen.Condition) IOmsOrderSettingDo
	Or(conds ...gen.Condition) IOmsOrderSettingDo
	Select(conds ...field.Expr) IOmsOrderSettingDo
	Where(conds ...gen.Condition) IOmsOrderSettingDo
	Order(conds ...field.Expr) IOmsOrderSettingDo
	Distinct(cols ...field.Expr) IOmsOrderSettingDo
	Omit(cols ...field.Expr) IOmsOrderSettingDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo
	Group(cols ...field.Expr) IOmsOrderSettingDo
	Having(conds ...gen.Condition) IOmsOrderSettingDo
	Limit(limit int) IOmsOrderSettingDo
	Offset(offset int) IOmsOrderSettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderSettingDo
	Unscoped() IOmsOrderSettingDo
	Create(values ...*model.OmsOrderSetting) error
	CreateInBatches(values []*model.OmsOrderSetting, batchSize int) error
	Save(values ...*model.OmsOrderSetting) error
	First() (*model.OmsOrderSetting, error)
	Take() (*model.OmsOrderSetting, error)
	Last() (*model.OmsOrderSetting, error)
	Find() ([]*model.OmsOrderSetting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderSetting, err error)
	FindInBatches(result *[]*model.OmsOrderSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderSetting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderSettingDo
	Assign(attrs ...field.AssignExpr) IOmsOrderSettingDo
	Joins(fields ...field.RelationField) IOmsOrderSettingDo
	Preload(fields ...field.RelationField) IOmsOrderSettingDo
	FirstOrInit() (*model.OmsOrderSetting, error)
	FirstOrCreate() (*model.OmsOrderSetting, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderSetting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderSettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderSettingDo) Debug() IOmsOrderSettingDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderSettingDo) WithContext(ctx context.Context) IOmsOrderSettingDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderSettingDo) ReadDB() IOmsOrderSettingDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderSettingDo) WriteDB() IOmsOrderSettingDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderSettingDo) Session(config *gorm.Session) IOmsOrderSettingDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderSettingDo) Clauses(conds ...clause.Expression) IOmsOrderSettingDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderSettingDo) Returning(value interface{}, columns ...string) IOmsOrderSettingDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderSettingDo) Not(conds ...gen.Condition) IOmsOrderSettingDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderSettingDo) Or(conds ...gen.Condition) IOmsOrderSettingDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderSettingDo) Select(conds ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderSettingDo) Where(conds ...gen.Condition) IOmsOrderSettingDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderSettingDo) Order(conds ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderSettingDo) Distinct(cols ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderSettingDo) Omit(cols ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderSettingDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderSettingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderSettingDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderSettingDo) Group(cols ...field.Expr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderSettingDo) Having(conds ...gen.Condition) IOmsOrderSettingDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderSettingDo) Limit(limit int) IOmsOrderSettingDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderSettingDo) Offset(offset int) IOmsOrderSettingDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderSettingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderSettingDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderSettingDo) Unscoped() IOmsOrderSettingDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderSettingDo) Create(values ...*model.OmsOrderSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderSettingDo) CreateInBatches(values []*model.OmsOrderSetting, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderSettingDo) Save(values ...*model.OmsOrderSetting) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderSettingDo) First() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Take() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Last() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) Find() ([]*model.OmsOrderSetting, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderSetting), err
}

func (o omsOrderSettingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderSetting, err error) {
	buf := make([]*model.OmsOrderSetting, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderSettingDo) FindInBatches(result *[]*model.OmsOrderSetting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderSettingDo) Attrs(attrs ...field.AssignExpr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderSettingDo) Assign(attrs ...field.AssignExpr) IOmsOrderSettingDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderSettingDo) Joins(fields ...field.RelationField) IOmsOrderSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderSettingDo) Preload(fields ...field.RelationField) IOmsOrderSettingDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderSettingDo) FirstOrInit() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) FirstOrCreate() (*model.OmsOrderSetting, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderSetting), nil
	}
}

func (o omsOrderSettingDo) FindByPage(offset int, limit int) (result []*model.OmsOrderSetting, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o omsOrderSettingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderSettingDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderSettingDo) Delete(models ...*model.OmsOrderSetting) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderSettingDo) withDO(do gen.Dao) *omsOrderSettingDo {
	o.DO = *do.(*gen.DO)
	return o
}