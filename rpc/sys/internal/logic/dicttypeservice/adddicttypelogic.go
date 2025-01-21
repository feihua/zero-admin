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
)

// AddDictTypeLogic 添加字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type AddDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictTypeLogic {
	return &AddDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDictType 添加字典信息
// 1.查询字典名称是否已存在,如果字典名称已存在,则直接返回
// 2.查询字典类型是否已存在,如果字典类型已存在,则直接返回
// 3.字典不存在时,则直接添加字典
func (l *AddDictTypeLogic) AddDictType(in *sysclient.AddDictTypeReq) (*sysclient.AddDictTypeResp, error) {
	q := query.SysDictType

	// 1.查询字典名称是否已存在,如果字典名称已存在,则直接返回
	count, err := q.WithContext(l.ctx).Where(q.DictName.Eq(in.DictName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, ".查询字典名称失败,参数：%+v,,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加字典信息"))
	}

	if count > 0 {
		return nil, errors.New("添加字典类型失败,字典名称已存在")
	}

	// 2.查询字典类型是否已存在,如果字典类型已存在,则直接返回
	count, err = q.WithContext(l.ctx).Where(q.DictType.Eq(in.DictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	if count > 0 {
		return nil, errors.New("添加字典类型失败,字典类型已存在")
	}

	dict := &model.SysDictType{
		DictName:   in.DictName,   // 字典名称
		DictType:   in.DictType,   // 字典类型
		DictStatus: in.DictStatus, // 字典状态
		Remark:     in.Remark,     // 备注信息
		IsSystem:   in.IsSystem,   // 是否系统预留  0：否  1：是
		IsDeleted:  0,             // 是否删除  0：否  1：是
		CreateBy:   in.CreateBy,   // 创建者
	}
	err = q.WithContext(l.ctx).Create(dict)

	if err != nil {
		logc.Errorf(l.ctx, "添加字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("添加字典信息失败")
	}

	return &sysclient.AddDictTypeResp{}, nil
}
