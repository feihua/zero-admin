package member_level

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

// QueryMemberLevelDetailLogic 查询会员等级表详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type QueryMemberLevelDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLevelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelDetailLogic {
	return &QueryMemberLevelDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberLevelDetail 查询会员等级表详情
func (l *QueryMemberLevelDetailLogic) QueryMemberLevelDetail(req *types.QueryMemberLevelDetailReq) (resp *types.QueryMemberLevelDetailResp, err error) {

	detail, err := l.svcCtx.MemberLevelService.QueryMemberLevelDetail(l.ctx, &umsclient.QueryMemberLevelDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员等级表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberLevelDetailData{
		Id:                 detail.Id,                 //
		LevelName:          detail.LevelName,          // 等级名称
		GrowthPoint:        detail.GrowthPoint,        // 成长点
		DefaultStatus:      detail.DefaultStatus,      // 是否为默认等级：0->不是；1->是
		FreeFreightPoint:   detail.FreeFreightPoint,   // 免运费标准
		CommentGrowthPoint: detail.CommentGrowthPoint, // 每次评价获取的成长值
		IsFreeFreight:      detail.IsFreeFreight,      // 是否有免邮特权
		IsSignIn:           detail.IsSignIn,           // 是否有签到特权
		IsComment:          detail.IsComment,          // 是否有评论获奖励特权
		IsPromotion:        detail.IsPromotion,        // 是否有专享活动特权
		IsMemberPrice:      detail.IsMemberPrice,      // 是否有会员价格特权
		IsBirthday:         detail.IsBirthday,         // 是否有生日特权
		Remark:             detail.Remark,             // 备注
	}
	return &types.QueryMemberLevelDetailResp{
		Code:    "000000",
		Message: "查询会员等级表成功",
		Data:    data,
	}, nil
}
