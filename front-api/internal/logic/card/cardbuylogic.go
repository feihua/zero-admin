package card

import (
	"context"
	"encoding/json"

	"zero-admin/common/ctxdata"
	"zero-admin/common/uniqueid"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CardBuyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCardBuyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CardBuyLogic {
	return &CardBuyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CardBuyLogic) CardBuy(req *types.CardBuyReq) (resp *types.CardBuyResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	prepaidCardSn := uniqueid.GenSn(uniqueid.SN_PREFIX_PREPAID_CARD)
	_, err = l.svcCtx.Ums.MemberPrepaidCardRelationAdd(l.ctx, &umsclient.MemberPrepaidCardRelationAddReq{
		MemberId:      memberId,
		PrepaidCardId: req.CardId,
		PrepaidCardSn: prepaidCardSn,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("购买预付卡失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.CardBuyResp{
			Errno:  1,
			Errmsg: "购买预付卡失败",
		}, nil
	}

	return &types.CardBuyResp{
		Errno:  0,
		Errmsg: "购买预付卡完成",
	}, nil
}
