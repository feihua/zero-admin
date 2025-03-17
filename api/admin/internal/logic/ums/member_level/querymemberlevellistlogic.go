package member_level

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

// QueryMemberLevelListLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:39
*/
type QueryMemberLevelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelListLogic {
	return &QueryMemberLevelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberLevelList 查询会员等级列表
func (l *QueryMemberLevelListLogic) QueryMemberLevelList(req *types.QueryMemberLevelListReq) (resp *types.QueryMemberLevelListResp, err error) {
	result, err := l.svcCtx.MemberLevelService.QueryMemberLevelList(l.ctx, &umsclient.QueryMemberLevelListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		LevelName: strings.TrimSpace(req.LevelName),
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员等级列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员等级失败")
	}

	var list []*types.QueryMemberLevelListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberLevelListData{
			Id:                 item.Id,                 //
			LevelName:          item.LevelName,          // 等级名称
			GrowthPoint:        item.GrowthPoint,        // 成长点
			DefaultStatus:      item.DefaultStatus,      // 是否为默认等级：0->不是；1->是
			FreeFreightPoint:   item.FreeFreightPoint,   // 免运费标准
			CommentGrowthPoint: item.CommentGrowthPoint, // 每次评价获取的成长值
			IsFreeFreight:      item.IsFreeFreight,      // 是否有免邮特权
			IsSignIn:           item.IsSignIn,           // 是否有签到特权
			IsComment:          item.IsComment,          // 是否有评论获奖励特权
			IsPromotion:        item.IsPromotion,        // 是否有专享活动特权
			IsMemberPrice:      item.IsMemberPrice,      // 是否有会员价格特权
			IsBirthday:         item.IsBirthday,         // 是否有生日特权
			Remark:             item.Remark,             // 备注
		})
	}

	return &types.QueryMemberLevelListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员等级成功",
	}, nil
}
