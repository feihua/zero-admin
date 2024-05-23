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
func (l *QueryMemberLevelListLogic) QueryMemberLevelList(req *types.ListMemberLevelReq) (resp *types.ListMemberLevelResp, err error) {
	result, err := l.svcCtx.MemberLevelService.MemberLevelList(l.ctx, &umsclient.MemberLevelListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员等级列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员等级失败")
	}

	var list []*types.ListMemberLevelData

	for _, item := range result.List {
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
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员等级成功",
	}, nil
}
