package cartitemservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCartItemLogic 删除购物车
/*
Author: LiuFeiHua
Date: 2024/6/12 9:53
*/
type DeleteCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartItemLogic {
	return &DeleteCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCartItem 删除购物车
func (l *DeleteCartItemLogic) DeleteCartItem(in *omsclient.DeleteCartItemReq) (*omsclient.DeleteCartItemResp, error) {
	q := query.OmsCartItem.WithContext(l.ctx).Where(query.OmsCartItem.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		q = q.Where(query.OmsCartItem.ID.In(in.Ids...))
	}
	_, err := q.Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除购物车失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除购物车失败")
	}

	return &omsclient.DeleteCartItemResp{}, nil
}
