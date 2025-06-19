package member

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2025/6/19 11:11
*/
type UpdateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMember 更新会员信息
func (l *UpdateMemberLogic) UpdateMember(req *types.UpdateMemberReq) (resp *types.MemberResp, err error) {
	_, err = l.svcCtx.MemberService.UpdateMemberInfo(l.ctx, &umsclient.UpdateMemberInfoReq{
		Id:        req.Id,        // 主键ID
		Nickname:  req.Nickname,  // 昵称
		Mobile:    req.Mobile,    // 手机号码
		Avatar:    req.Avatar,    // 头像
		Signature: req.Signature, // 个性签名
		Gender:    req.Gender,    // 性别：0-未知，1-男，2-女
		Birthday:  req.Birthday,  // 生日
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.MemberResp{
		Code:    "000000",
		Message: "更新会员信息成功",
	}, nil
}
