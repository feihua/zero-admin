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

func newSysRole(db *gorm.DB, opts ...gen.DOOption) sysRole {
	_sysRole := sysRole{}

	_sysRole.sysRoleDo.UseDB(db, opts...)
	_sysRole.sysRoleDo.UseModel(&model.SysRole{})

	tableName := _sysRole.sysRoleDo.TableName()
	_sysRole.ALL = field.NewAsterisk(tableName)
	_sysRole.ID = field.NewInt64(tableName, "id")
	_sysRole.RoleName = field.NewString(tableName, "role_name")
	_sysRole.RoleKey = field.NewString(tableName, "role_key")
	_sysRole.DataScope = field.NewInt32(tableName, "data_scope")
	_sysRole.Status = field.NewInt32(tableName, "status")
	_sysRole.Remark = field.NewString(tableName, "remark")
	_sysRole.DelFlag = field.NewInt32(tableName, "del_flag")
	_sysRole.CreateBy = field.NewString(tableName, "create_by")
	_sysRole.CreateTime = field.NewTime(tableName, "create_time")
	_sysRole.UpdateBy = field.NewString(tableName, "update_by")
	_sysRole.UpdateTime = field.NewTime(tableName, "update_time")

	_sysRole.fillFieldMap()

	return _sysRole
}

// sysRole 角色信息
type sysRole struct {
	sysRoleDo sysRoleDo

	ALL        field.Asterisk
	ID         field.Int64  // 角色id
	RoleName   field.String // 名称
	RoleKey    field.String // 角色权限字符串
	DataScope  field.Int32  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status     field.Int32  // 状态(1:正常，0:禁用)
	Remark     field.String // 备注
	DelFlag    field.Int32  // 删除标志（0代表删除 1代表存在）
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (s sysRole) Table(newTableName string) *sysRole {
	s.sysRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysRole) As(alias string) *sysRole {
	s.sysRoleDo.DO = *(s.sysRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysRole) updateTableName(table string) *sysRole {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.RoleName = field.NewString(table, "role_name")
	s.RoleKey = field.NewString(table, "role_key")
	s.DataScope = field.NewInt32(table, "data_scope")
	s.Status = field.NewInt32(table, "status")
	s.Remark = field.NewString(table, "remark")
	s.DelFlag = field.NewInt32(table, "del_flag")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")

	s.fillFieldMap()

	return s
}

func (s *sysRole) WithContext(ctx context.Context) ISysRoleDo { return s.sysRoleDo.WithContext(ctx) }

func (s sysRole) TableName() string { return s.sysRoleDo.TableName() }

func (s sysRole) Alias() string { return s.sysRoleDo.Alias() }

func (s sysRole) Columns(cols ...field.Expr) gen.Columns { return s.sysRoleDo.Columns(cols...) }

func (s *sysRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["role_name"] = s.RoleName
	s.fieldMap["role_key"] = s.RoleKey
	s.fieldMap["data_scope"] = s.DataScope
	s.fieldMap["status"] = s.Status
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["del_flag"] = s.DelFlag
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
}

func (s sysRole) clone(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysRole) replaceDB(db *gorm.DB) sysRole {
	s.sysRoleDo.ReplaceDB(db)
	return s
}

type sysRoleDo struct{ gen.DO }

type ISysRoleDo interface {
	gen.SubQuery
	Debug() ISysRoleDo
	WithContext(ctx context.Context) ISysRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysRoleDo
	WriteDB() ISysRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysRoleDo
	Not(conds ...gen.Condition) ISysRoleDo
	Or(conds ...gen.Condition) ISysRoleDo
	Select(conds ...field.Expr) ISysRoleDo
	Where(conds ...gen.Condition) ISysRoleDo
	Order(conds ...field.Expr) ISysRoleDo
	Distinct(cols ...field.Expr) ISysRoleDo
	Omit(cols ...field.Expr) ISysRoleDo
	Join(table schema.Tabler, on ...field.Expr) ISysRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleDo
	Group(cols ...field.Expr) ISysRoleDo
	Having(conds ...gen.Condition) ISysRoleDo
	Limit(limit int) ISysRoleDo
	Offset(offset int) ISysRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleDo
	Unscoped() ISysRoleDo
	Create(values ...*model.SysRole) error
	CreateInBatches(values []*model.SysRole, batchSize int) error
	Save(values ...*model.SysRole) error
	First() (*model.SysRole, error)
	Take() (*model.SysRole, error)
	Last() (*model.SysRole, error)
	Find() ([]*model.SysRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRole, err error)
	FindInBatches(result *[]*model.SysRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysRoleDo
	Assign(attrs ...field.AssignExpr) ISysRoleDo
	Joins(fields ...field.RelationField) ISysRoleDo
	Preload(fields ...field.RelationField) ISysRoleDo
	FirstOrInit() (*model.SysRole, error)
	FirstOrCreate() (*model.SysRole, error)
	FindByPage(offset int, limit int) (result []*model.SysRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysRoleDo) Debug() ISysRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysRoleDo) WithContext(ctx context.Context) ISysRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysRoleDo) ReadDB() ISysRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysRoleDo) WriteDB() ISysRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysRoleDo) Session(config *gorm.Session) ISysRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysRoleDo) Clauses(conds ...clause.Expression) ISysRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysRoleDo) Returning(value interface{}, columns ...string) ISysRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysRoleDo) Not(conds ...gen.Condition) ISysRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysRoleDo) Or(conds ...gen.Condition) ISysRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysRoleDo) Select(conds ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysRoleDo) Where(conds ...gen.Condition) ISysRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysRoleDo) Order(conds ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysRoleDo) Distinct(cols ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysRoleDo) Omit(cols ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysRoleDo) Join(table schema.Tabler, on ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysRoleDo) Group(cols ...field.Expr) ISysRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysRoleDo) Having(conds ...gen.Condition) ISysRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysRoleDo) Limit(limit int) ISysRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysRoleDo) Offset(offset int) ISysRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysRoleDo) Unscoped() ISysRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysRoleDo) Create(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysRoleDo) CreateInBatches(values []*model.SysRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysRoleDo) Save(values ...*model.SysRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysRoleDo) First() (*model.SysRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Take() (*model.SysRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Last() (*model.SysRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) Find() ([]*model.SysRole, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysRole), err
}

func (s sysRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysRole, err error) {
	buf := make([]*model.SysRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysRoleDo) FindInBatches(result *[]*model.SysRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysRoleDo) Attrs(attrs ...field.AssignExpr) ISysRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysRoleDo) Assign(attrs ...field.AssignExpr) ISysRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysRoleDo) Joins(fields ...field.RelationField) ISysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysRoleDo) Preload(fields ...field.RelationField) ISysRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysRoleDo) FirstOrInit() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FirstOrCreate() (*model.SysRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysRole), nil
	}
}

func (s sysRoleDo) FindByPage(offset int, limit int) (result []*model.SysRole, count int64, err error) {
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

func (s sysRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysRoleDo) Delete(models ...*model.SysRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysRoleDo) withDO(do gen.Dao) *sysRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}
