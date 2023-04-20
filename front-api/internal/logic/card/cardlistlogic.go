package card

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CardListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCardListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CardListLogic {
	return &CardListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CardListLogic) CardList(req *types.CardListReq) (resp *types.CardListResp, err error) {
	result, err := l.svcCtx.Ums.MemberPrepaidCardList(l.ctx, &umsclient.MemberPrepaidCardListReq{
		Current:  1,
		PageSize: 100,
	})
	if err != nil {
		// reqStr, _ := json.Marshal(req)
		// logx.WithContext(l.ctx).Errorf("查询用户收货地址列表失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.CardListResp{
			Errno:  1,
			Errmsg: "预付卡功能维护中",
		}, nil
	}
	var list []types.CardList

	for _, v := range result.List {
		list = append(list, types.CardList{
			ID:             v.Id,
			Name:           v.Name,
			FaceValue:      v.FaceValue,
			ExpiredDay:     v.ExpiredDay,
			Status:         v.Status,
			CommissionRate: v.CommissionRate,
			DiscountRate:   v.DiscountRate,
			Note:           "",
		})
	}
	return &types.CardListResp{
		Errno:  0,
		Errmsg: "查询完成",
		Data: types.CardListData{
			// CartTotal: count,
			CardList: list,
		},
	}, nil

}
