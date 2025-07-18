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

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
)

func newOmsOrderReturn(db *gorm.DB, opts ...gen.DOOption) omsOrderReturn {
	_omsOrderReturn := omsOrderReturn{}

	_omsOrderReturn.omsOrderReturnDo.UseDB(db, opts...)
	_omsOrderReturn.omsOrderReturnDo.UseModel(&model.OmsOrderReturn{})

	tableName := _omsOrderReturn.omsOrderReturnDo.TableName()
	_omsOrderReturn.ALL = field.NewAsterisk(tableName)
	_omsOrderReturn.ID = field.NewInt64(tableName, "id")
	_omsOrderReturn.OrderID = field.NewInt64(tableName, "order_id")
	_omsOrderReturn.ReturnNo = field.NewString(tableName, "return_no")
	_omsOrderReturn.MemberID = field.NewInt64(tableName, "member_id")
	_omsOrderReturn.Status = field.NewInt32(tableName, "status")
	_omsOrderReturn.Type = field.NewInt32(tableName, "type")
	_omsOrderReturn.Reason = field.NewString(tableName, "reason")
	_omsOrderReturn.Description = field.NewString(tableName, "description")
	_omsOrderReturn.ProofPic = field.NewString(tableName, "proof_pic")
	_omsOrderReturn.RefundAmount = field.NewFloat64(tableName, "refund_amount")
	_omsOrderReturn.ReturnName = field.NewString(tableName, "return_name")
	_omsOrderReturn.ReturnPhone = field.NewString(tableName, "return_phone")
	_omsOrderReturn.CompanyAddress = field.NewString(tableName, "company_address")
	_omsOrderReturn.CreateTime = field.NewTime(tableName, "create_time")
	_omsOrderReturn.HandleTime = field.NewTime(tableName, "handle_time")
	_omsOrderReturn.HandleNote = field.NewString(tableName, "handle_note")
	_omsOrderReturn.HandleMan = field.NewString(tableName, "handle_man")
	_omsOrderReturn.ReceiveTime = field.NewTime(tableName, "receive_time")
	_omsOrderReturn.ReceiveNote = field.NewString(tableName, "receive_note")
	_omsOrderReturn.ReceiveMan = field.NewString(tableName, "receive_man")
	_omsOrderReturn.RefundTime = field.NewTime(tableName, "refund_time")
	_omsOrderReturn.CloseTime = field.NewTime(tableName, "close_time")
	_omsOrderReturn.Remark = field.NewString(tableName, "remark")

	_omsOrderReturn.fillFieldMap()

	return _omsOrderReturn
}

// omsOrderReturn 退货/售后主表
type omsOrderReturn struct {
	omsOrderReturnDo omsOrderReturnDo

	ALL            field.Asterisk
	ID             field.Int64   // 主键ID
	OrderID        field.Int64   // 关联订单ID
	ReturnNo       field.String  // 退货单号
	MemberID       field.Int64   // 会员ID
	Status         field.Int32   // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
	Type           field.Int32   // 售后类型（0退货退款 1仅退款 2换货）
	Reason         field.String  // 退货原因
	Description    field.String  // 问题描述
	ProofPic       field.String  // 凭证图片，逗号分隔
	RefundAmount   field.Float64 // 退款金额
	ReturnName     field.String  // 退货人姓名
	ReturnPhone    field.String  // 退货人电话
	CompanyAddress field.String  // 退货收货地址
	CreateTime     field.Time    // 申请时间
	HandleTime     field.Time    // 处理时间
	HandleNote     field.String  // 处理备注
	HandleMan      field.String  // 处理人员
	ReceiveTime    field.Time    // 收货时间
	ReceiveNote    field.String  // 收货备注
	ReceiveMan     field.String  // 收货人
	RefundTime     field.Time    // 退款时间
	CloseTime      field.Time    // 关闭时间
	Remark         field.String  // 备注

	fieldMap map[string]field.Expr
}

