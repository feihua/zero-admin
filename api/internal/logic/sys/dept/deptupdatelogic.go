package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *DeptUpdateLogic) DeptUpdate(req types.UpdateDeptReq) (*types.UpdateDeptResp, error) {
	_, err := l.svcCtx.Sys.DeptUpdate(l.ctx, &sysclient.DeptUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		LastUpdateBy: "admin",
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
