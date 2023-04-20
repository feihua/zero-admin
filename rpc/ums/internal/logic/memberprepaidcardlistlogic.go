package logic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardListLogic {
	return &MemberPrepaidCardListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardListLogic) MemberPrepaidCardList(in *umsclient.MemberPrepaidCardListReq) (*umsclient.MemberPrepaidCardListResp, error) {

	all, err := l.svcCtx.UmsMemberPrepaidCardModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberPrepaidCardModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员登录记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.MemberPrepaidCardListData
	for _, item := range *all {
		list = append(list, &ums.MemberPrepaidCardListData{
			Id:             item.Id,
			Name:           item.Name,
			FaceValue:      item.FaceValue,
			ExpiredDay:     item.ExpiredDay,
			Status:         item.Status,
			CommissionRate: item.CommissionRate,
			DiscountRate:   item.DiscountRate,
			Note:           item.Note,
		})
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员登录记录列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &umsclient.MemberPrepaidCardListResp{
		Total: count,
		List:  list,
	}, nil
}
