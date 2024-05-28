package dictservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddDictLogic 添加字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type AddDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictLogic {
	return &AddDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddDict 添加字典信息
// 1.根据字典名称或者类型查询字典是否已存在
// 2.如果字典已存在,则直接返回
// 3.字典不存在时,则直接添加字典
func (l *AddDictLogic) AddDict(in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	q := query.SysDict

	// 1.根据字典名称或者类型查询字典是否已存在
	count, err := q.WithContext(l.ctx).Where(q.DictName.Eq(in.DictName)).Or(q.DictType.Eq(in.DictType)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据字典名称：%+v,查询字典信息失败,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("查询字典信息失败"))
	}

	//2.如果字典已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "字典信息已存在：%+v", in)
		return nil, errors.New("字典已存在")
	}

	dict := &model.SysDict{
		DictName:   in.DictName,
		DictType:   in.DictType,
		DictStatus: in.DictStatus,
		Remark:     in.Remark,
		IsSystem:   0,
		DelFlag:    0,
		CreateBy:   in.CreateBy,
		CreateTime: time.Now(),
	}
	err = q.WithContext(l.ctx).Create(dict)

	if err != nil {
		logc.Errorf(l.ctx, "添加字典信息失败,参数:%+v,异常:%s", dict, err.Error())
		return nil, errors.New("添加字典信息失败")
	}

	return &sysclient.DictAddResp{}, nil
}
