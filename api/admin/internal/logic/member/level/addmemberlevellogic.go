package level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberLevelLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:37
*/
type AddMemberLevelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLevelLogic {
	return &AddMemberLevelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberLevel 添加会员等级
func (l *AddMemberLevelLogic) AddMemberLevel(req *types.AddMemberLevelReq) (resp *types.AddMemberLevelResp, err error) {
	_, err = l.svcCtx.MemberLevelService.AddMemberLevel(l.ctx, &umsclient.AddMemberLevelReq{
		LevelName:          req.LevelName,          // 等级名称
		GrowthPoint:        req.GrowthPoint,        // 成长点
		DefaultStatus:      req.DefaultStatus,      // 是否为默认等级：0->不是；1->是
		FreeFreightPoint:   req.FreeFreightPoint,   // 免运费标准
		CommentGrowthPoint: req.CommentGrowthPoint, // 每次评价获取的成长值
		IsFreeFreight:      req.IsFreeFreight,      // 是否有免邮特权
		IsSignIn:           req.IsSignIn,           // 是否有签到特权
		IsComment:          req.IsComment,          // 是否有评论获奖励特权
		IsPromotion:        req.IsPromotion,        // 是否有专享活动特权
		IsMemberPrice:      req.IsMemberPrice,      // 是否有会员价格特权
		IsBirthday:         req.IsBirthday,         // 是否有生日特权
		Remark:             req.Remark,             // 备注
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员等级信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员等级失败")
	}

	return &types.AddMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
