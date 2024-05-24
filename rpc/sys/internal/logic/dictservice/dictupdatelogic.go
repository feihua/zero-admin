package dictservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DictUpdate 更新字典信息
// 1.根据部门id查询部门是否已存在
// 2.如果部门不已存在,则直接返回
// 3.部门存在时,则直接更新部门
func (l *DictUpdateLogic) DictUpdate(in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	q := query.SysDict
	//更新之前查询记录是否存在
	dict, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	_, err = q.WithContext(l.ctx).Updates(&model.SysDict{
		ID:          in.Id,
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        in.Sort,
		CreateBy:    dict.CreateBy,
		CreateTime:  dict.CreateTime,
		UpdateBy:    in.UpdateBy,
		UpdateTime:  &now,
		Remarks:     in.Remarks,
		DelFlag:     in.DelFlag,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.DictUpdateResp{}, nil
}
