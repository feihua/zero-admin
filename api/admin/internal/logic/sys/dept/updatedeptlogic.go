package dept

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDeptLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type UpdateDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDeptLogic {
	return UpdateDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDept 更新部门信息
func (l *UpdateDeptLogic) UpdateDept(req *types.UpdateDeptReq) (*types.UpdateDeptResp, error) {

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
		logc.Errorf(l.ctx, "更新部门信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDeptResp{
		Code:    "000000",
		Message: "更新部门成功",
	}, nil
}
