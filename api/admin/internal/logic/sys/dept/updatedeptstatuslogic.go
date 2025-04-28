package dept

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDeptStatusLogic 更新部门状态
/*
Author: LiuFeiHua
Date: 2024/5/29 16:40
*/
type UpdateDeptStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeptStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptStatusLogic {
	return &UpdateDeptStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDeptStatus 更新部门状态
func (l *UpdateDeptStatusLogic) UpdateDeptStatus(req *types.UpdateDeptStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.DeptService.UpdateDeptStatus(l.ctx, &sysclient.UpdateDeptStatusReq{
		Id:       req.Id,     // 部门id
		Status:   req.Status, // 部门状态（0：停用，1:正常）
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新部门状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
