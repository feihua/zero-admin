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

// AddUserLogic 新增用户
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
	userAddReq := sysclient.AddUserReq{
		Avatar:     req.Avatar,
		CreateBy:   l.ctx.Value("userName").(string),
		DeptId:     req.DeptId,
		Email:      req.Email,
		Mobile:     req.Mobile,
		NickName:   req.NickName,
		Password:   req.Password,
		Remark:     req.Remark,
		UserName:   req.UserName,
		UserStatus: req.UserStatus,
	}
	_, err := l.svcCtx.UserService.AddUser(l.ctx, &userAddReq)

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
