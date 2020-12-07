package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberRuleSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingListLogic {
	return &MemberRuleSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(in *ums.MemberRuleSettingListReq) (*ums.MemberRuleSettingListResp, error) {
	all, _ := l.svcCtx.UmsMemberRuleSettingModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*ums.MemberRuleSettingListData
	for _, item := range *all {

		list = append(list, &ums.MemberRuleSettingListData{
			Id:                item.Id,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   int64(item.ConsumePerPoint),
			LowOrderAmount:    int64(item.LowOrderAmount),
			MaxPointPerOrder:  item.MaxPointPerOrder,
			Type:              item.Type,
		})
	}

	fmt.Println(list)
	return &ums.MemberRuleSettingListResp{
		Total: 10,
		List:  list,
	}, nil

}
