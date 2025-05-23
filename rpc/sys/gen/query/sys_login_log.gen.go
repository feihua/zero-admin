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

	"github.com/feihua/zero-admin/rpc/sys/gen/model"
)

func newSysLoginLog(db *gorm.DB, opts ...gen.DOOption) sysLoginLog {
	_sysLoginLog := sysLoginLog{}

	_sysLoginLog.sysLoginLogDo.UseDB(db, opts...)
	_sysLoginLog.sysLoginLogDo.UseModel(&model.SysLoginLog{})

	tableName := _sysLoginLog.sysLoginLogDo.TableName()
	_sysLoginLog.ALL = field.NewAsterisk(tableName)
	_sysLoginLog.ID = field.NewInt64(tableName, "id")
	_sysLoginLog.LoginName = field.NewString(tableName, "login_name")
	_sysLoginLog.Ipaddr = field.NewString(tableName, "ipaddr")
	_sysLoginLog.LoginLocation = field.NewString(tableName, "login_location")
	_sysLoginLog.Platform = field.NewString(tableName, "platform")
	_sysLoginLog.Browser = field.NewString(tableName, "browser")
	_sysLoginLog.Version = field.NewString(tableName, "version")
	_sysLoginLog.Os = field.NewString(tableName, "os")
	_sysLoginLog.Arch = field.NewString(tableName, "arch")
	_sysLoginLog.Engine = field.NewString(tableName, "engine")
	_sysLoginLog.EngineDetails = field.NewString(tableName, "engine_details")
	_sysLoginLog.Extra = field.NewString(tableName, "extra")
	_sysLoginLog.Status = field.NewInt32(tableName, "status")
	_sysLoginLog.Msg = field.NewString(tableName, "msg")
	_sysLoginLog.LoginTime = field.NewTime(tableName, "login_time")

	_sysLoginLog.fillFieldMap()

	return _sysLoginLog
}

// sysLoginLog 系统登录日志表
type sysLoginLog struct {
	sysLoginLogDo sysLoginLogDo

	ALL           field.Asterisk
	ID            field.Int64  // 登录日志id
	LoginName     field.String // 登录账号
	Ipaddr        field.String // 登录IP地址
	LoginLocation field.String // 登录地点
	Platform      field.String // 平台信息
	Browser       field.String // 浏览器类型
	Version       field.String // 浏览器版本
	Os            field.String // 操作系统
	Arch          field.String // 体系结构信息
	Engine        field.String // 渲染引擎信息
	EngineDetails field.String // 渲染引擎详细信息
	Extra         field.String // 其他信息（可选）
	Status        field.Int32  // 登录状态(0:失败,1:成功)
	Msg           field.String // 提示消息
	LoginTime     field.Time   // 访问时间

	fieldMap map[string]field.Expr
}

func (s sysLoginLog) Table(newTableName string) *sysLoginLog {
	s.sysLoginLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysLoginLog) As(alias string) *sysLoginLog {
	s.sysLoginLogDo.DO = *(s.sysLoginLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysLoginLog) updateTableName(table string) *sysLoginLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.LoginName = field.NewString(table, "login_name")
	s.Ipaddr = field.NewString(table, "ipaddr")
	s.LoginLocation = field.NewString(table, "login_location")
	s.Platform = field.NewString(table, "platform")
	s.Browser = field.NewString(table, "browser")
	s.Version = field.NewString(table, "version")
	s.Os = field.NewString(table, "os")
	s.Arch = field.NewString(table, "arch")
	s.Engine = field.NewString(table, "engine")
	s.EngineDetails = field.NewString(table, "engine_details")
	s.Extra = field.NewString(table, "extra")
	s.Status = field.NewInt32(table, "status")
	s.Msg = field.NewString(table, "msg")
	s.LoginTime = field.NewTime(table, "login_time")

	s.fillFieldMap()

	return s
}

func (s *sysLoginLog) WithContext(ctx context.Context) ISysLoginLogDo {
	return s.sysLoginLogDo.WithContext(ctx)
}

func (s sysLoginLog) TableName() string { return s.sysLoginLogDo.TableName() }

func (s sysLoginLog) Alias() string { return s.sysLoginLogDo.Alias() }

func (s sysLoginLog) Columns(cols ...field.Expr) gen.Columns { return s.sysLoginLogDo.Columns(cols...) }

func (s *sysLoginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysLoginLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["login_name"] = s.LoginName
	s.fieldMap["ipaddr"] = s.Ipaddr
	s.fieldMap["login_location"] = s.LoginLocation
	s.fieldMap["platform"] = s.Platform
	s.fieldMap["browser"] = s.Browser
	s.fieldMap["version"] = s.Version
	s.fieldMap["os"] = s.Os
	s.fieldMap["arch"] = s.Arch
	s.fieldMap["engine"] = s.Engine
	s.fieldMap["engine_details"] = s.EngineDetails
	s.fieldMap["extra"] = s.Extra
	s.fieldMap["status"] = s.Status
	s.fieldMap["msg"] = s.Msg
	s.fieldMap["login_time"] = s.LoginTime
}

func (s sysLoginLog) clone(db *gorm.DB) sysLoginLog {
	s.sysLoginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysLoginLog) replaceDB(db *gorm.DB) sysLoginLog {
	s.sysLoginLogDo.ReplaceDB(db)
	return s
}