func (o omsOrderReturn) Table(newTableName string) *omsOrderReturn {
	o.omsOrderReturnDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderReturn) As(alias string) *omsOrderReturn {
	o.omsOrderReturnDo.DO = *(o.omsOrderReturnDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderReturn) updateTableName(table string) *omsOrderReturn {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.OrderID = field.NewInt64(table, "order_id")
	o.ReturnNo = field.NewString(table, "return_no")
	o.MemberID = field.NewInt64(table, "member_id")
	o.Status = field.NewInt32(table, "status")
	o.Type = field.NewInt32(table, "type")
	o.Reason = field.NewString(table, "reason")
	o.Description = field.NewString(table, "description")
	o.ProofPic = field.NewString(table, "proof_pic")
	o.RefundAmount = field.NewFloat64(table, "refund_amount")
	o.ReturnName = field.NewString(table, "return_name")
	o.ReturnPhone = field.NewString(table, "return_phone")
	o.CompanyAddress = field.NewString(table, "company_address")
	o.CreateTime = field.NewTime(table, "create_time")
	o.HandleTime = field.NewTime(table, "handle_time")
	o.HandleNote = field.NewString(table, "handle_note")
	o.HandleMan = field.NewString(table, "handle_man")
	o.ReceiveTime = field.NewTime(table, "receive_time")
	o.ReceiveNote = field.NewString(table, "receive_note")
	o.ReceiveMan = field.NewString(table, "receive_man")
	o.RefundTime = field.NewTime(table, "refund_time")
	o.CloseTime = field.NewTime(table, "close_time")
	o.Remark = field.NewString(table, "remark")

	o.fillFieldMap()

	return o
}

func (o *omsOrderReturn) WithContext(ctx context.Context) IOmsOrderReturnDo {
	return o.omsOrderReturnDo.WithContext(ctx)
}

func (o omsOrderReturn) TableName() string { return o.omsOrderReturnDo.TableName() }

func (o omsOrderReturn) Alias() string { return o.omsOrderReturnDo.Alias() }

func (o omsOrderReturn) Columns(cols ...field.Expr) gen.Columns {
	return o.omsOrderReturnDo.Columns(cols...)
}

func (o *omsOrderReturn) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderReturn) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 23)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["return_no"] = o.ReturnNo
	o.fieldMap["member_id"] = o.MemberID
	o.fieldMap["status"] = o.Status
	o.fieldMap["type"] = o.Type
	o.fieldMap["reason"] = o.Reason
	o.fieldMap["description"] = o.Description
	o.fieldMap["proof_pic"] = o.ProofPic
	o.fieldMap["refund_amount"] = o.RefundAmount
	o.fieldMap["return_name"] = o.ReturnName
	o.fieldMap["return_phone"] = o.ReturnPhone
	o.fieldMap["company_address"] = o.CompanyAddress
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["handle_time"] = o.HandleTime
	o.fieldMap["handle_note"] = o.HandleNote
	o.fieldMap["handle_man"] = o.HandleMan
	o.fieldMap["receive_time"] = o.ReceiveTime
	o.fieldMap["receive_note"] = o.ReceiveNote
	o.fieldMap["receive_man"] = o.ReceiveMan
	o.fieldMap["refund_time"] = o.RefundTime
	o.fieldMap["close_time"] = o.CloseTime
	o.fieldMap["remark"] = o.Remark
}

func (o omsOrderReturn) clone(db *gorm.DB) omsOrderReturn {
	o.omsOrderReturnDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderReturn) replaceDB(db *gorm.DB) omsOrderReturn {
	o.omsOrderReturnDo.ReplaceDB(db)
	return o
}

type omsOrderReturnDo struct{ gen.DO }

