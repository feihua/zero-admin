package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeptUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptUpdateLogic {
	return DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeptUpdate 更新部门信息
func (l *DeptUpdateLogic) DeptUpdate(req types.UpdateDeptReq) (*types.UpdateDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeptUpdate(l.ctx, &sysclient.DeptUpdateReq{
		Id:        req.Id,
		DeptName:  req.DeptName,
		ParentId:  req.ParentId,
		OrderNum:  req.OrderNum,
		UpdateBy:  l.ctx.Value("userName").(string),
		ParentIds: req.ParentIds,
		DelFlag:   req.DelFlag,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新机构信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新机构失败")
	}

	return &types.UpdateDeptResp{
		Code:    "000000",
		Message: "更新机构成功",
	}, nil
}
