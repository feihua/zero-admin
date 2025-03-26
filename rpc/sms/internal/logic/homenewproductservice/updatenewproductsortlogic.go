package homenewproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNewProductSortLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:31
*/
type UpdateNewProductSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductSortLogic {
	return &UpdateNewProductSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNewProductSort 修改首页新品排序
func (l *UpdateNewProductSortLogic) UpdateNewProductSort(in *smsclient.UpdateNewProductSortReq) (*smsclient.UpdateNewProductSortResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Sort, in.Sort)

	if err != nil {
		logc.Errorf(l.ctx, "修改首页新品排序失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("修改首页新品排序失败")
	}

	return &smsclient.UpdateNewProductSortResp{}, nil
}