type IOmsOrderReturnDo interface {
	gen.SubQuery
	Debug() IOmsOrderReturnDo
	WithContext(ctx context.Context) IOmsOrderReturnDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderReturnDo
	WriteDB() IOmsOrderReturnDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderReturnDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderReturnDo
	Not(conds ...gen.Condition) IOmsOrderReturnDo
	Or(conds ...gen.Condition) IOmsOrderReturnDo
	Select(conds ...field.Expr) IOmsOrderReturnDo
	Where(conds ...gen.Condition) IOmsOrderReturnDo
	Order(conds ...field.Expr) IOmsOrderReturnDo
	Distinct(cols ...field.Expr) IOmsOrderReturnDo
	Omit(cols ...field.Expr) IOmsOrderReturnDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo
	Group(cols ...field.Expr) IOmsOrderReturnDo
	Having(conds ...gen.Condition) IOmsOrderReturnDo
	Limit(limit int) IOmsOrderReturnDo
	Offset(offset int) IOmsOrderReturnDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderReturnDo
	Unscoped() IOmsOrderReturnDo
	Create(values ...*model.OmsOrderReturn) error
	CreateInBatches(values []*model.OmsOrderReturn, batchSize int) error
	Save(values ...*model.OmsOrderReturn) error
	First() (*model.OmsOrderReturn, error)
	Take() (*model.OmsOrderReturn, error)
	Last() (*model.OmsOrderReturn, error)
	Find() ([]*model.OmsOrderReturn, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderReturn, err error)
	FindInBatches(result *[]*model.OmsOrderReturn, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderReturn) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderReturnDo
	Assign(attrs ...field.AssignExpr) IOmsOrderReturnDo
	Joins(fields ...field.RelationField) IOmsOrderReturnDo
	Preload(fields ...field.RelationField) IOmsOrderReturnDo
	FirstOrInit() (*model.OmsOrderReturn, error)
	FirstOrCreate() (*model.OmsOrderReturn, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderReturn, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderReturnDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderReturnDo) Debug() IOmsOrderReturnDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderReturnDo) WithContext(ctx context.Context) IOmsOrderReturnDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderReturnDo) ReadDB() IOmsOrderReturnDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderReturnDo) WriteDB() IOmsOrderReturnDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderReturnDo) Session(config *gorm.Session) IOmsOrderReturnDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderReturnDo) Clauses(conds ...clause.Expression) IOmsOrderReturnDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderReturnDo) Returning(value interface{}, columns ...string) IOmsOrderReturnDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderReturnDo) Not(conds ...gen.Condition) IOmsOrderReturnDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderReturnDo) Or(conds ...gen.Condition) IOmsOrderReturnDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderReturnDo) Select(conds ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderReturnDo) Where(conds ...gen.Condition) IOmsOrderReturnDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderReturnDo) Order(conds ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderReturnDo) Distinct(cols ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderReturnDo) Omit(cols ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderReturnDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderReturnDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderReturnDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderReturnDo) Group(cols ...field.Expr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderReturnDo) Having(conds ...gen.Condition) IOmsOrderReturnDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderReturnDo) Limit(limit int) IOmsOrderReturnDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderReturnDo) Offset(offset int) IOmsOrderReturnDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderReturnDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderReturnDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderReturnDo) Unscoped() IOmsOrderReturnDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderReturnDo) Create(values ...*model.OmsOrderReturn) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderReturnDo) CreateInBatches(values []*model.OmsOrderReturn, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderReturnDo) Save(values ...*model.OmsOrderReturn) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderReturnDo) First() (*model.OmsOrderReturn, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturn), nil
	}
}

func (o omsOrderReturnDo) Take() (*model.OmsOrderReturn, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturn), nil
	}
}

func (o omsOrderReturnDo) Last() (*model.OmsOrderReturn, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturn), nil
	}
}

func (o omsOrderReturnDo) Find() ([]*model.OmsOrderReturn, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderReturn), err
}

func (o omsOrderReturnDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderReturn, err error) {
	buf := make([]*model.OmsOrderReturn, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderReturnDo) FindInBatches(result *[]*model.OmsOrderReturn, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderReturnDo) Attrs(attrs ...field.AssignExpr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderReturnDo) Assign(attrs ...field.AssignExpr) IOmsOrderReturnDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderReturnDo) Joins(fields ...field.RelationField) IOmsOrderReturnDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderReturnDo) Preload(fields ...field.RelationField) IOmsOrderReturnDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderReturnDo) FirstOrInit() (*model.OmsOrderReturn, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturn), nil
	}
}

func (o omsOrderReturnDo) FirstOrCreate() (*model.OmsOrderReturn, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderReturn), nil
	}
}

func (o omsOrderReturnDo) FindByPage(offset int, limit int) (result []*model.OmsOrderReturn, count int64, err error) {
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

func (o omsOrderReturnDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderReturnDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderReturnDo) Delete(models ...*model.OmsOrderReturn) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderReturnDo) withDO(do gen.Dao) *omsOrderReturnDo {
	o.DO = *do.(*gen.DO)
	return o
}
