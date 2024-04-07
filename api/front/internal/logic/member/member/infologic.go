package member

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// InfoLogic 获取会员个人信息
/*
Author: LiuFeiHua
Date: 2023/11/28 14:53
*/
type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Info 获取个人信息
func (l *InfoLogic) Info() (resp *types.InfoResp, err error) {
	//从jwt中提取会员id
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	member, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})

	// 获取用户优惠券
	result, err := l.svcCtx.CouponHistoryService.CouponCount(l.ctx, &smsclient.CouponCountReq{
		MemberId: memberId,
	})

	var count int64 = 0
	//获取不到优惠券数量的时候,默认返回0
	if result != nil {
		count = result.Total
	}
	return &types.InfoResp{
		Code:    0,
		Message: "查询会员信息",
		Data: types.MemberData{
			Id:                    member.Id,
			MemberLevelId:         member.MemberLevelId,
			Username:              member.Username,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			Status:                member.Status,
			CreateTime:            member.CreateTime,
			Icon:                  member.Icon,
			Gender:                member.Gender,
			Birthday:              member.Birthday,
			City:                  member.City,
			Job:                   member.Job,
			PersonalizedSignature: member.PersonalizedSignature,
			SourceType:            member.SourceType,
			Integration:           member.Integration,
			Growth:                member.Growth,
			LuckeyCount:           member.LuckeyCount,
			HistoryIntegration:    member.HistoryIntegration,
			CouponCount:           count,
		},
	}, nil
}
