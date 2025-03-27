package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLogic 更新会员表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:58
*/
type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMember 更新会员表
func (l *UpdateMemberLogic) UpdateMember(in *umsclient.UpdateMemberReq) (*umsclient.UpdateMemberResp, error) {
	birthday, _ := time.Parse("2006-01-02", in.Birthday)
	_, err := query.UmsMember.WithContext(l.ctx).Updates(&model.UmsMember{
		ID:                    in.Id,                    //
		MemberLevelID:         in.MemberLevelId,         // 会员等级id
		MemberName:            in.MemberName,            // 用户名
		Nickname:              in.Nickname,              // 昵称
		Phone:                 in.Phone,                 // 手机号码
		MemberStatus:          in.MemberStatus,          // 帐号启用状态:0->禁用；1->启用
		Icon:                  in.Icon,                  // 头像
		Gender:                in.Gender,                // 性别：0->未知；1->男；2->女
		Birthday:              &birthday,                // 生日
		City:                  in.City,                  // 所做城市
		Job:                   in.Job,                   // 职业
		PersonalizedSignature: in.PersonalizedSignature, // 个性签名
		SourceType:            in.SourceType,            // 用户来源
		Integration:           in.Integration,           // 积分
		Growth:                in.Growth,                // 成长值
		LotteryCount:          in.LotteryCount,          // 剩余抽奖次数
		HistoryIntegration:    in.HistoryIntegration,    // 历史积分数量
	})
	if err != nil {
		logc.Errorf(l.ctx, "更新会员失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员失败")
	}

	return &umsclient.UpdateMemberResp{}, nil
}
