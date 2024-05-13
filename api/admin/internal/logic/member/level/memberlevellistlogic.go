package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelListLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:39
*/
type MemberLevelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLevelListLogic {
	return MemberLevelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MemberLevelList 查询会员等级列表
func (l *MemberLevelListLogic) MemberLevelList(req types.ListMemberLevelReq) (*types.ListMemberLevelResp, error) {
	resp, err := l.svcCtx.MemberLevelService.MemberLevelList(l.ctx, &umsclient.MemberLevelListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员等级列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员等级失败")
	}

	var list []*types.ListMemberLevelData

	for _, item := range resp.List {
		list = append(list, &types.ListMemberLevelData{
			Id:                 item.Id,
			Name:               item.LevelName,
			GrowthPoint:        item.GrowthPoint,
			DefaultStatus:      item.DefaultStatus,
			FreeFreightPoint:   item.FreeFreightPoint,
			CommentGrowthPoint: item.CommentGrowthPoint,
			IsFreeFreight:      item.IsFreeFreight,
			IsSignIn:           item.IsSignIn,
			IsComment:          item.IsComment,
			IsPromotion:        item.IsPromotion,
			IsMemberPrice:      item.IsMemberPrice,
			IsBirthday:         item.IsBirthday,
			Note:               item.Remark,
		})
	}

	return &types.ListMemberLevelResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员等级成功",
	}, nil
}
