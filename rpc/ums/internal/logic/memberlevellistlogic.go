package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.UmsMemberLevelModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberLevelModel.Count()

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

	fmt.Println(list)
	return &ums.MemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
