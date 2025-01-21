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

// DeleteDeptLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:16
*/
type DeleteDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteDeptLogic {
	return DeleteDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteDept 删除部门信息
func (l *DeleteDeptLogic) DeleteDept(req *types.DeleteDeptReq) (*types.DeleteDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeleteDept(l.ctx, &sysclient.DeleteDeptReq{
		Id: req.Id, // 编号
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据deptIds: %+v,删除部门异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteDeptResp{
		Code:    "000000",
		Message: "删除部门成功",
	}, nil
}
