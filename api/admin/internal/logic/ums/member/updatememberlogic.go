package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLogic 更新会员信息
/*
Author: LiuFeiHua
Date: 2024/5/23 9:26
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
func (l *UpdateMemberLogic) UpdateMember(req *types.UpdateMemberReq) (resp *types.UpdateMemberResp, err error) {
	_, err = l.svcCtx.MemberService.UpdateMember(l.ctx, &umsclient.UpdateMemberReq{
		Id:                    req.Id,                    //
		MemberLevelId:         req.MemberLevelId,         // 会员等级id
		MemberName:            req.MemberName,            // 用户名
		Password:              req.Password,              // 密码
		Nickname:              req.Nickname,              // 昵称
		Phone:                 req.Phone,                 // 手机号码
		MemberStatus:          req.MemberStatus,          // 帐号启用状态:0->禁用；1->启用
		Icon:                  req.Icon,                  // 头像
		Gender:                req.Gender,                // 性别：0->未知；1->男；2->女
		Birthday:              req.Birthday,              // 生日
		City:                  req.City,                  // 所做城市
		Job:                   req.Job,                   // 职业
		PersonalizedSignature: req.PersonalizedSignature, // 个性签名
		SourceType:            req.SourceType,            // 用户来源
		Integration:           req.Integration,           // 积分
		Growth:                req.Growth,                // 成长值
		LotteryCount:          req.LotteryCount,          // 剩余抽奖次数
		HistoryIntegration:    req.HistoryIntegration,    // 历史积分数量
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员信息失败")
	}

	return &types.UpdateMemberResp{
		Code:    "000000",
		Message: "更新会员信息成功",
	}, nil
}
