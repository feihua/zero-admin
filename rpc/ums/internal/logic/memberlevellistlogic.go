package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelListLogic {
	return &MemberLevelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelListLogic) MemberLevelList(in *ums.MemberLevelListReq) (*ums.MemberLevelListResp, error) {
	all, err := l.svcCtx.UmsMemberLevelModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberLevelModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员等级列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*ums.MemberLevelListData
	for _, item := range *all {

		list = append(list, &ums.MemberLevelListData{
			Id:                    item.Id,
			Name:                  item.Name,
			GrowthPoint:           item.GrowthPoint,
			DefaultStatus:         item.DefaultStatus,
			FreeFreightPoint:      int64(item.FreeFreightPoint),
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员等级列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
