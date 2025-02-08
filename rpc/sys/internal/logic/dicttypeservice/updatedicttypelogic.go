package dicttypeservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateDictTypeLogic 更新字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type UpdateDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeLogic {
	return &UpdateDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateDictType 更新字典表
// 1.根据字典id查询字典是否已存在
// 2.查询字典名称是否已存在,如果字典名称已存在,则直接返回
// 3.查询字典类型是否已存在,如果字典类型已存在,则直接返回
// 4.字典存在时,则直接更新字典
func (l *UpdateDictTypeLogic) UpdateDictType(in *sysclient.UpdateDictTypeReq) (*sysclient.UpdateDictTypeResp, error) {
	dictType := query.SysDictType
	q := dictType.WithContext(l.ctx)

	// 1.根据字典id查询字典是否已存在
	res, err := query.SysDictType.WithContext(l.ctx).Where(query.SysDictType.ID.Eq(in.Id)).First()

	// 1.判断字典类型是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "字典类型不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("字典类型不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询字典类型异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询字典类型异常")
	}

	// 2.查询字典名称是否已存在,如果字典名称已存在,则直接返回
	count, err := q.WithContext(l.ctx).Where(dictType.ID.Neq(in.Id), dictType.DictName.Eq(in.DictName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, ".查询字典名称失败,参数：%+v,,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新字典信息失败"))
	}

	if count > 0 {
		return nil, errors.New("更新字典类型失败,字典名称已存在")
	}

	// 3.查询字典类型是否已存在,如果字典类型已存在,则直接返回
	count, err = q.WithContext(l.ctx).Where(dictType.ID.Neq(in.Id), dictType.DictType.Eq(in.DictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	if count > 0 {
		return nil, errors.New("更新字典类型失败,字典类型已存在")
	}

	// 4.字典存在时,则直接更新字典
	now := time.Now()
	dict := &model.SysDictType{
		ID:         in.Id,          // 编号
		DictName:   in.DictName,    // 字典名称
		DictType:   in.DictType,    // 字典类型
		DictStatus: in.DictStatus,  // 字典状态
		Remark:     in.Remark,      // 备注信息
		IsSystem:   in.IsSystem,    // 是否系统预留  0：否  1：是
		CreateBy:   res.CreateBy,   // 创建者
		CreateTime: res.CreateTime, // 创建时间
		UpdateBy:   in.UpdateBy,    // 更新者
		UpdateTime: &now,           // 更新时间
	}
	err = l.svcCtx.DB.Model(&model.SysDictType{}).WithContext(l.ctx).Where(dictType.ID.Eq(in.Id)).Save(dict).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("更新字典信息失败")
	}

	return &sysclient.UpdateDictTypeResp{}, nil
}
