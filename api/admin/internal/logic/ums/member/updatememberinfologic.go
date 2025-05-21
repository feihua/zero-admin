package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberInfoLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type UpdateMemberInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberInfoLogic {
	return &UpdateMemberInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberInfo 更新会员信息
func (l *UpdateMemberInfoLogic) UpdateMemberInfo(req *types.UpdateMemberInfoReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.MemberInfoService.UpdateMemberInfo(l.ctx, &umsclient.UpdateMemberInfoReq{
		Id:        req.Id,        // 主键ID
		Nickname:  req.Nickname,  // 昵称
		Mobile:    req.Mobile,    // 手机号码
		Avatar:    req.Avatar,    // 头像
		Signature: req.Signature, // 个性签名
		Gender:    req.Gender,    // 性别：0-未知，1-男，2-女
		Birthday:  req.Birthday,  // 生日
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新会员信息成功",
	}, nil
}
