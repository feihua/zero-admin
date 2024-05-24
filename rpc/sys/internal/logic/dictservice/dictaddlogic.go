package dictservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictAdd 添加字典信息
// 1.根据字典名称查询字典是否已存在
// 2.如果字典已存在,则直接返回
// 3.字典不存在时,则直接添加字典
func (l *DictAddLogic) DictAdd(in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	err := query.SysDict.WithContext(l.ctx).Create(&model.SysDict{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        in.Sort,
		CreateBy:    in.CreateBy,
		Remarks:     in.Remarks,
		DelFlag:     in.DelFlag,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.DictAddResp{}, nil
}
