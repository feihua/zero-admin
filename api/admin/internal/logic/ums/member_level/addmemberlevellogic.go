package member_level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
func (l *AddMemberLevelLogic) AddMemberLevel(req *types.AddMemberLevelReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberLevelService.AddMemberLevel(l.ctx, &umsclient.AddMemberLevelReq{
		Name:         req.Name,         // 等级名称
		Level:        req.Level,        // 等级
		GrowthPoint:  req.GrowthPoint,  // 升级所需成长值
		DiscountRate: req.DiscountRate, // 折扣率(0-100)
		FreeFreight:  req.FreeFreight,  // 是否免运费
		CommentExtra: req.CommentExtra, // 是否可评论获取奖励
		Privileges:   req.Privileges,   // 会员特权JSON
		Remark:       req.Remark,       // 备注
		IsEnabled:    req.IsEnabled,    // 是否启用
		CreateBy:     userId,           // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员等级信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
