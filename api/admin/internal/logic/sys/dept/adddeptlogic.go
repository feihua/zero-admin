package dept

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddDeptLogic 添加部门信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:16
*/
type AddDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddDeptLogic {
	return AddDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddDept 添加部门信息
func (l *AddDeptLogic) AddDept(req *types.AddDeptReq) (*types.AddDeptResp, error) {
	_, err := l.svcCtx.DeptService.DeptAdd(l.ctx, &sysclient.DeptAddReq{
		DeptName:  req.DeptName,
		ParentId:  req.ParentId,
		OrderNum:  req.OrderNum,
		CreateBy:  l.ctx.Value("userName").(string),
		ParentIds: req.ParentIds,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加部门信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddDeptResp{
		Code:    "000000",
		Message: "添加部门成功",
	}, nil
}
