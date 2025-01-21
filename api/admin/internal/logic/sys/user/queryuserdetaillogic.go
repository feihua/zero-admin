package user

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

// QueryUserDetailLogic 查询用户详情
/*
Author: LiuFeiHua
Date: 2024/5/29 18:20
*/
type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryUserDetail 查询用户详情
func (l *QueryUserDetailLogic) QueryUserDetail(req *types.QueryUserDetailReq) (resp *types.QueryUserDetailResp, err error) {
	detail, err := l.svcCtx.UserService.QueryUserDetail(l.ctx, &sysclient.QueryUserDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryUserDetailData{
		Id:           detail.Id,           // 编号
		UserName:     detail.UserName,     // 用户名
		NickName:     detail.NickName,     // 昵称
		Avatar:       detail.Avatar,       // 头像
		Email:        detail.Email,        // 邮箱
		Mobile:       detail.Mobile,       // 手机号
		UserStatus:   detail.UserStatus,   // 帐号状态（0正常 1停用）
		DeptId:       detail.DeptId,       // 部门id
		Remark:       detail.Remark,       // 备注信息
		IsDeleted:    detail.IsDeleted,    // 是否删除  0：否  1：是
		LoginTime:    detail.LoginTime,    // 登录时间
		LoginIp:      detail.LoginIp,      // 登录ip
		LoginOs:      detail.LoginOs,      // 登录os
		LoginBrowser: detail.LoginBrowser, // 登录浏览器
		CreateBy:     detail.CreateBy,     // 创建者
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新者
		UpdateTime:   detail.UpdateTime,   // 更新时间
	}

	return &types.QueryUserDetailResp{
		Code:    "000000",
		Message: "查询用户详情成功",
		Data:    data,
	}, nil
}
