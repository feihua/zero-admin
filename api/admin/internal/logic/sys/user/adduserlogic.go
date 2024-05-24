package user

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

// AddUserLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 13:57
*/
type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddUser 新增用户
func (l *AddUserLogic) AddUser(req *types.AddUserReq) (*types.AddUserResp, error) {
	userAddReq := sysclient.UserAddReq{
		Email:      req.Email,
		Mobile:     req.Mobile,
		UserName:   req.Name,
		NickName:   req.NickName,
		DeptId:     req.DeptId,
		CreateBy:   l.ctx.Value("userName").(string),
		JobId:      req.JobId,
		UserStatus: req.Status,
	}
	_, err := l.svcCtx.UserService.UserAdd(l.ctx, &userAddReq)

	if err != nil {
		logc.Errorf(l.ctx, "添加用户信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddUserResp{
		Code:    "000000",
		Message: "添加用户成功",
	}, nil
}
