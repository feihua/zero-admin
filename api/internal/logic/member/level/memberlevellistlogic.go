package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberLevelListLogic) MemberLevelList(req types.ListMemberLevelReq) (*types.ListMemberLevelResp, error) {
	resp, err := l.svcCtx.Ums.MemberLevelList(l.ctx, &umsclient.MemberLevelListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员等级列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员等级失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberLevelData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberLevelData{
			Id:                    item.Id,
			Name:                  item.Name,
			GrowthPoint:           item.GrowthPoint,
			DefaultStatus:         item.DefaultStatus,
			FreeFreightPoint:      float64(item.FreeFreightPoint),
			CommentGrowthPoint:    item.CommentGrowthPoint,
			PriviledgeFreeFreight: item.PriviledgeFreeFreight,
			PriviledgeSignIn:      item.PriviledgeSignIn,
			PriviledgeComment:     item.PriviledgeComment,
			PriviledgePromotion:   item.PriviledgePromotion,
			PriviledgeMemberPrice: item.PriviledgeMemberPrice,
			PriviledgeBirthday:    item.PriviledgeBirthday,
			Note:                  item.Note,
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