type sysLoginLogDo struct{ gen.DO }

type ISysLoginLogDo interface {
	gen.SubQuery
	Debug() ISysLoginLogDo
	WithContext(ctx context.Context) ISysLoginLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysLoginLogDo
	WriteDB() ISysLoginLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysLoginLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysLoginLogDo
	Not(conds ...gen.Condition) ISysLoginLogDo
	Or(conds ...gen.Condition) ISysLoginLogDo
	Select(conds ...field.Expr) ISysLoginLogDo
	Where(conds ...gen.Condition) ISysLoginLogDo
	Order(conds ...field.Expr) ISysLoginLogDo
	Distinct(cols ...field.Expr) ISysLoginLogDo
	Omit(cols ...field.Expr) ISysLoginLogDo
	Join(table schema.Tabler, on ...field.Expr) ISysLoginLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysLoginLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysLoginLogDo
	Group(cols ...field.Expr) ISysLoginLogDo
	Having(conds ...gen.Condition) ISysLoginLogDo
	Limit(limit int) ISysLoginLogDo
	Offset(offset int) ISysLoginLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLoginLogDo
	Unscoped() ISysLoginLogDo
	Create(values ...*model.SysLoginLog) error
	CreateInBatches(values []*model.SysLoginLog, batchSize int) error
	Save(values ...*model.SysLoginLog) error
	First() (*model.SysLoginLog, error)
	Take() (*model.SysLoginLog, error)
	Last() (*model.SysLoginLog, error)
	Find() ([]*model.SysLoginLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLoginLog, err error)
	FindInBatches(result *[]*model.SysLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysLoginLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysLoginLogDo
	Assign(attrs ...field.AssignExpr) ISysLoginLogDo
	Joins(fields ...field.RelationField) ISysLoginLogDo
	Preload(fields ...field.RelationField) ISysLoginLogDo
	FirstOrInit() (*model.SysLoginLog, error)
	FirstOrCreate() (*model.SysLoginLog, error)
	FindByPage(offset int, limit int) (result []*model.SysLoginLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysLoginLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysLoginLogDo) Debug() ISysLoginLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysLoginLogDo) WithContext(ctx context.Context) ISysLoginLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysLoginLogDo) ReadDB() ISysLoginLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysLoginLogDo) WriteDB() ISysLoginLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysLoginLogDo) Session(config *gorm.Session) ISysLoginLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysLoginLogDo) Clauses(conds ...clause.Expression) ISysLoginLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysLoginLogDo) Returning(value interface{}, columns ...string) ISysLoginLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysLoginLogDo) Not(conds ...gen.Condition) ISysLoginLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysLoginLogDo) Or(conds ...gen.Condition) ISysLoginLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysLoginLogDo) Select(conds ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysLoginLogDo) Where(conds ...gen.Condition) ISysLoginLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysLoginLogDo) Order(conds ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysLoginLogDo) Distinct(cols ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysLoginLogDo) Omit(cols ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysLoginLogDo) Join(table schema.Tabler, on ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysLoginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysLoginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysLoginLogDo) Group(cols ...field.Expr) ISysLoginLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysLoginLogDo) Having(conds ...gen.Condition) ISysLoginLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysLoginLogDo) Limit(limit int) ISysLoginLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysLoginLogDo) Offset(offset int) ISysLoginLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysLoginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLoginLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysLoginLogDo) Unscoped() ISysLoginLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysLoginLogDo) Create(values ...*model.SysLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysLoginLogDo) CreateInBatches(values []*model.SysLoginLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysLoginLogDo) Save(values ...*model.SysLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysLoginLogDo) First() (*model.SysLoginLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLoginLog), nil
	}
}

func (s sysLoginLogDo) Take() (*model.SysLoginLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLoginLog), nil
	}
}

func (s sysLoginLogDo) Last() (*model.SysLoginLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLoginLog), nil
	}
}

func (s sysLoginLogDo) Find() ([]*model.SysLoginLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysLoginLog), err
}

func (s sysLoginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLoginLog, err error) {
	buf := make([]*model.SysLoginLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysLoginLogDo) FindInBatches(result *[]*model.SysLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysLoginLogDo) Attrs(attrs ...field.AssignExpr) ISysLoginLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysLoginLogDo) Assign(attrs ...field.AssignExpr) ISysLoginLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysLoginLogDo) Joins(fields ...field.RelationField) ISysLoginLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysLoginLogDo) Preload(fields ...field.RelationField) ISysLoginLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysLoginLogDo) FirstOrInit() (*model.SysLoginLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLoginLog), nil
	}
}

func (s sysLoginLogDo) FirstOrCreate() (*model.SysLoginLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLoginLog), nil
	}
}

func (s sysLoginLogDo) FindByPage(offset int, limit int) (result []*model.SysLoginLog, count int64, err error) {
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

func (s sysLoginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysLoginLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysLoginLogDo) Delete(models ...*model.SysLoginLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysLoginLogDo) withDO(do gen.Dao) *sysLoginLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
