package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberListLogic 查询会员信息列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:28
*/
type QueryMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberListLogic {
	return &QueryMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberList 查询会员信息列表
func (l *QueryMemberListLogic) QueryMemberList(req *types.ListMemberReq) (resp *types.ListMemberResp, err error) {
	result, err := l.svcCtx.MemberService.QueryMemberList(l.ctx, &umsclient.QueryMemberListReq{
		PageNum:      req.Current,
		PageSize:     req.PageSize,
		MemberName:   strings.TrimSpace(req.Username),
		Phone:        strings.TrimSpace(req.Phone),
		MemberStatus: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员信息失败")
	}

	var list []*types.ListMemberData

	for _, item := range result.List {
		list = append(list, &types.ListMemberData{
			Id:                    item.Id,                    //
			MemberLevelId:         item.MemberLevelId,         // 会员等级id
			MemberName:            item.MemberName,            // 用户名
			Nickname:              item.Nickname,              // 昵称
			Phone:                 item.Phone,                 // 手机号码
			MemberStatus:          item.MemberStatus,          // 帐号启用状态:0->禁用；1->启用
			Icon:                  item.Icon,                  // 头像
			Gender:                item.Gender,                // 性别：0->未知；1->男；2->女
			Birthday:              item.Birthday,              // 生日
			City:                  item.City,                  // 所做城市
			Job:                   item.Job,                   // 职业
			PersonalizedSignature: item.PersonalizedSignature, // 个性签名
			SourceType:            item.SourceType,            // 用户来源
			Integration:           item.Integration,           // 积分
			Growth:                item.Growth,                // 成长值
			LotteryCount:          item.LotteryCount,          // 剩余抽奖次数
			HistoryIntegration:    item.HistoryIntegration,    // 历史积分数量
			CreateTime:            item.CreateTime,            // 创建时间
			UpdateTime:            item.UpdateTime,            // 更新时间
		})
	}

	return &types.ListMemberResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员信息成功",
	}, nil
}
