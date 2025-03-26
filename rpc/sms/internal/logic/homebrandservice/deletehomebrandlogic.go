package homebrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeBrandLogic 删除首页推荐品牌
/*
Author: LiuFeiHua
Date: 2024/6/12 17:53
*/
type DeleteHomeBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeBrandLogic {
	return &DeleteHomeBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeBrand 删除首页推荐品牌
func (l *DeleteHomeBrandLogic) DeleteHomeBrand(in *smsclient.DeleteHomeBrandReq) (*smsclient.DeleteHomeBrandResp, error) {
	q := query.SmsHomeBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除首页推荐品牌失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除首页推荐品牌失败")
	}

	return &smsclient.DeleteHomeBrandResp{}, nil
}
