package card

import (
	"context"
	"encoding/json"

	"zero-admin/common/ctxdata"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CardMyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCardMyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CardMyListLogic {
	return &CardMyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CardMyListLogic) CardMyList(req *types.CardMyListReq) (resp *types.CardMyListResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	result, err := l.svcCtx.Ums.MemberPrepaidCardRelationList(l.ctx, &umsclient.MemberPrepaidCardRelationListReq{
		Current:  req.Page,
		PageSize: req.Limit,
		MemberId: memberId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询用户收货地址列表失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.CardMyListResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	var list []types.CardMyList

	for _, v := range result.List {
		list = append(list, types.CardMyList{
			ID:            v.Id,
			PrepaidCardId: v.PrepaidCardId,
			PrepaidCardSn: v.PrepaidCardSn,
			Balance:       v.Balance,
			Status:        v.Status,
			CreateTime:    v.CreateTime,
			ExpiredTime:   v.ExpiredTime,
			UpdateTime:    v.UpdateTime,
			Note:          v.Note,
		})
	}

	return &types.CardMyListResp{
		Errno: 0,
		Data: types.CardMyListData{
			CardMyList: list,
		},
		Errmsg: "查询完成",
	}, nil
}
